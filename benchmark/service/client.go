package service

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"text/tabwriter"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"github.com/montanaflynn/stats"
	"golang.org/x/net/context"

	"github.com/hb-go/micro/benchmark/proto"
)

type NewClient func() client.Client

func ClientWith(serviceName string, nc NewClient, concurrency, total int) {
	flag.Parse()
	n := concurrency
	m := total / n

	log.Logf("service: %s testing", serviceName)
	log.Logf("concurrency: %d, requests per client: %d \n", n, m)

	args := prepareArgs()

	// Hystrix配置
	// hystrix.ConfigureCommand(serviceName+".Hello.Say", hystrix.CommandConfig{
	// 	MaxConcurrentRequests: hystrix.DefaultMaxConcurrent * 100,
	// })

	var wg sync.WaitGroup
	wg.Add(n * m)

	var trans uint64
	var transOK uint64

	d := make([][]int64, n, n)

	// it contains warmup time but we can ignore it
	totalT := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		dt := make([]int64, 0, m)
		d = append(d, dt)

		go func(i int) {
			helloClient := nc()
			c := proto.NewHelloService(serviceName, helloClient)

			// warmup
			for j := 0; j < 5; j++ {
				c.Say(context.Background(), args)
			}

			for j := 0; j < m; j++ {
				t := time.Now().UnixNano()
				reply, err := c.Say(context.Background(), args)
				t = time.Now().UnixNano() - t

				d[i] = append(d[i], t)

				if err == nil && *reply.Field1 == "OK" {
					atomic.AddUint64(&transOK, 1)
				} else {
					// log.Logf("error:%v", err)
				}

				atomic.AddUint64(&trans, 1)
				wg.Done()
			}
		}(i)

	}

	wg.Wait()
	totalT = time.Now().UnixNano() - totalT
	totalT = totalT / 1000000

	totalD := make([]int64, 0, n*m)
	for _, k := range d {
		totalD = append(totalD, k...)
	}
	totalD2 := make([]float64, 0, n*m)
	for _, k := range totalD {
		totalD2 = append(totalD2, float64(k))
	}

	mean, _ := stats.Mean(totalD2)
	median, _ := stats.Median(totalD2)
	max, _ := stats.Max(totalD2)
	min, _ := stats.Min(totalD2)
	p90, _ := stats.Percentile(totalD2, 90)
	p99, _ := stats.Percentile(totalD2, 99)
	tps := int64(n*m) * 1000 / totalT

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', 0)

	fmt.Println()
	fmt.Fprintf(w, "took\t(ms)\t: %d\n", totalT)
	fmt.Fprintf(w, "sent\trequests\t: %d\n", n*m)
	fmt.Fprintf(w, "received\trequests\t: %d\n", atomic.LoadUint64(&trans))
	fmt.Fprintf(w, "received\trequests_OK\t: %d\n", atomic.LoadUint64(&transOK))
	fmt.Fprintf(w, "throughput\t(TPS)\t: %d\n", tps)
	w.Flush()

	fmt.Println()
	fmt.Fprintf(w, "concurrency\tmean\tmedian\tmax\tmin\tp90\tp99\tTPS\n")
	fmt.Fprintf(w, "%d\t%.fns\t%.fns\t%.fns\t%.fns\t%.fns\t%.fns\t%d\n", concurrency, mean, median, max, min, p90, p99, tps)
	fmt.Fprintf(w, "%d\t%.3fms\t%.3fms\t%.3fms\t%.3fms\t%.3fms\t%.3fms\t%d\n", concurrency, mean/1000000, median/1000000, max/1000000, min/1000000, p90/1000000, p99/1000000, tps)
	w.Flush()
}

func prepareArgs() *proto.BenchmarkMessage {
	b := true
	var i int32 = 100000
	var i64 int64 = 100000
	var s = "许多往事在眼前一幕一幕，变的那麼模糊"

	var args proto.BenchmarkMessage

	v := reflect.ValueOf(&args).Elem()
	num := v.NumField()
	for k := 0; k < num; k++ {
		field := v.Field(k)
		if field.Type().Kind() == reflect.Ptr {
			switch v.Field(k).Type().Elem().Kind() {
			case reflect.Int, reflect.Int32:
				field.Set(reflect.ValueOf(&i))
			case reflect.Int64:
				field.Set(reflect.ValueOf(&i64))
			case reflect.Bool:
				field.Set(reflect.ValueOf(&b))
			case reflect.String:
				field.Set(reflect.ValueOf(&s))
			}
		} else {
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.SetInt(100000)
			case reflect.Bool:
				field.SetBool(true)
			case reflect.String:
				field.SetString(s)
			}
		}

	}
	return &args
}
