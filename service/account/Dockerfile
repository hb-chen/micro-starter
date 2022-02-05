FROM alpine:3.10
ADD dist/micro_linux_amd64/bin/example /opt/service/main
WORKDIR /opt/service
ENTRYPOINT [ "main" ]
