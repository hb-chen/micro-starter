FROM alpine:3.10

ADD dist/micro_linux_amd64/bin/micro /opt/service/micro

EXPOSE 8080
WORKDIR /opt/service

ENTRYPOINT [ "./micro" ]
