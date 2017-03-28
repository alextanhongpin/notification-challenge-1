FROM alpine:latest

MANTAINER Alex Tan <alextan220990@gmail.com>

WORKDIR "/opt"

ADD .docker_build/notification-challenge /opt/bin/notification-challenge

CMD ["/opt/bin/notification-challenge"]