Retrieve server certificate with `openssl s_client`.
```sh
$ openssl s_client -connect www.google.com:443
CONNECTED(00000003)
depth=2 C = US, O = Google Trust Services LLC, CN = GTS Root R1
verify return:1
depth=1 C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
verify return:1
depth=0 CN = www.google.com
verify return:1
---
Certificate chain
 0 s:CN = www.google.com
   i:C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
 1 s:C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
   i:C = US, O = Google Trust Services LLC, CN = GTS Root R1
 2 s:C = US, O = Google Trust Services LLC, CN = GTS Root R1
   i:C = BE, O = GlobalSign nv-sa, OU = Root CA, CN = GlobalSign Root CA
---
Server certificate
-----BEGIN CERTIFICATE-----
MIIEiTCCA3GgAwIBAgIRAJe1AQS4zHSACSQRxVDlVq0wDQYJKoZIhvcNAQELBQAw
RjELMAkGA1UEBhMCVVMxIjAgBgNVBAoTGUdvb2dsZSBUcnVzdCBTZXJ2aWNlcyBM
TEMxEzARBgNVBAMTCkdUUyBDQSAxQzMwHhcNMjMxMDA5MDgxMTI2WhcNMjQwMTAx
MDgxMTI1WjAZMRcwFQYDVQQDEw53d3cuZ29vZ2xlLmNvbTBZMBMGByqGSM49AgEG
CCqGSM49AwEHA0IABFW7lhdHYCFrndwECat/kpAbVyofcHLsGPcC6DLOMqdQmNV5
hEiWeEKTcd29TF9rC6Z0g4lo72To58ZeO5YWyg+jggJoMIICZDAOBgNVHQ8BAf8E
BAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQUAqcP0O7KeZczpW7niF8/vsiwV2YwHwYDVR0jBBgwFoAUinR/r4XN7pXNPZzQ
4kYU83E1HScwagYIKwYBBQUHAQEEXjBcMCcGCCsGAQUFBzABhhtodHRwOi8vb2Nz
cC5wa2kuZ29vZy9ndHMxYzMwMQYIKwYBBQUHMAKGJWh0dHA6Ly9wa2kuZ29vZy9y
ZXBvL2NlcnRzL2d0czFjMy5kZXIwGQYDVR0RBBIwEIIOd3d3Lmdvb2dsZS5jb20w
IQYDVR0gBBowGDAIBgZngQwBAgEwDAYKKwYBBAHWeQIFAzA8BgNVHR8ENTAzMDGg
L6AthitodHRwOi8vY3Jscy5wa2kuZ29vZy9ndHMxYzMvbW9WRGZJU2lhMmsuY3Js
MIIBBQYKKwYBBAHWeQIEAgSB9gSB8wDxAHcA2ra/az+1tiKfm8K7XGvocJFxbLtR
hIU0vaQ9MEjX+6sAAAGLE7SmVgAABAMASDBGAiEA/w5JHHAqyDFgAxT+L032qeWj
hzHxrIJxy5vdGUbI+o0CIQDkagdSK0yMLILOvQpkBq4xi+H2N3UVhQXaL3HJgLSw
lQB2AHb/iD8KtvuVUcJhzPWHujS0pM27KdxoQgqf5mdMWjp0AAABixO0plEAAAQD
AEcwRQIgBtL8IGSWqxiazOG4DPc1E+b1fmBWg0R881J6NwDI6lQCIQCwf5GmVYOa
wgqmBf6YN6YAu9i6n8e29wF+xZ/177LbczANBgkqhkiG9w0BAQsFAAOCAQEAdrWE
U9eXgwuYnsJqD9dozsWbkUPPpW+22RzOecX7eqYsxj+c+D9ZDSmOplUG7PmuJzSs
9bquP7YGJeROp2LgxrOfuJb6dKM+FwQ9HnD16D+zw+K0HIkOqydyHyz9ZDT8znQs
02SXSTbnssa/t0dXApLS1e4buOc7uVXsEcoEqfzVD9+AIJ3bsTlIVzb6NVBJcUj3
wvA2x8nAo0EWt0IOh5ZS+FF2s5uwGTzC3dB/vjT9sy7X0Rp+QZNr8Zjwn5KVt4B5
AVavn/+tKT5JwIqltYvj7/p1SyPpY++zT8YXSnyzk7bzRqBwmV1lmfX1FXZQoNa0
U5dXPL22Ilrxb4fV+g==
-----END CERTIFICATE-----
subject=CN = www.google.com

issuer=C = US, O = Google Trust Services LLC, CN = GTS CA 1C3

---
No client certificate CA names sent
Peer signing digest: SHA256
Peer signature type: ECDSA
Server Temp Key: X25519, 253 bits
---
SSL handshake has read 4295 bytes and written 396 bytes
Verification: OK
---
New, TLSv1.3, Cipher is TLS_AES_256_GCM_SHA384
Server public key is 256 bit
Secure Renegotiation IS NOT supported
Compression: NONE
Expansion: NONE
No ALPN negotiated
Early data was not sent
Verify return code: 0 (ok)
---
```

