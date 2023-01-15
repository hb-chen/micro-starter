FROM alpine:3.10
ADD dist/account_linux_amd64_v1/bin/account /opt/service/main
WORKDIR /opt/service
ENTRYPOINT [ "./main" ]
