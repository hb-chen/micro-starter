FROM alpine:3.10
ADD dist/example_linux_amd64_v1/bin/example /opt/service/main
WORKDIR /opt/service
ENTRYPOINT [ "./main" ]
