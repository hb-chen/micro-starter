package main

import (
	"flag"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/montanaflynn/stats"
	"golang.org/x/net/context"
	"github.com/afex/hystrix-go/hystrix"

	"github.com/micro/go-log"
	//micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector/cache"
	"github.com/micro/go-plugins/transport/tcp"
	breaker "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber"

	proto "github.com/hb-go/micro/benchmark/proto"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")
var host = flag.String("s", "127.0.0.1:8972", "server ip and port")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	selected := -1
	servers := strings.Split(*host, ",")
	sNum := len(servers)

	log.Logf("Servers: %+v\n\n", servers)

	log.Logf("concurrency: %d\nrequests per client: %d\n\n", n, m)

	args := prepareArgs()

	// Hystrix配置
	hystrix.ConfigureCommand("hello.Hello.Say", hystrix.CommandConfig{
		MaxConcurrentRequests: hystrix.DefaultMaxConcurrent * 2,
	})

	//b, _ := proto.Marshal(args)
	//log.Logf("message size: %d bytes\n\n", len(b))

	var wg sync.WaitGroup
	wg.Add(n * m)

	var trans uint64
	var transOK uint64

	d := make([][]int64, n, n)

	//it contains warmup time but we can ignore it
	totalT := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		dt := make([]int64, 0, m)
		d = append(d, dt)
		selected = (selected + 1) % sNum

		go func(i int, selected int) {
			//service := micro.NewService(
			//	micro.Name("hello.client"),
			//	micro.Transport(tcp.NewTransport()),
			//	micro.Selector(cache.NewSelector(cache.TTL(time.Second*3000))),
			//)
			//c := proto.NewHelloService("hello", service.Client())

			helloClient := client.NewClient(
				client.Transport(tcp.NewTransport()),
				//client.ContentType("application/protobuf"),
				client.Selector(cache.NewSelector(cache.TTL(time.Second*3000))),
				client.Retries(1),
				client.PoolSize(10),
				client.RequestTimeout(time.Millisecond*100),
				client.Wrap(breaker.NewClientWrapper()),
				client.Wrap(ratelimit.NewClientWrapper(10)),
			)
			c := proto.NewHelloService("hello", helloClient)

			//warmup
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
					log.Logf("error:%v", err)
				}

				atomic.AddUint64(&trans, 1)
				wg.Done()
			}

			// c.Close()

		}(i, selected)

	}

	wg.Wait()
	totalT = time.Now().UnixNano() - totalT
	totalT = totalT / 1000000
	log.Logf("took %d ms for %d requests\n", totalT, n*m)

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
	p99, _ := stats.Percentile(totalD2, 99.9)

	log.Logf("sent     requests    : %d\n", n*m)
	log.Logf("received requests    : %d\n", atomic.LoadUint64(&trans))
	log.Logf("received requests_OK : %d\n", atomic.LoadUint64(&transOK))
	log.Logf("throughput  (TPS)    : %d\n", int64(n*m)*1000/totalT)
	log.Logf("mean: %.f ns, median: %.f ns, max: %.f ns, min: %.f ns, p99: %.f ns\n", mean, median, max, min, p99)
	log.Logf("mean: %d ms, median: %d ms, max: %d ms, min: %d ms, p99: %d ms\n", int64(mean/1000000), int64(median/1000000), int64(max/1000000), int64(min/1000000), int64(p99/1000000))

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
