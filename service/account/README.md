# Console Service

This is the Console service

## Usage

Generate the proto code

```
make proto
```

Run the service

```
micro run --name console .
```

Call service

```shell script
$ curl "http://localhost:8080/console/account/login?username=hbchen"
{"token":"token"}

$ curl "http://localhost:8080/console/account/info"
{"avatar":"https://avatars3.githubusercontent.com/u/730866?s=460&v=4","name":"Hobo"}
```

```shell script
$ micro console --help      
NAME:
        micro console

VERSION:
        latest

USAGE:
        micro console [command]

COMMANDS:
        account info
        account login
        account logout
```

```shell script
$ micro console account login --username=test
{
        "token": "token"
}

$ micro console account info
{
        "name": "Hobo",
        "avatar": "https://avatars3.githubusercontent.com/u/730866?s=460\u0026v=4"
}
```
