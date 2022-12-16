FROM alpine:latest

RUN apk update && apk upgrade && apk add jq && rm -rf /var/cache/apk/*

WORKDIR /jq

ENTRYPOINT ["jq","-r"]


# docker build . -t yudady/jq
# docker run -it -v ${PWD}:/jq yudady/jq /bin/sh
# docker run -it -v ${PWD}:/jq yudady/jq '.workers' employee.json



