FROM alpine:3.10

ADD dist/micro_linux_amd64_v1/bin/micro /opt/service/micro

EXPOSE 8080
WORKDIR /opt/service

CMD [ "./micro"]