Print server certificate.
```sh
$ openssl x509 -text -noout
-----BEGIN CERTIFICATE-----
MIIEiTCCA3GgAwIBAgIRAJe1AQS4zHSACSQRxVDlVq0wDQYJKoZIhvcNAQELBQAw
RjELMAkGA1UEBhMCVVMxIjAgBgNVBAoTGUdvb2dsZSBUcnVzdCBTZXJ2aWNlcyBM
TEMxEzARBgNVBAMTCkdUUyBDQSAxQzMwHhcNMjMxMDA5MDgxMTI2WhcNMjQwMTAx
MDgxMTI1WjAZMRcwFQYDVQQDEw53d3cuZ29vZ2xlLmNvbTBZMBMGByqGSM49AgEG
CCqGSM49AwEHA0IABFW7lhdHYCFrndwECat/kpAbVyofcHLsGPcC6DLOMqdQmNV5
hEiWeEKTcd29TF9rC6Z0g4lo72To58ZeO5YWyg+jggJoMIICZDAOBgNVHQ8BAf8E
BAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQUAqcP0O7KeZczpW7niF8/vsiwV2YwHwYDVR0jBBgwFoAUinR/r4XN7pXNPZzQ
4kYU83E1HScwagYIKwYBBQUHAQEEXjBcMCcGCCsGAQUFBzABhhtodHRwOi8vb2Nz
cC5wa2kuZ29vZy9ndHMxYzMwMQYIKwYBBQUHMAKGJWh0dHA6Ly9wa2kuZ29vZy9y
ZXBvL2NlcnRzL2d0czFjMy5kZXIwGQYDVR0RBBIwEIIOd3d3Lmdvb2dsZS5jb20w
IQYDVR0gBBowGDAIBgZngQwBAgEwDAYKKwYBBAHWeQIFAzA8BgNVHR8ENTAzMDGg
L6AthitodHRwOi8vY3Jscy5wa2kuZ29vZy9ndHMxYzMvbW9WRGZJU2lhMmsuY3Js
MIIBBQYKKwYBBAHWeQIEAgSB9gSB8wDxAHcA2ra/az+1tiKfm8K7XGvocJFxbLtR
hIU0vaQ9MEjX+6sAAAGLE7SmVgAABAMASDBGAiEA/w5JHHAqyDFgAxT+L032qeWj
hzHxrIJxy5vdGUbI+o0CIQDkagdSK0yMLILOvQpkBq4xi+H2N3UVhQXaL3HJgLSw
lQB2AHb/iD8KtvuVUcJhzPWHujS0pM27KdxoQgqf5mdMWjp0AAABixO0plEAAAQD
AEcwRQIgBtL8IGSWqxiazOG4DPc1E+b1fmBWg0R881J6NwDI6lQCIQCwf5GmVYOa
wgqmBf6YN6YAu9i6n8e29wF+xZ/177LbczANBgkqhkiG9w0BAQsFAAOCAQEAdrWE
U9eXgwuYnsJqD9dozsWbkUPPpW+22RzOecX7eqYsxj+c+D9ZDSmOplUG7PmuJzSs
9bquP7YGJeROp2LgxrOfuJb6dKM+FwQ9HnD16D+zw+K0HIkOqydyHyz9ZDT8znQs
02SXSTbnssa/t0dXApLS1e4buOc7uVXsEcoEqfzVD9+AIJ3bsTlIVzb6NVBJcUj3
wvA2x8nAo0EWt0IOh5ZS+FF2s5uwGTzC3dB/vjT9sy7X0Rp+QZNr8Zjwn5KVt4B5
AVavn/+tKT5JwIqltYvj7/p1SyPpY++zT8YXSnyzk7bzRqBwmV1lmfX1FXZQoNa0
U5dXPL22Ilrxb4fV+g==
-----END CERTIFICATE-----
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            97:b5:01:04:b8:cc:74:80:09:24:11:c5:50:e5:56:ad
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
        Validity
            Not Before: Oct  9 08:11:26 2023 GMT
            Not After : Jan  1 08:11:25 2024 GMT
        Subject: CN = www.google.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:55:bb:96:17:47:60:21:6b:9d:dc:04:09:ab:7f:
                    92:90:1b:57:2a:1f:70:72:ec:18:f7:02:e8:32:ce:
                    32:a7:50:98:d5:79:84:48:96:78:42:93:71:dd:bd:
                    4c:5f:6b:0b:a6:74:83:89:68:ef:64:e8:e7:c6:5e:
                    3b:96:16:ca:0f
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier: 
                02:A7:0F:D0:EE:CA:79:97:33:A5:6E:E7:88:5F:3F:BE:C8:B0:57:66
            X509v3 Authority Key Identifier: 
                keyid:8A:74:7F:AF:85:CD:EE:95:CD:3D:9C:D0:E2:46:14:F3:71:35:1D:27

            Authority Information Access: 
                OCSP - URI:http://ocsp.pki.goog/gts1c3
                CA Issuers - URI:http://pki.goog/repo/certs/gts1c3.der

            X509v3 Subject Alternative Name: 
                DNS:www.google.com
            X509v3 Certificate Policies: 
                Policy: 2.23.140.1.2.1
                Policy: 1.3.6.1.4.1.11129.2.5.3

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crls.pki.goog/gts1c3/moVDfISia2k.crl

            CT Precertificate SCTs: 
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : DA:B6:BF:6B:3F:B5:B6:22:9F:9B:C2:BB:5C:6B:E8:70:
                                91:71:6C:BB:51:84:85:34:BD:A4:3D:30:48:D7:FB:AB
                    Timestamp : Oct  9 09:11:28.086 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:46:02:21:00:FF:0E:49:1C:70:2A:C8:31:60:03:14:
                                FE:2F:4D:F6:A9:E5:A3:87:31:F1:AC:82:71:CB:9B:DD:
                                19:46:C8:FA:8D:02:21:00:E4:6A:07:52:2B:4C:8C:2C:
                                82:CE:BD:0A:64:06:AE:31:8B:E1:F6:37:75:15:85:05:
                                DA:2F:71:C9:80:B4:B0:95
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 76:FF:88:3F:0A:B6:FB:95:51:C2:61:CC:F5:87:BA:34:
                                B4:A4:CD:BB:29:DC:68:42:0A:9F:E6:67:4C:5A:3A:74
                    Timestamp : Oct  9 09:11:28.081 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:20:06:D2:FC:20:64:96:AB:18:9A:CC:E1:B8:
                                0C:F7:35:13:E6:F5:7E:60:56:83:44:7C:F3:52:7A:37:
                                00:C8:EA:54:02:21:00:B0:7F:91:A6:55:83:9A:C2:0A:
                                A6:05:FE:98:37:A6:00:BB:D8:BA:9F:C7:B6:F7:01:7E:
                                C5:9F:F5:EF:B2:DB:73
    Signature Algorithm: sha256WithRSAEncryption
         76:b5:84:53:d7:97:83:0b:98:9e:c2:6a:0f:d7:68:ce:c5:9b:
         91:43:cf:a5:6f:b6:d9:1c:ce:79:c5:fb:7a:a6:2c:c6:3f:9c:
         f8:3f:59:0d:29:8e:a6:55:06:ec:f9:ae:27:34:ac:f5:ba:ae:
         3f:b6:06:25:e4:4e:a7:62:e0:c6:b3:9f:b8:96:fa:74:a3:3e:
         17:04:3d:1e:70:f5:e8:3f:b3:c3:e2:b4:1c:89:0e:ab:27:72:
         1f:2c:fd:64:34:fc:ce:74:2c:d3:64:97:49:36:e7:b2:c6:bf:
         b7:47:57:02:92:d2:d5:ee:1b:b8:e7:3b:b9:55:ec:11:ca:04:
         a9:fc:d5:0f:df:80:20:9d:db:b1:39:48:57:36:fa:35:50:49:
         71:48:f7:c2:f0:36:c7:c9:c0:a3:41:16:b7:42:0e:87:96:52:
         f8:51:76:b3:9b:b0:19:3c:c2:dd:d0:7f:be:34:fd:b3:2e:d7:
         d1:1a:7e:41:93:6b:f1:98:f0:9f:92:95:b7:80:79:01:56:af:
         9f:ff:ad:29:3e:49:c0:8a:a5:b5:8b:e3:ef:fa:75:4b:23:e9:
         63:ef:b3:4f:c6:17:4a:7c:b3:93:b6:f3:46:a0:70:99:5d:65:
         99:f5:f5:15:76:50:a0:d6:b4:53:97:57:3c:bd:b6:22:5a:f1:
         6f:87:d5:fa
```

