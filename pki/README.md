# Configuring Trusted CA Certificates

## macOS

Use Keychain Access to set an explicit trust.

1. Open Keychain Access
2. Click on the System keychain in the list of keychains on the left
3. In the list on the right, find the certificate for the server and select it
4. Press the i (get info) button
5. Reveal the Trust arrow and change the "When using this certificate"
   to Always Trust

# Certificates

## CA Certificate

```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            4b:a8:52:93:f7:9a:2f:a2:73:06:4b:a8:04:8d:75:d0
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Internet Security Research Group, CN = ISRG Root X1
        Validity
            Not Before: Mar 13 00:00:00 2024 GMT
            Not After : Mar 12 23:59:59 2027 GMT
        Subject: C = US, O = Let's Encrypt, CN = R10
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:cf:57:e5:e6:c4:54:12:ed:b4:47:fe:c9:27:58:
                    76:46:50:28:8c:1d:3e:88:df:05:9d:d5:b5:18:29:
                    bd:dd:b5:5a:bf:fa:f6:ce:a3:be:af:00:21:4b:62:
                    5a:5a:3c:01:2f:c5:58:03:f6:89:ff:8e:11:43:eb:
                    c1:b5:e0:14:07:96:8f:6f:1f:d7:e7:ba:81:39:09:
                    75:65:b7:c2:af:18:5b:37:26:28:e7:a3:f4:07:2b:
                    6d:1a:ff:ab:58:bc:95:ae:40:ff:e9:cb:57:c4:b5:
                    5b:7f:78:0d:18:61:bc:17:e7:54:c6:bb:49:91:cd:
                    6e:18:d1:80:85:ee:a6:65:36:bc:74:ea:bc:50:4c:
                    ea:fc:21:f3:38:16:93:94:ba:b0:d3:6b:38:06:cd:
                    16:12:7a:ca:52:75:c8:ad:76:b2:c2:9c:5d:98:45:
                    5c:6f:61:7b:c6:2d:ee:3c:13:52:86:01:d9:57:e6:
                    38:1c:df:8d:b5:1f:92:91:9a:e7:4a:1c:cc:45:a8:
                    72:55:f0:b0:e6:a3:07:ec:fd:a7:1b:66:9e:3f:48:
                    8b:71:84:71:58:c9:3a:fa:ef:5e:f2:5b:44:2b:3c:
                    74:e7:8f:b2:47:c1:07:6a:cd:9a:b7:0d:96:f7:12:
                    81:26:51:54:0a:ec:61:f6:f7:f5:e2:f2:8a:c8:95:
                    0d:8d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage:
                TLS Web Client Authentication, TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier:
                BB:BC:C3:47:A5:E4:BC:A9:C6:C3:A4:72:0C:10:8D:A2:35:E1:C8:E8
            X509v3 Authority Key Identifier:
                keyid:79:B4:59:E6:7B:B6:E5:E4:01:73:80:08:88:C8:1A:58:F6:E9:9B:6E

            Authority Information Access:
                CA Issuers - URI:http://x1.i.lencr.org/

            X509v3 Certificate Policies:
                Policy: 2.23.140.1.2.1

            X509v3 CRL Distribution Points:

                Full Name:
                  URI:http://x1.c.lencr.org/

    Signature Algorithm: sha256WithRSAEncryption
         92:b1:e7:41:37:eb:79:9d:81:e6:cd:e2:25:e1:3a:20:e9:90:
         44:95:a3:81:5c:cf:c3:5d:fd:bd:a0:70:d5:b1:96:28:22:0b:
         d2:f2:28:cf:0c:e7:d4:e6:43:8c:24:22:1d:c1:42:92:d1:09:
         af:9f:4b:f4:c8:70:4f:20:16:b1:5a:dd:01:f6:1f:f8:1f:61:
         6b:14:27:b0:72:8d:63:ae:ee:e2:ce:4b:cf:37:dd:bb:a3:d4:
         cd:e7:ad:50:ad:bd:bf:e3:ec:3e:62:36:70:99:31:a7:e8:8d:
         dd:ea:62:e2:12:ae:f5:9c:d4:3d:2c:0c:aa:d0:9c:79:be:ea:
         3d:5c:44:6e:96:31:63:5a:7d:d6:7e:4f:24:a0:4b:05:7f:5e:
         6f:d2:d4:ea:5f:33:4b:13:d6:57:b6:ca:de:51:b8:5d:a3:09:
         82:74:fd:c7:78:9e:b3:b9:ac:16:da:4a:2b:96:c3:b6:8b:62:
         8f:f9:74:19:a2:9e:03:de:e9:6f:9b:b0:0f:d2:a0:5a:f6:85:
         5c:c2:04:b7:c8:d5:4e:32:c4:bf:04:5d:bc:29:f6:f7:81:8f:
         0c:5d:3c:53:c9:40:90:8b:fb:b6:08:65:b9:a4:21:d5:09:e5:
         13:84:84:37:82:ce:10:28:fc:76:c2:06:25:7a:46:52:4d:da:
         53:72:a4:27:3f:62:70:ac:be:69:48:00:fb:67:0f:db:5b:a1:
         e8:d7:03:21:2d:d7:c9:f6:99:42:39:83:43:df:77:0a:12:08:
         f1:25:d6:ba:94:19:54:18:88:a5:c5:8e:e1:1a:99:93:79:6b:
         ec:1c:f9:31:40:b0:cc:32:00:df:9f:5e:e7:b4:92:ab:90:82:
         91:8d:0d:e0:1e:95:ba:59:3b:2e:4b:5f:c2:b7:46:35:52:39:
         06:c0:bd:aa:ac:52:c1:22:a0:44:97:99:f7:0c:a0:21:a7:a1:
         6c:71:47:16:17:01:68:c0:ca:a6:26:65:04:7c:b3:ae:c9:e7:
         94:55:c2:6f:9b:3c:1c:a9:f9:2e:c5:20:1a:f0:76:e0:be:ec:
         18:d6:4f:d8:25:fb:76:11:e8:bf:e6:21:0f:e8:e8:cc:b5:b6:
         a7:d5:b8:f7:9f:41:cf:61:22:46:6a:83:b6:68:97:2e:7c:ea:
         4e:95:db:23:eb:2e:c8:2b:28:84:a4:60:e9:49:f4:44:2e:3b:
         f9:ca:62:57:01:e2:5d:90:16:f9:c9:fc:7a:23:48:8e:a6:d5:
         81:72:f1:28:fa:5d:ce:fb:ed:4e:73:8f:94:2e:d2:41:94:98:
         99:db:a7:af:70:5f:f5:be:fb:02:20:bf:66:27:6c:b4:ad:fa:
         75:12:0b:2b:3e:ce:03:9e
```

