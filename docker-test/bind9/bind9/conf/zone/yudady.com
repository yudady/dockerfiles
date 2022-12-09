$TTL    3600

@       IN      SOA     ns1.yudady.com. root.yudady.com. (
        2022120901      ; serial
        3600            ; refresh after 2 hours
        3600            ; retry after 1 hour
        604800          ; expire after 1 week
        3600  )         ; minimum TTL of 1 day

;
; Primary nameserver
        IN      NS      ns1.yudady.com.

;
; Define A records (forward lookups)
; @     IN      NS      ns1.yudady.com.
ns1  IN A  127.0.0.1
app  IN A  127.0.0.1
db   IN A  127.0.0.1
test IN A  127.0.0.1