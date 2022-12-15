# [docker alpine](https://alpinelinux.org/about/)


[github docker alpine](https://github.com/alpinelinux/docker-alpine)



```shell
docker pull alpine:3.17
```


```dockerfile
FROM alpine:3.17
RUN apk add --no-cache mysql-client
ENTRYPOINT ["mysql"]
```


## apk command
https://wangchujiang.com/linux-command/c/apk.html


## 刪除docker name is none
```shell
docker rmi -f $(docker images --filter "dangling=true" -q --no-trunc)

docker rmi $(docker images -f "dangling=true" -q)

PS C:\curl-cmd> docker system prune
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all dangling images
  - all dangling build cache

Are you sure you want to continue? [y/N] y
Deleted Containers:
fc6bc00814a03d33778d9ebea8188893b01eb8e0cbbf7b131bcfafe7b1159af4

```