## CRL

```
Certificate Revocation List (CRL):
        Version 2 (0x1)
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Internet Security Research Group, CN = ISRG Root X1
        Last Update: Feb  5 00:00:00 2024 GMT
        Next Update: Jan  4 23:59:59 2025 GMT
        CRL extensions:
            X509v3 Authority Key Identifier:
                79:B4:59:E6:7B:B6:E5:E4:01:73:80:08:88:C8:1A:58:F6:E9:9B:6E
            X509v3 CRL Number:
                104
No Revoked Certificates.
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        59:26:d6:a5:01:52:f5:e1:20:f8:e7:5d:6d:28:5c:a1:6f:39:
        1e:ee:92:a8:4d:07:f9:a4:65:af:37:db:f8:a9:4f:df:a1:b4:
        96:e5:61:1a:84:3c:03:66:0d:4f:6c:33:1c:97:b1:e5:33:9e:
        4a:d9:1e:88:c7:42:8e:fd:36:21:24:e6:a0:87:b0:d2:c4:34:
        41:7c:d3:68:9c:50:f2:a5:a6:09:8c:8b:c2:62:63:dc:26:a4:
        12:ae:c3:81:65:c0:44:2a:35:01:49:b2:cd:59:6a:e7:5d:f1:
        1f:63:84:aa:a1:53:3a:5f:7f:f3:9a:ed:42:4b:64:21:52:fa:
        9d:e9:b9:af:bb:c7:5c:e4:78:3c:47:f3:be:16:78:c4:23:63:
        c1:6a:e9:8e:65:31:9b:00:24:0f:91:20:98:1f:47:55:ca:ab:
        6a:72:ad:ac:b9:c0:f9:3c:4f:1b:46:58:d8:50:8f:e7:13:7b:
        ff:fb:5f:8b:c1:ba:01:97:37:77:34:20:a8:d5:4d:b0:9c:f2:
        8f:6d:22:b2:dd:5f:05:b6:2c:de:99:a2:b6:ea:ef:59:64:d5:
        c1:b0:7f:80:45:cc:68:87:7c:63:eb:63:07:f1:49:1a:8f:38:
        e4:05:2d:7c:e0:42:98:ae:07:07:8b:f7:9c:3b:a9:09:70:bf:
        8f:52:d3:30:ea:df:42:67:88:6b:d2:de:ab:3d:28:a4:7a:d7:
        7d:bb:82:6f:6a:10:96:01:4a:3f:81:d7:e1:e3:5a:91:58:9e:
        2d:f2:f9:5e:58:cf:ce:63:a3:bd:46:8f:0c:97:6c:4f:97:d5:
        48:de:9c:cb:57:c3:9a:c6:a2:92:78:e6:05:3d:d5:4e:14:d9:
        f8:f4:09:9e:d2:fe:13:38:5b:e9:af:0a:ec:92:e7:bf:ee:5a:
        33:48:ee:31:82:d7:6f:0b:cd:ec:aa:db:66:9f:d8:a1:63:20:
        57:7b:76:aa:d0:d6:a5:1e:c9:44:45:dd:3c:18:bd:6f:05:b8:
        19:58:a0:e9:c5:8a:58:70:3b:e4:22:bf:0d:c8:a3:e0:53:a6:
        7f:2b:a6:39:14:ad:2b:0d:b7:4a:46:d2:78:21:67:6b:25:33:
        23:d6:ab:17:80:bb:66:22:ec:ee:6d:b1:e5:01:ae:4e:5b:5c:
        3b:35:54:3e:5a:94:51:f5:81:eb:cb:10:ca:d6:39:7e:17:ae:
        f0:4d:25:81:64:cd:b6:06:09:ea:75:eb:0e:06:e5:a4:c0:1e:
        0e:24:9f:33:bf:fd:1f:12:48:57:60:e1:a4:e8:aa:b2:30:e9:
        ec:e0:52:76:44:4e:bd:42:69:69:b5:de:51:ef:84:a4:16:19:
        49:a2:2b:d2:3d:62:b4:6e
```

