## systemd

``` shell
sudo systemctl start unspammer
```

``` shell
sudo systemctl stop unspammer
```

## ACME Testing

``` shell
sudo ./unspammer-linux-amd64 -ca nap -blacklist default.bl -acme ephemelier.com -email 'mtr@iki.fi' -dry-run
```
