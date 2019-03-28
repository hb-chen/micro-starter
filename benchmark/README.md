# Benchmark
[rpcx-benchmark](https://github.com/rpcx-ecosystem/rpcx-benchmark)

- MBP开发环境
```bash
go run micro_mclient.go -c 100 -n 50000
2018/06/27 17:48:00 Servers: [127.0.0.1:8972]

2018/06/27 17:51:28 concurrency: 100
requests per client: 500

2018/06/27 17:51:32 took 3885 ms for 50000 requests
2018/06/27 17:51:32 sent     requests    : 50000
2018/06/27 17:51:32 received requests    : 50000
2018/06/27 17:51:32 received requests_OK : 50000
2018/06/27 17:51:32 throughput  (TPS)    : 12870
2018/06/27 17:51:32 mean: 5376750 ns, median: 4338000 ns, max: 96008000 ns, min: 156000 ns, p99: 45684000 ns
2018/06/27 17:51:32 mean: 5 ms, median: 4 ms, max: 96 ms, min: 0 ms, p99: 45 ms
```

- 限流
```go
helloClient := client.NewClient(
    client.Transport(tcp.NewTransport()),
    client.ContentType("application/proto-rpc"),
    client.Selector(cache.NewSelector(cache.TTL(time.Second*3000))),
    client.Retries(1),
    client.PoolSize(100),
    client.RequestTimeout(time.Millisecond*1000),
    client.Wrap(ratelimit.NewClientWrapper(1)),
)
```
```bash
# client.Wrap(ratelimit.NewClientWrapper(1))

➜  benchmark git:(master) ✗ go run micro_client.go -c 50 -n 1000
2018/07/06 21:27:13 concurrency: 50
requests per client: 20

2018/07/06 21:27:37 took 24018 ms for 1000 requests
2018/07/06 21:27:37 sent     requests    : 1000
2018/07/06 21:27:37 received requests    : 1000
2018/07/06 21:27:37 received requests_OK : 1000
2018/07/06 21:27:37 throughput  (TPS)    : 41
2018/07/06 21:27:37 mean: 1000159226 ns, median: 999944000 ns, max: 1009774000 ns, min: 992741000 ns, p99: 1009760500 ns
2018/07/06 21:27:37 mean: 1000 ms, median: 999 ms, max: 1009 ms, min: 992 ms, p99: 1009 ms

# client.Wrap(ratelimit.NewClientWrapper(10))
➜  benchmark git:(master) ✗ go run micro_client.go -c 50 -n 1000
2018/07/06 21:27:53 concurrency: 50
requests per client: 20

2018/07/06 21:27:56 took 2414 ms for 1000 requests
2018/07/06 21:27:56 sent     requests    : 1000
2018/07/06 21:27:56 received requests    : 1000
2018/07/06 21:27:56 received requests_OK : 1000
2018/07/06 21:27:56 throughput  (TPS)    : 414
2018/07/06 21:27:56 mean: 100238663 ns, median: 99990000 ns, max: 109895000 ns, min: 93010000 ns, p99: 109669000 ns
2018/07/06 21:27:56 mean: 100 ms, median: 99 ms, max: 109 ms, min: 93 ms, p99: 109 ms

# client.Wrap(ratelimit.NewClientWrapper(100))
➜  benchmark git:(master) ✗ go run micro_client.go -c 50 -n 1000
2018/07/06 21:28:18 concurrency: 50
requests per client: 20

2018/07/06 21:28:18 took 267 ms for 1000 requests
2018/07/06 21:28:18 sent     requests    : 1000
2018/07/06 21:28:18 received requests    : 1000
2018/07/06 21:28:18 received requests_OK : 1000
2018/07/06 21:28:18 throughput  (TPS)    : 3745
2018/07/06 21:28:18 mean: 8221555 ns, median: 8683500 ns, max: 29595000 ns, min: 204000 ns, p99: 29448500 ns
2018/07/06 21:28:18 mean: 8 ms, median: 8 ms, max: 29 ms, min: 0 ms, p99: 29 ms

```