# Benchmark
>ref:[rpcx-benchmark](https://github.com/rpcx-ecosystem/rpcx-benchmark)

测试影响`go-micro`服务间通信效率的三个组件：`transport`、`server`以及`codec`，主要做不同插件间的横向对比。

**测试环境**

- MBP
- go **1.12.5**
- go-micro **v1.2.0**
- go-plugins **v1.1.0** 

## Transport + Server对比

`transport`和`server`的对比使用`100`并发，完成`10W`请求进行测试

### 结果对比
从结果看`tcp`+`rpc`吞吐量最高，分别比较：

- `transport`比较结果`tcp`>`grpc`>`utp`
- `server`比较结果`rpc`>`grpc`

T+S|平均<br/>(ms)|中位<br/>(ms)|最大<br/>(ms)|最小<br/>(ms)|P90<br/>(ms)|P99<br/>(ms)|TPS
---|---|---|---|---|---|---|---
tcp+rpc|7.236|5.629|101.506|0.177|13.338|35.880|13192
grpc+rpc|8.668|7.964|101.280|0.251|12.744|21.672|11166
utp+rpc|11.824|11.600|53.183|0.204|15.575|21.334|8252
grpc+grpc|8.924|8.181|134.434|0.286|13.211|22.973|10845

> 在开始的测试中有个误区，`grpc`服务并不使用`transport`，包括`http`服务，`transport`仅在使用`go-micro`的`rpc`服务时有效

## Codec对比
`transport`和`server`分别使用`tcp`和`rpc`对比不同`codec`性能，因为并发`100`时不同`codec`的失败率差别比较大，所以使用`50`并发，完成`10W`请求进行测试

### 结果对比
对比结果：`protobuf`>`proto-rpc`>`grpc`>`json`>`grpc+json`>`json-rpc`>`bsonrpc`

CODEC|平均<br/>(ms)|中位<br/>(ms)|最大<br/>(ms)|最小<br/>(ms)|P90<br/>(ms)|P99<br/>(ms)|TPS
-----|------|------|------|------|------|------|------
grpc|3.937|2.979|90.004|0.180|7.184|19.355|12310
grpc+json|6.085|4.694|149.861|0.342|10.365|31.837|8000
protobuf|3.661|2.707|96.636|0.156|6.542|20.261|13150
json|5.402|4.122|122.360|0.225|9.186|30.474|8896
json-rpc|6.380|4.878|115.141|0.288|11.150|33.395|7631
proto-rpc|3.692|2.729|101.010|0.180|6.701|19.454|13041
bsonrpc|7.912|5.979|132.041|0.354|14.789|40.414|6145

默认`chdec`如下
```go
DefaultCodecs = map[string]codec.NewCodec{
    "application/grpc":         grpc.NewCodec,
    "application/grpc+json":    grpc.NewCodec,
    "application/grpc+proto":   grpc.NewCodec,
    "application/json":         json.NewCodec,
    "application/json-rpc":     jsonrpc.NewCodec,
    "application/protobuf":     proto.NewCodec,
    "application/proto-rpc":    protorpc.NewCodec,
    "application/octet-stream": raw.NewCodec,
}
```

另外`go-plugins`提供三个`codec`插件，在`server`和`client`初始化时自定义添加
```go
server.Codec("application/msgpackrpc", msgpackrpc.NewCodec),
server.Codec("application/bsonrpc", bsonrpc.NewCodec),
server.Codec("application/jsonrpc2", jsonrpc2.NewCodec),

client.Codec("application/msgpackrpc", msgpackrpc.NewCodec),
client.Codec("application/bsonrpc", bsonrpc.NewCodec),
client.Codec("application/jsonrpc2", jsonrpc2.NewCodec),
```

实际测试时由于不同原因`raw` 、`msgpackrpc`和`jsonrpc2`运行失败未测试，`grpc+proto`与`grpc`实现一致未测试

>raw.NewCodec<br/>error:{"id":"go.micro.client.codec","code":500,"detail":"failed to write: field1:……<br/>
 msgpackrpc.NewCodec，需要实现EncodeMsg(*Writer)<br/>error:{"id":"go.micro.client.codec","code":500,"detail":"Not encodable","status":"Internal Server Error"}<br/>
 jsonrpc2.NewCodec<br/>error:{"id":"go.micro.client.transport","code":500,"detail":"EOF","status":"Internal Server Error"}

## 详细数据

### Transport + Server 

#### 测试命令
```bash
$ go run server.go
$ go run client.go -c 100 -n 100000
```

**tcp + rpc**
```bash
took       (ms)        : 7580
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99999
throughput (TPS)       : 13192

concurrency mean      median    max         min       p90        p99        TPS
100         7235584ns 5629000ns 101506000ns 177000ns  13338000ns 35880000ns 13192
100         7.236ms   5.629ms   101.506ms   0.177ms   13.338ms   35.880ms   13192

```

**tcp + grpc**
```bash
took       (ms)        : 9295
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99994
throughput (TPS)       : 10758

concurrency mean      median    max         min       p90        p99        TPS
100         9005583ns 8171000ns 110580000ns 300000ns  13510000ns 24317000ns 10758
100         9.006ms   8.171ms   110.580ms   0.300ms   13.510ms   24.317ms   10758
```

**grpc + rpc**
```bash
took       (ms)        : 8955
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99999
throughput (TPS)       : 11166

concurrency mean      median    max         min       p90        p99        TPS
100         8668191ns 7964000ns 101280000ns 251000ns  12744000ns 21672000ns 11166
100         8.668ms   7.964ms   101.280ms   0.251ms   12.744ms   21.672ms   11166
```


**grpc + gRPC**
```bash
took       (ms)        : 9220
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99995
throughput (TPS)       : 10845

concurrency mean      median    max         min       p90        p99        TPS
100         8924043ns 8181000ns 134434000ns 286000ns  13211000ns 22973000ns 10845
100         8.924ms   8.181ms   134.434ms   0.286ms   13.211ms   22.973ms   10845
```


**utp + rpc**
```bash
took       (ms)        : 12117
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 100000
throughput (TPS)       : 8252

concurrency mean       median     max        min       p90        p99        TPS
100         11823520ns 11600000ns 53183000ns 204000ns  15575000ns 21334000ns 8252
100         11.824ms   11.600ms   53.183ms   0.204ms   15.575ms   21.334ms   8252
```

**utp + grpc**
```bash
took       (ms)        : 9367
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99999
throughput (TPS)       : 10675

concurrency mean      median    max         min       p90        p99        TPS
100         9063046ns 8213000ns 102052000ns 289000ns  13553000ns 23920000ns 10675
100         9.063ms   8.213ms   102.052ms   0.289ms   13.553ms   23.920ms   10675
```

### Codec

#### 测试命令

```bash
$ cd benchmark/tcp_rpc
$ go run server.go
$ go run client.go -c 50 -n 100000 -ct grpc
$ go run client.go -c 50 -n 100000 -ct grpc+json
$ go run client.go -c 50 -n 100000 -ct protobuf
$ go run client.go -c 50 -n 100000 -ct json
$ go run client.go -c 50 -n 100000 -ct json-rpc
$ go run client.go -c 50 -n 100000 -ct proto-rpc
$ go run client.go -c 50 -n 100000 -ct bsonrpc
```

**grpc**
```bash
took       (ms)        : 8123
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 100000
throughput (TPS)       : 12310

concurrency mean      median    max        min       p90       p99        TPS
50          3936652ns 2979000ns 90004000ns 180000ns  7184000ns 19355000ns 12310
50          3.937ms   2.979ms   90.004ms   0.180ms   7.184ms   19.355ms   12310
```

**grpc+json**
```bash
took       (ms)        : 12500
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99961
throughput (TPS)       : 8000

concurrency mean      median    max         min       p90        p99        TPS
50          6084709ns 4694000ns 149861000ns 342000ns  10365000ns 31837000ns 8000
50          6.085ms   4.694ms   149.861ms   0.342ms   10.365ms   31.837ms   8000
```

**protobuf**
```bash
took       (ms)        : 7604
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 100000
throughput (TPS)       : 13150

concurrency mean      median    max        min       p90       p99        TPS
50          3660854ns 2707000ns 96636000ns 156000ns  6542000ns 20261000ns 13150
50          3.661ms   2.707ms   96.636ms   0.156ms   6.542ms   20.261ms   13150
```

**json**
```bash
took       (ms)        : 11241
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99922
throughput (TPS)       : 8896

concurrency mean      median    max         min       p90       p99        TPS
50          5401532ns 4121500ns 122360000ns 225000ns  9186000ns 30474000ns 8896
50          5.402ms   4.122ms   122.360ms   0.225ms   9.186ms   30.474ms   8896
```

**json-rpc**
```bash
took       (ms)        : 13103
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99971
throughput (TPS)       : 7631

concurrency mean      median    max         min       p90        p99        TPS
50          6380308ns 4878000ns 115141000ns 288000ns  11150000ns 33395000ns 7631
50          6.380ms   4.878ms   115.141ms   0.288ms   11.150ms   33.395ms   7631
```

**proto-rpc**
```bash
took       (ms)        : 7668
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99998
throughput (TPS)       : 13041

concurrency mean      median    max         min       p90       p99        TPS
50          3692281ns 2729000ns 101010000ns 180000ns  6701000ns 19454000ns 13041
50          3.692ms   2.729ms   101.010ms   0.180ms   6.701ms   19.454ms   13041
```

**bsonrpc**
```bash
took       (ms)        : 16272
sent       requests    : 100000
received   requests    : 100000
received   requests_OK : 99933
throughput (TPS)       : 6145

concurrency mean      median    max         min       p90        p99        TPS
50          7911887ns 5979000ns 132041000ns 354000ns  14789000ns 40414000ns 6145
50          7.912ms   5.979ms   132.041ms   0.354ms   14.789ms   40.414ms   6145
```