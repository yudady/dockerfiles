FROM alpine:latest

RUN apk add bind iputils bind-tools openrc

RUN mkdir -p /run/openrc && touch /run/openrc/softlevel
RUN rc-update -u named

# Volume ["/sys/fs/cgroup"]

EXPOSE 53/tcp
EXPOSE 53/udp
