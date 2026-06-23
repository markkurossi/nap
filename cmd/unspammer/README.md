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

## OpenSSL

``` shell
openssl x509 -in ee-cert.pem -text -noout
```

## Creating an EE Certificate for a P-256 Public Key

``` shell
/unspammer -ca nap -create-ee "localhost" -pubkey 045025d9e221d24dcf972634fef75cea0474e98a7db906506a13c5b0aeada05d2476b35bfba44c849fcb55db6e556a1113ac8413e81197175c135feeb42bee8ded
```