Grab GTS (intermediate CA) public key.
```sh
$ openssl x509 -text -noout
-----BEGIN CERTIFICATE-----
MIIFljCCA36gAwIBAgINAgO8U1lrNMcY9QFQZjANBgkqhkiG9w0BAQsFADBHMQsw
CQYDVQQGEwJVUzEiMCAGA1UEChMZR29vZ2xlIFRydXN0IFNlcnZpY2VzIExMQzEU
MBIGA1UEAxMLR1RTIFJvb3QgUjEwHhcNMjAwODEzMDAwMDQyWhcNMjcwOTMwMDAw
MDQyWjBGMQswCQYDVQQGEwJVUzEiMCAGA1UEChMZR29vZ2xlIFRydXN0IFNlcnZp
Y2VzIExMQzETMBEGA1UEAxMKR1RTIENBIDFDMzCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAPWI3+dijB43+DdCkH9sh9D7ZYIl/ejLa6T/belaI+KZ9hzp
kgOZE3wJCor6QtZeViSqejOEH9Hpabu5dOxXTGZok3c3VVP+ORBNtzS7XyV3NzsX
lOo85Z3VvMO0Q+sup0fvsEQRY9i0QYXdQTBIkxu/t/bgRQIh4JZCF8/ZK2VWNAcm
BA2o/X3KLu/qSHw3TT8An4Pf73WELnlXXPxXbhqW//yMmqaZviXZf5YsBvcRKgKA
gOtjGDxQSYflispfGStZloEAoPtR28p3CwvJlk/vcEnHXG0g/Zm0tOLKLnf9LdwL
tmsTDIwZKxeWmLnwi/agJ7u2441Rj72ux5uxiZ0CAwEAAaOCAYAwggF8MA4GA1Ud
DwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwEgYDVR0T
AQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUinR/r4XN7pXNPZzQ4kYU83E1HScwHwYD
VR0jBBgwFoAU5K8rJnEaK0gnhS9SZizv8IkTcT4waAYIKwYBBQUHAQEEXDBaMCYG
CCsGAQUFBzABhhpodHRwOi8vb2NzcC5wa2kuZ29vZy9ndHNyMTAwBggrBgEFBQcw
AoYkaHR0cDovL3BraS5nb29nL3JlcG8vY2VydHMvZ3RzcjEuZGVyMDQGA1UdHwQt
MCswKaAnoCWGI2h0dHA6Ly9jcmwucGtpLmdvb2cvZ3RzcjEvZ3RzcjEuY3JsMFcG
A1UdIARQME4wOAYKKwYBBAHWeQIFAzAqMCgGCCsGAQUFBwIBFhxodHRwczovL3Br
aS5nb29nL3JlcG9zaXRvcnkvMAgGBmeBDAECATAIBgZngQwBAgIwDQYJKoZIhvcN
AQELBQADggIBAIl9rCBcDDy+mqhXlRu0rvqrpXJxtDaV/d9AEQNMwkYUuxQkq/BQ
cSLbrcRuf8/xam/IgxvYzolfh2yHuKkMo5uhYpSTld9brmYZCwKWnvy15xBpPnrL
RklfRuFBsdeYTWU0AIAaP0+fbH9JAIFTQaSSIYKCGvGjRFsqUBITTcFTNvNCCK9U
+o53UxtkOCcXCb1YyRt8OS1b887U7ZfbFAO/CVMkH8IMBHmYJvJh8VNS/UKMG2Yr
PxWhu//2m+OBmgEGcYk1KCTd4b3rGS3hSMs9WYNRtHTGnXzGsYZbr8w0xNPM1IER
lQCh9BIiAfq0g3GvjLeMcySsN1PCAJA/Ef5c7TaUEDu9Ka7ixzpiO2xj2YC/WXGs
Yye5TBeg2vZzFb8q3o/zpWwygTMD0IZRcZk0upONXbVRWPeyk+gB9lm+cZv9TSjO
z23HFtz30dZGm6fKa+l3D/2gthsjgx0QGtkJAITgRNOidSOzNIb2ILCkXhAd4FJG
AJ2xDx8hcFH1mt0G/FX0Kw4zd8NLQsLxdxP8c4CU6x+7Nz/OAipmsHMdMqUybDKw
juDEI/9bfU1lcKwrmz3O2+BtjjKAvpafkmO8l7tdufThcV4q5O8DIrGKZTqPwJNl
1IXNDw9bg1kWRxYtnCQ6yICmJhSFm/Y3m6xv+cXDBlHz4n/FsRC6UfTd
-----END CERTIFICATE-----
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            02:03:bc:53:59:6b:34:c7:18:f5:01:50:66
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Google Trust Services LLC, CN = GTS Root R1
        Validity
            Not Before: Aug 13 00:00:42 2020 GMT
            Not After : Sep 30 00:00:42 2027 GMT
        Subject: C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:f5:88:df:e7:62:8c:1e:37:f8:37:42:90:7f:6c:
                    87:d0:fb:65:82:25:fd:e8:cb:6b:a4:ff:6d:e9:5a:
                    23:e2:99:f6:1c:e9:92:03:99:13:7c:09:0a:8a:fa:
                    42:d6:5e:56:24:aa:7a:33:84:1f:d1:e9:69:bb:b9:
                    74:ec:57:4c:66:68:93:77:37:55:53:fe:39:10:4d:
                    b7:34:bb:5f:25:77:37:3b:17:94:ea:3c:e5:9d:d5:
                    bc:c3:b4:43:eb:2e:a7:47:ef:b0:44:11:63:d8:b4:
                    41:85:dd:41:30:48:93:1b:bf:b7:f6:e0:45:02:21:
                    e0:96:42:17:cf:d9:2b:65:56:34:07:26:04:0d:a8:
                    fd:7d:ca:2e:ef:ea:48:7c:37:4d:3f:00:9f:83:df:
                    ef:75:84:2e:79:57:5c:fc:57:6e:1a:96:ff:fc:8c:
                    9a:a6:99:be:25:d9:7f:96:2c:06:f7:11:2a:02:80:
                    80:eb:63:18:3c:50:49:87:e5:8a:ca:5f:19:2b:59:
                    96:81:00:a0:fb:51:db:ca:77:0b:0b:c9:96:4f:ef:
                    70:49:c7:5c:6d:20:fd:99:b4:b4:e2:ca:2e:77:fd:
                    2d:dc:0b:b6:6b:13:0c:8c:19:2b:17:96:98:b9:f0:
                    8b:f6:a0:27:bb:b6:e3:8d:51:8f:bd:ae:c7:9b:b1:
                    89:9d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Basic Constraints: critical
                CA:TRUE, pathlen:0
            X509v3 Subject Key Identifier: 
                8A:74:7F:AF:85:CD:EE:95:CD:3D:9C:D0:E2:46:14:F3:71:35:1D:27
            X509v3 Authority Key Identifier: 
                keyid:E4:AF:2B:26:71:1A:2B:48:27:85:2F:52:66:2C:EF:F0:89:13:71:3E

            Authority Information Access: 
                OCSP - URI:http://ocsp.pki.goog/gtsr1
                CA Issuers - URI:http://pki.goog/repo/certs/gtsr1.der

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crl.pki.goog/gtsr1/gtsr1.crl

            X509v3 Certificate Policies: 
                Policy: 1.3.6.1.4.1.11129.2.5.3
                  CPS: https://pki.goog/repository/
                Policy: 2.23.140.1.2.1
                Policy: 2.23.140.1.2.2

    Signature Algorithm: sha256WithRSAEncryption
         89:7d:ac:20:5c:0c:3c:be:9a:a8:57:95:1b:b4:ae:fa:ab:a5:
         72:71:b4:36:95:fd:df:40:11:03:4c:c2:46:14:bb:14:24:ab:
         f0:50:71:22:db:ad:c4:6e:7f:cf:f1:6a:6f:c8:83:1b:d8:ce:
         89:5f:87:6c:87:b8:a9:0c:a3:9b:a1:62:94:93:95:df:5b:ae:
         66:19:0b:02:96:9e:fc:b5:e7:10:69:3e:7a:cb:46:49:5f:46:
         e1:41:b1:d7:98:4d:65:34:00:80:1a:3f:4f:9f:6c:7f:49:00:
         81:53:41:a4:92:21:82:82:1a:f1:a3:44:5b:2a:50:12:13:4d:
         c1:53:36:f3:42:08:af:54:fa:8e:77:53:1b:64:38:27:17:09:
         bd:58:c9:1b:7c:39:2d:5b:f3:ce:d4:ed:97:db:14:03:bf:09:
         53:24:1f:c2:0c:04:79:98:26:f2:61:f1:53:52:fd:42:8c:1b:
         66:2b:3f:15:a1:bb:ff:f6:9b:e3:81:9a:01:06:71:89:35:28:
         24:dd:e1:bd:eb:19:2d:e1:48:cb:3d:59:83:51:b4:74:c6:9d:
         7c:c6:b1:86:5b:af:cc:34:c4:d3:cc:d4:81:11:95:00:a1:f4:
         12:22:01:fa:b4:83:71:af:8c:b7:8c:73:24:ac:37:53:c2:00:
         90:3f:11:fe:5c:ed:36:94:10:3b:bd:29:ae:e2:c7:3a:62:3b:
         6c:63:d9:80:bf:59:71:ac:63:27:b9:4c:17:a0:da:f6:73:15:
         bf:2a:de:8f:f3:a5:6c:32:81:33:03:d0:86:51:71:99:34:ba:
         93:8d:5d:b5:51:58:f7:b2:93:e8:01:f6:59:be:71:9b:fd:4d:
         28:ce:cf:6d:c7:16:dc:f7:d1:d6:46:9b:a7:ca:6b:e9:77:0f:
         fd:a0:b6:1b:23:83:1d:10:1a:d9:09:00:84:e0:44:d3:a2:75:
         23:b3:34:86:f6:20:b0:a4:5e:10:1d:e0:52:46:00:9d:b1:0f:
         1f:21:70:51:f5:9a:dd:06:fc:55:f4:2b:0e:33:77:c3:4b:42:
         c2:f1:77:13:fc:73:80:94:eb:1f:bb:37:3f:ce:02:2a:66:b0:
         73:1d:32:a5:32:6c:32:b0:8e:e0:c4:23:ff:5b:7d:4d:65:70:
         ac:2b:9b:3d:ce:db:e0:6d:8e:32:80:be:96:9f:92:63:bc:97:
         bb:5d:b9:f4:e1:71:5e:2a:e4:ef:03:22:b1:8a:65:3a:8f:c0:
         93:65:d4:85:cd:0f:0f:5b:83:59:16:47:16:2d:9c:24:3a:c8:
         80:a6:26:14:85:9b:f6:37:9b:ac:6f:f9:c5:c3:06:51:f3:e2:
         7f:c5:b1:10:ba:51:f4:dd
```

