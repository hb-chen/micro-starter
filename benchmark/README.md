# Benchmark
[rpcx-benchmark](https://github.com/rpcx-ecosystem/rpcx-benchmark)

- MPB开发环境
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