## EE Certificate

```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            04:70:84:fa:3e:af:76:dc:26:92:d4:53:42:b6:1d:25:6a:88
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Let's Encrypt, CN = R10
        Validity
            Not Before: Jun 28 02:32:15 2024 GMT
            Not After : Sep 26 02:32:14 2024 GMT
        Subject: CN = www.markkurossi.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:c0:84:1b:b0:d1:ef:af:e3:1c:b2:c0:ca:8f:aa:
                    9d:e5:1e:a0:7a:d2:91:45:09:c7:56:c4:97:a8:55:
                    33:ce:55:1d:be:e8:44:da:48:5f:fa:92:d4:04:00:
                    bb:8c:e7:4e:e5:c4:0c:08:96:f3:5d:b4:41:e0:50:
                    94:ec:b2:7f:00:e9:8c:76:0b:04:ef:cd:13:9f:1c:
                    91:fa:e9:53:c3:78:a4:ac:fa:0b:a8:da:35:b5:bb:
                    4a:96:83:56:23:93:c0:e8:59:e6:f8:1c:2b:0e:fb:
                    a9:4a:70:01:85:df:a2:25:aa:00:82:13:d3:c0:a1:
                    79:9d:3a:09:0b:06:4c:3e:3b:9f:f1:53:ef:a2:6d:
                    b5:29:d2:50:33:a4:50:41:95:ce:21:65:e9:c4:41:
                    93:23:5d:83:83:9d:78:2b:8d:c7:c2:08:3c:39:52:
                    3f:41:7a:92:08:0b:d2:85:b7:90:bf:bd:cc:84:4b:
                    1b:50:b8:97:a5:5a:51:be:6d:fd:13:00:1a:d0:01:
                    ad:53:74:5e:9b:61:b1:2d:ff:47:c1:d8:60:25:49:
                    ed:d7:e3:41:19:55:77:2e:1b:fc:29:3d:49:37:82:
                    11:00:29:d5:89:cc:51:6c:5a:49:cc:70:9a:58:b0:
                    60:44:c1:18:e4:43:ea:5b:eb:b7:75:65:8f:57:ed:
                    c0:2f
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage:
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier:
                4E:C4:80:2F:5A:38:EC:F1:86:F6:3A:7A:F7:47:13:1E:3B:97:6C:B3
            X509v3 Authority Key Identifier:
                BB:BC:C3:47:A5:E4:BC:A9:C6:C3:A4:72:0C:10:8D:A2:35:E1:C8:E8
            Authority Information Access:
                OCSP - URI:http://r10.o.lencr.org
                CA Issuers - URI:http://r10.i.lencr.org/
            X509v3 Subject Alternative Name:
                DNS:markkurossi.com, DNS:www.markkurossi.com
            X509v3 Certificate Policies:
                Policy: 2.23.140.1.2.1
            CT Precertificate SCTs:
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 48:B0:E3:6B:DA:A6:47:34:0F:E5:6A:02:FA:9D:30:EB:
                                1C:52:01:CB:56:DD:2C:81:D9:BB:BF:AB:39:D8:84:73
                    Timestamp : Jun 28 03:32:15.517 2024 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:44:02:20:69:7E:3F:A2:7F:B2:30:E8:C1:68:B4:4C:
                                6C:E4:F9:13:D6:AF:C6:E0:82:28:27:72:BF:5A:4A:63:
                                28:58:58:C4:02:20:4C:06:90:BB:F5:E7:01:41:4C:77:
                                9B:C1:0A:DE:50:96:31:22:F7:C2:6B:FB:3B:23:07:E8:
                                3D:A3:E9:BF:F3:8E
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 19:98:10:71:09:F0:D6:52:2E:30:80:D2:9E:3F:64:BB:
                                83:6E:28:CC:F9:0F:52:8E:EE:DF:CE:4A:3F:16:B4:CA
                    Timestamp : Jun 28 03:32:15.921 2024 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:21:00:84:54:2A:08:F8:27:69:DA:99:92:3C:
                                0A:57:43:4A:88:15:2D:87:83:ED:8C:02:74:49:4B:94:
                                68:5C:9B:49:37:02:20:11:8E:4A:BE:99:5F:19:6B:6B:
                                06:FF:B2:FE:CB:F8:8B:B9:D6:BD:60:65:17:1E:26:3B:
                                8B:C4:55:A0:EB:B2:24
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        ad:81:51:d5:6f:18:04:31:15:db:1a:52:a6:8c:89:ee:30:c4:
        22:8f:5c:dc:81:07:85:92:1d:d0:38:5a:a3:4a:49:1e:78:a8:
        0f:4a:25:73:c2:35:31:c5:97:1c:de:16:67:d6:13:32:e2:42:
        9e:be:e5:63:b9:11:e9:1b:02:61:41:f1:23:d3:de:b2:03:cc:
        92:4e:e3:f1:4d:bd:75:00:0b:64:ee:18:1f:9e:e4:4c:fb:90:
        6a:af:4b:1a:5e:35:05:68:15:2b:96:7c:93:cb:a2:5f:dc:0a:
        2f:0b:4f:93:07:0c:6b:a9:0c:60:d5:01:a8:67:87:6f:32:b8:
        bb:50:4f:ae:83:ac:e9:e9:2d:36:69:05:24:a8:f9:71:66:a6:
        2b:5f:47:5c:a7:ac:e2:1d:f9:58:ad:93:d9:16:3b:92:7d:f1:
        6b:f0:de:26:40:65:25:cb:48:43:c0:94:bc:d1:a3:55:53:15:
        09:51:ea:11:ba:7d:ef:bc:3d:d2:63:c6:21:b2:96:e3:96:f6:
        ab:a4:0b:f7:80:b7:36:05:b6:75:e9:0f:e7:a0:38:92:80:8e:
        81:5e:44:3a:98:46:b1:26:27:7f:cd:c6:4a:15:38:3f:24:ab:
        a9:f0:55:21:d1:8b:a2:28:4e:9e:11:29:65:2a:a7:e7:28:27:
        86:98:61:f0
```