```sh
$ openssl s_client -help
Usage: s_client [options]
Valid options are:
 -help                      Display this summary
 -host val                  Use -connect instead
 -port +int                 Use -connect instead
 -connect val               TCP/IP where to connect (default is :4433)
 -bind val                  bind local address for connection
 -proxy val                 Connect to via specified proxy to the real server
 -unix val                  Connect over the specified Unix-domain socket
 -4                         Use IPv4 only
 -6                         Use IPv6 only
 -verify +int               Turn on peer certificate verification
 -cert infile               Certificate file to use, PEM format assumed
 -certform PEM|DER          Certificate format (PEM or DER) PEM default
 -nameopt val               Various certificate name options
 -key val                   Private key file to use, if not in -cert file
 -keyform PEM|DER|ENGINE    Key format (PEM, DER or engine) PEM default
 -pass val                  Private key file pass phrase source
 -CApath dir                PEM format directory of CA's
 -CAfile infile             PEM format file of CA's
 -no-CAfile                 Do not load the default certificates file
 -no-CApath                 Do not load certificates from the default certificates directory
 -requestCAfile infile      PEM format file of CA names to send to the server
 -dane_tlsa_domain val      DANE TLSA base domain
 -dane_tlsa_rrdata val      DANE TLSA rrdata presentation form
 -dane_ee_no_namechecks     Disable name checks when matching DANE-EE(3) TLSA records
 -reconnect                 Drop and re-make the connection with the same Session-ID
 -showcerts                 Show all certificates sent by the server
 -debug                     Extra output
 -msg                       Show protocol messages
 -msgfile outfile           File to send output of -msg or -trace, instead of stdout
 -nbio_test                 More ssl protocol testing
 -state                     Print the ssl states
 -crlf                      Convert LF from terminal into CRLF
 -quiet                     No s_client output
 -ign_eof                   Ignore input eof (default when -quiet)
 -no_ign_eof                Don't ignore input eof
 -starttls val              Use the appropriate STARTTLS command before starting TLS
 -xmpphost val              Alias of -name option for "-starttls xmpp[-server]"
 -rand val                  Load the file(s) into the random number generator
 -writerand outfile         Write random data to the specified file
 -sess_out outfile          File to write SSL session to
 -sess_in infile            File to read SSL session from
 -use_srtp val              Offer SRTP key management with a colon-separated profile list
 -keymatexport val          Export keying material using label
 -keymatexportlen +int      Export len bytes of keying material (default 20)
 -maxfraglen +int           Enable Maximum Fragment Length Negotiation (len values: 512, 1024, 2048 and 4096)
 -fallback_scsv             Send the fallback SCSV
 -name val                  Hostname to use for "-starttls lmtp", "-starttls smtp" or "-starttls xmpp[-server]"
 -CRL infile                CRL file to use
 -crl_download              Download CRL from distribution points
 -CRLform PEM|DER           CRL format (PEM or DER) PEM is default
 -verify_return_error       Close connection on verification error
 -verify_quiet              Restrict verify output to errors
 -brief                     Restrict output to brief summary of connection parameters
 -prexit                    Print session information when the program exits
 -security_debug            Enable security debug messages
 -security_debug_verbose    Output more security debug output
 -cert_chain infile         Certificate chain file (in PEM format)
 -chainCApath dir           Use dir as certificate store path to build CA certificate chain
 -verifyCApath dir          Use dir as certificate store path to verify CA certificate
 -build_chain               Build certificate chain
 -chainCAfile infile        CA file for certificate chain (PEM format)
 -verifyCAfile infile       CA file for certificate verification (PEM format)
 -nocommands                Do not use interactive command letters
 -servername val            Set TLS extension servername (SNI) in ClientHello (default)
 -noservername              Do not send the server name (SNI) extension in the ClientHello
 -tlsextdebug               Hex dump of all TLS extensions received
 -status                    Request certificate status from server
 -serverinfo val            types  Send empty ClientHello extensions (comma-separated numbers)
 -alpn val                  Enable ALPN extension, considering named protocols supported (comma-separated list)
 -async                     Support asynchronous operation
 -ssl_config val            Use specified configuration file
 -max_send_frag +int        Maximum Size of send frames 
 -split_send_frag +int      Size used to split data for encrypt pipelines
 -max_pipelines +int        Maximum number of encrypt/decrypt pipelines to be used
 -read_buf +int             Default read buffer size to be used for connections
 -no_ssl3                   Just disable SSLv3
 -no_tls1                   Just disable TLSv1
 -no_tls1_1                 Just disable TLSv1.1
 -no_tls1_2                 Just disable TLSv1.2
 -no_tls1_3                 Just disable TLSv1.3
 -bugs                      Turn on SSL bug compatibility
 -no_comp                   Disable SSL/TLS compression (default)
 -comp                      Use SSL/TLS-level compression
 -no_ticket                 Disable use of TLS session tickets
 -serverpref                Use server's cipher preferences
 -legacy_renegotiation      Enable use of legacy renegotiation (dangerous)
 -no_renegotiation          Disable all renegotiation.
 -legacy_server_connect     Allow initial connection to servers that don't support RI
 -no_resumption_on_reneg    Disallow session resumption on renegotiation
 -no_legacy_server_connect  Disallow initial connection to servers that don't support RI
 -allow_no_dhe_kex          In TLSv1.3 allow non-(ec)dhe based key exchange on resumption
 -prioritize_chacha         Prioritize ChaCha ciphers when preferred by clients
 -strict                    Enforce strict certificate checks as per TLS standard
 -sigalgs val               Signature algorithms to support (colon-separated list)
 -client_sigalgs val        Signature algorithms to support for client certificate authentication (colon-separated list)
 -groups val                Groups to advertise (colon-separated list)
 -curves val                Groups to advertise (colon-separated list)
 -named_curve val           Elliptic curve used for ECDHE (server-side only)
 -cipher val                Specify TLSv1.2 and below cipher list to be used
 -ciphersuites val          Specify TLSv1.3 ciphersuites to be used
 -min_protocol val          Specify the minimum protocol version to be used
 -max_protocol val          Specify the maximum protocol version to be used
 -record_padding val        Block size to pad TLS 1.3 records to.
 -debug_broken_protocol     Perform all sorts of protocol violations for testing purposes
 -no_middlebox              Disable TLSv1.3 middlebox compat mode
 -policy val                adds policy to the acceptable policy set
 -purpose val               certificate chain purpose
 -verify_name val           verification policy name
 -verify_depth int          chain depth limit
 -auth_level int            chain authentication security level
 -attime intmax             verification epoch time
 -verify_hostname val       expected peer hostname
 -verify_email val          expected peer email
 -verify_ip val             expected peer IP address
 -ignore_critical           permit unhandled critical extensions
 -issuer_checks             (deprecated)
 -crl_check                 check leaf certificate revocation
 -crl_check_all             check full chain revocation
 -policy_check              perform rfc5280 policy checks
 -explicit_policy           set policy variable require-explicit-policy
 -inhibit_any               set policy variable inhibit-any-policy
 -inhibit_map               set policy variable inhibit-policy-mapping
 -x509_strict               disable certificate compatibility work-arounds
 -extended_crl              enable extended CRL features
 -use_deltas                use delta CRLs
 -policy_print              print policy processing diagnostics
 -check_ss_sig              check root CA self-signatures
 -trusted_first             search trust store first (default)
 -suiteB_128_only           Suite B 128-bit-only mode
 -suiteB_128                Suite B 128-bit mode allowing 192-bit algorithms
 -suiteB_192                Suite B 192-bit-only mode
 -partial_chain             accept chains anchored by intermediate trust-store CAs
 -no_alt_chains             (deprecated)
 -no_check_time             ignore certificate validity time
 -allow_proxy_certs         allow the use of proxy certificates
 -xkey infile               key for Extended certificates
 -xcert infile              cert for Extended certificates
 -xchain infile             chain for Extended certificates
 -xchain_build              build certificate chain for the extended certificates
 -xcertform PEM|DER         format of Extended certificate (PEM or DER) PEM default 
 -xkeyform PEM|DER          format of Extended certificate's key (PEM or DER) PEM default
 -tls1                      Just use TLSv1
 -tls1_1                    Just use TLSv1.1
 -tls1_2                    Just use TLSv1.2
 -tls1_3                    Just use TLSv1.3
 -dtls                      Use any version of DTLS
 -timeout                   Enable send/receive timeout on DTLS connections
 -mtu +int                  Set the link layer MTU
 -dtls1                     Just use DTLSv1
 -dtls1_2                   Just use DTLSv1.2
 -nbio                      Use non-blocking IO
 -psk_identity val          PSK identity
 -psk val                   PSK in hex (without 0x)
 -psk_session infile        File to read PSK SSL session from
 -srpuser val               SRP authentication for 'user'
 -srppass val               Password for 'user'
 -srp_lateuser              SRP username into second ClientHello message
 -srp_moregroups            Tolerate other than the known g N values.
 -srp_strength +int         Minimal length in bits for N
 -nextprotoneg val          Enable NPN extension, considering named protocols supported (comma-separated list)
 -engine val                Use engine, possibly a hardware device
 -ssl_client_engine val     Specify engine to be used for client certificate operations
 -ct                        Request and parse SCTs (also enables OCSP stapling)
 -noct                      Do not request or parse SCTs (default)
 -ctlogfile infile          CT log list CONF file
 -keylogfile outfile        Write TLS secrets to file
 -early_data infile         File to send as early data
 -enable_pha                Enable post-handshake-authentication
```

```sh
$ openssl version
OpenSSL 1.1.1w  11 Sep 2023
```
