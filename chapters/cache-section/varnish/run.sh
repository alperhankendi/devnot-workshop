/usr/local/sbin/varnishd \
-f $PWD/default.vcl \
-s malloc,1G -T 127.0.0.1:2000 -a 0.0.0.0:8080