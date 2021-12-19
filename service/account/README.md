# Account Service

This is the Account service

## Usage

Generate the proto code

```
make proto
```

Run the service

```
go run main.go --profile starter-local
```

Call service

```shell script
curl -XPOST -H"Content-Type: application/json" "http://localhost:8080/account/login" -d '{"username":"admin","password":"123456"}'

curl "http://localhost:8080/account/info?id=1"
```

