FROM alpine:3.10
ADD dist/greeting_linux_amd64_v1/bin/greeting /opt/service/main
WORKDIR /opt/service
ENTRYPOINT [ "./main" ]
