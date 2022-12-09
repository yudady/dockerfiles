$TTL 38400
@ IN SOA ns.yudady.com. admin.yudady.com. (
1       ;Serial
600     ;Refresh
300     ;Retry
60480   ;Expire
600 )   ;Negative Cache TTL

@                   IN      MX  1   127.0.0.1
ns                  IN      A       127.0.0.1
www                 IN      A       114.32.146.154