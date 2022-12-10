# docker bind9

## github 
- https://github.com/sameersbn/docker-bind
- web login https://localhost:10000   =>   root / password
- docker exec -it bind bash

```dockerfile
docker run --name bind -d --restart=always \
  --publish 53:53/tcp --publish 53:53/udp --publish 10000:10000/tcp \
  --volume ${PWD}/data:/data \
  sameersbn/bind:9.16.1-20200524


docker run --name bind -d --restart=always \
  --publish 53:53/tcp --publish 53:53/udp --publish 10000:10000/tcp \
  --volume ${PWD}/data:/data \
  --volume ${PWD}/conf/bind/etc/named.conf.local:/data/bind/etc/named.conf.local \
  --volume ${PWD}/conf/bind/lib/named.conf.local:/data/bind/lib/getcharzp.cn.hosts \
  sameersbn/bind:9.16.1-20200524


docker exec -it bind bash
```

### STOP
```shell
docker stop bind
docker rm -v bind
```




```shell
apt install bind9-utils
apt install iputils-ping -y
```

## info

```shell
cat /data/bind/etc/named.conf
cat /data/bind/etc/named.conf.local
```



## ref
https://www.bilibili.com/video/BV1cT4y1o7mL/?spm_id_from=333.337.search-card.all.click&vd_source=6bd04a20c72eb5cca642210346af7081
https://www.cnblogs.com/GetcharZp/p/15391326.html
https://bind9.readthedocs.io/en/latest/reference.html
https://www.bilibili.com/video/BV1JW4y1J7cw/?p=2&spm_id_from=pageDriver&vd_source=6bd04a20c72eb5cca642210346af7081
