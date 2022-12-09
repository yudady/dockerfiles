$TTL 38400
@ IN SOA ns.yudady.com. admin.yudady.com. (
1       ;Serial
600     ;Refresh
300     ;Retry
60480   ;Expire
600 )   ;Negative Cache TTL

@       IN      NS      ns.yudady.com.
@       IN      NS      ns2.yudady.com.
