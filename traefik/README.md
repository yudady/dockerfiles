# DNS泛域名解析應用 
## [Dead simple wildcard DNS for any IP Address](https://nip.io/) 是一個免費的域名解析服務


> xxxxx.114.32.146.154.nip.io => 114.32.146.154


```shell
tommy@tommy-msi:~$ dig www.114.32.146.154.nip.io

; <<>> DiG 9.16.1-Ubuntu <<>> www.114.32.146.154.nip.io
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 64857
;; flags: qr rd ad; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;www.114.32.146.154.nip.io.     IN      A

;; ANSWER SECTION:
www.114.32.146.154.nip.io. 0    IN      A       114.32.146.154

;; Query time: 300 msec
;; SERVER: 172.21.16.1#53(172.21.16.1)
;; WHEN: Sat Dec 17 12:51:41 CST 2022
;; MSG SIZE  rcvd: 84
```



## [sslip.io](https://sslip.io/)

> https://114.32.146.154.sslip.io

```shell
tommy@tommy-msi:~$ dig www.114.32.146.154.sslip.io

; <<>> DiG 9.16.1-Ubuntu <<>> www.114.32.146.154.sslip.io
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 39474
;; flags: qr rd ad; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;www.114.32.146.154.sslip.io.   IN      A

;; ANSWER SECTION:
www.114.32.146.154.sslip.io. 0  IN      A       114.32.146.154

;; Query time: 190 msec
;; SERVER: 172.21.16.1#53(172.21.16.1)
;; WHEN: Sat Dec 17 12:51:10 CST 2022
;; MSG SIZE  rcvd: 88
```

## [traefik.me](https://traefik.me/)


```shell
tommy@tommy-msi:~$ dig www.114.32.146.154.traefik.me

; <<>> DiG 9.16.1-Ubuntu <<>> www.114.32.146.154.traefik.me
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 64230
;; flags: qr rd ad; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;www.114.32.146.154.traefik.me. IN      A

;; ANSWER SECTION:
www.114.32.146.154.traefik.me. 0 IN     A       114.32.146.154

;; Query time: 10 msec
;; SERVER: 172.21.16.1#53(172.21.16.1)
;; WHEN: Sat Dec 17 12:50:07 CST 2022
;; MSG SIZE  rcvd: 92
```



## ============================================================================================
## ============================================================================================
## ============================================================================================
## ============================================================================================

## [local.gd](https://local.gd/) : 127.0.0.1

```shell
tommy@tommy-msi:~$ dig startup.local.gd

; <<>> DiG 9.16.1-Ubuntu <<>> startup.local.gd
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 52449
;; flags: qr rd ad; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;startup.local.gd.              IN      A

;; ANSWER SECTION:
startup.local.gd.       0       IN      A       127.0.0.1

;; Query time: 0 msec
;; SERVER: 172.21.16.1#53(172.21.16.1)
;; WHEN: Sat Dec 17 00:32:30 CST 2022
;; MSG SIZE  rcvd: 66



===================================

$ dig startup.local.gd
startup.local.gd.   	86400	IN	A	127.0.0.1

$ dig www.startup.local.gd
www.startup.local.gd.	86400	IN	A	127.0.0.1

$ dig my.project.company.local.gd
my.project.company.local.gd.	86400	IN	A	127.0.0.1

$ dig alderaan.local.gd
alderaan.local.gd.  	86400	IN	A	127.0.0.1
```





## lvh.me : 127.0.0.1

```shell
tommy@tommy-msi:~$ dig example.lvh.me

; <<>> DiG 9.16.1-Ubuntu <<>> example.lvh.me
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 47377
;; flags: qr rd ad; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;example.lvh.me.                        IN      A

;; ANSWER SECTION:
example.lvh.me.         0       IN      A       127.0.0.1

;; Query time: 0 msec
;; SERVER: 172.21.16.1#53(172.21.16.1)
;; WHEN: Sat Dec 17 00:34:46 CST 2022
;; MSG SIZE  rcvd: 62
```

## ngrok


## other
- https://github.com/traefik/traefik/
- https://doc.traefik.io/traefik/providers/docker/
- [[Unobtrusive local development with traefik2, docker and letsencrypt | Codementor]] [Unobtrusive local development with traefik2, docker and letsencrypt | Codementor](https://www.codementor.io/@slavko/unobtrusive-local-development-with-traefik2-docker-and-letsencrypt-15qw1ypoi8)
- [[acme.sh 自動化申請和更新 Let's Encrypt 萬用 SSL 憑證教學 | KJie Notes]] [acme.sh 自動化申請和更新 Let's Encrypt 萬用 SSL 憑證教學 | KJie Notes](https://www.kjnotes.com/devtools/103)
