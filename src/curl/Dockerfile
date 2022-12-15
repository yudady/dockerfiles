FROM alpine:latest

RUN apk update && apk upgrade && apk add curl && apk add curl-doc && rm -rf /var/cache/apk/*

ENTRYPOINT ["curl","-s","-i"]
CMD ["http://www.google.com"]

# docker build . -t yudady/curl
# docker run yudady/curl
# docker run yudady/curl http://www.yahoo.com


