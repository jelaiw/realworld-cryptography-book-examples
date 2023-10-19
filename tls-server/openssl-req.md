Generate a self-signed root certificate.
```sh
$ openssl req -new -x509 -days 29 -nodes -out cert.pem -keyout key.pem
Generating a RSA private key
......................................+++++
........+++++
writing new private key to 'key.pem'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:Raccoon City
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Umbrella Corp
Organizational Unit Name (eg, section) []:Research
Common Name (e.g. server FQDN or YOUR name) []:foobar.com
Email Address []:foo@foobar.com
```

Print certificate contents.
```sh
$ openssl x509 -in cert.pem -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            36:ed:ae:c7:fd:e7:f2:1b:b9:c9:78:91:67:c8:27:be:25:2d:52:26
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = AU, ST = Some-State, L = Raccoon City, O = Umbrella Corp, OU = Research, CN = foobar.com, emailAddress = foo@foobar.com
        Validity
            Not Before: Oct 19 23:40:40 2023 GMT
            Not After : Nov 17 23:40:40 2023 GMT
        Subject: C = AU, ST = Some-State, L = Raccoon City, O = Umbrella Corp, OU = Research, CN = foobar.com, emailAddress = foo@foobar.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:ed:91:d5:e4:a6:86:4e:04:54:74:b9:27:47:68:
                    a6:88:44:9a:17:5f:c2:83:f6:ca:43:ff:26:46:7a:
                    17:a4:e9:67:43:b5:36:3b:e2:ac:60:cc:a3:73:b2:
                    f1:5a:96:5e:39:e6:14:bd:ec:29:d3:14:2b:ab:9b:
                    d2:2f:4d:bd:d0:f0:6b:03:ea:91:e5:da:dc:14:5f:
                    af:07:7f:05:db:96:64:80:f6:41:09:00:1f:50:f5:
                    99:f4:bf:46:fe:ce:b2:ea:fe:4b:e4:cf:49:54:02:
                    6d:e2:82:24:35:07:e4:96:c0:8e:87:6d:77:25:c5:
                    9f:7f:0b:b9:31:67:96:14:a1:10:3d:72:8e:43:ec:
                    8d:0c:72:47:ba:bc:b1:55:b5:04:13:d7:84:5d:61:
                    35:40:f8:0e:5c:cb:2e:4f:71:c3:d5:19:78:34:bb:
                    de:69:10:ef:6f:3e:21:86:7a:fb:82:14:f2:e1:90:
                    80:d4:d0:ce:a3:c1:24:a2:5b:58:a6:c6:66:69:63:
                    c1:42:8f:01:ea:09:81:a6:3f:00:e9:8d:a4:22:b9:
                    37:a3:e0:3d:a0:51:e5:2e:04:3a:3f:c2:5b:36:44:
                    c0:53:ea:d7:78:47:19:3f:e0:01:c7:eb:aa:b8:82:
                    d9:1a:82:42:7c:f6:88:33:0b:8d:70:c1:02:c2:85:
                    5a:7d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Key Identifier: 
                BD:74:CF:50:86:76:32:EA:6E:96:11:14:BB:8F:5C:7B:02:C4:20:4D
            X509v3 Authority Key Identifier: 
                keyid:BD:74:CF:50:86:76:32:EA:6E:96:11:14:BB:8F:5C:7B:02:C4:20:4D

            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
         c1:3d:c6:30:12:b6:1c:0b:3b:45:ad:b2:12:06:67:38:70:dc:
         90:cc:99:36:b4:f0:36:84:4e:09:bb:15:76:de:d2:9a:bf:98:
         fc:5b:f9:98:70:92:42:7d:1e:01:4f:3c:8e:e3:b7:50:ea:cd:
         0c:7f:31:93:ac:d2:dc:8d:ce:65:8f:24:20:77:d2:14:1f:6d:
         32:2c:44:9f:50:e9:61:18:46:dc:a7:9b:a9:82:03:bc:d0:64:
         8e:72:a3:3a:56:9a:2f:a8:9a:4a:e6:c8:d3:a1:79:67:3c:4d:
         13:32:36:df:a6:6c:ee:07:b4:83:25:15:ca:e5:df:05:7f:96:
         ee:80:b5:75:a9:1f:f2:53:4a:4a:6c:c0:d2:87:6d:c4:73:fa:
         8c:6e:0e:30:67:19:71:41:fe:ad:ab:3b:21:4a:cd:9e:92:27:
         d6:fb:22:81:83:e1:1a:a6:96:1b:61:aa:96:b1:9e:3d:d8:45:
         eb:21:23:e4:b5:7e:1d:df:76:60:1e:30:99:e8:2f:bb:5c:bd:
         cc:16:ce:44:32:6f:d8:61:1c:65:6d:f8:3c:68:07:c7:85:0d:
         01:9d:c5:77:9a:43:41:00:63:ba:74:9a:fd:11:55:15:6b:28:
         71:ea:6f:1f:f5:03:ef:37:35:c6:7f:c7:b6:a8:52:e6:3e:88:
         b1:31:b6:80
-----BEGIN CERTIFICATE-----
MIIEEzCCAvugAwIBAgIUNu2ux/3n8hu5yXiRZ8gnviUtUiYwDQYJKoZIhvcNAQEL
BQAwgZgxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMRUwEwYDVQQH
DAxSYWNjb29uIENpdHkxFjAUBgNVBAoMDVVtYnJlbGxhIENvcnAxETAPBgNVBAsM
CFJlc2VhcmNoMRMwEQYDVQQDDApmb29iYXIuY29tMR0wGwYJKoZIhvcNAQkBFg5m
b29AZm9vYmFyLmNvbTAeFw0yMzEwMTkyMzQwNDBaFw0yMzExMTcyMzQwNDBaMIGY
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEVMBMGA1UEBwwMUmFj
Y29vbiBDaXR5MRYwFAYDVQQKDA1VbWJyZWxsYSBDb3JwMREwDwYDVQQLDAhSZXNl
YXJjaDETMBEGA1UEAwwKZm9vYmFyLmNvbTEdMBsGCSqGSIb3DQEJARYOZm9vQGZv
b2Jhci5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDtkdXkpoZO
BFR0uSdHaKaIRJoXX8KD9spD/yZGehek6WdDtTY74qxgzKNzsvFall455hS97CnT
FCurm9IvTb3Q8GsD6pHl2twUX68HfwXblmSA9kEJAB9Q9Zn0v0b+zrLq/kvkz0lU
Am3igiQ1B+SWwI6HbXclxZ9/C7kxZ5YUoRA9co5D7I0Mcke6vLFVtQQT14RdYTVA
+A5cyy5PccPVGXg0u95pEO9vPiGGevuCFPLhkIDU0M6jwSSiW1imxmZpY8FCjwHq
CYGmPwDpjaQiuTej4D2gUeUuBDo/wls2RMBT6td4Rxk/4AHH66q4gtkagkJ89ogz
C41wwQLChVp9AgMBAAGjUzBRMB0GA1UdDgQWBBS9dM9QhnYy6m6WERS7j1x7AsQg
TTAfBgNVHSMEGDAWgBS9dM9QhnYy6m6WERS7j1x7AsQgTTAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4IBAQDBPcYwErYcCztFrbISBmc4cNyQzJk2tPA2
hE4JuxV23tKav5j8W/mYcJJCfR4BTzyO47dQ6s0MfzGTrNLcjc5ljyQgd9IUH20y
LESfUOlhGEbcp5upggO80GSOcqM6VpovqJpK5sjToXlnPE0TMjbfpmzuB7SDJRXK
5d8Ff5bugLV1qR/yU0pKbMDSh23Ec/qMbg4wZxlxQf6tqzshSs2ekifW+yKBg+Ea
ppYbYaqWsZ492EXrISPktX4d33ZgHjCZ6C+7XL3MFs5EMm/YYRxlbfg8aAfHhQ0B
ncV3mkNBAGO6dJr9EVUVayhx6m8f9QPvNzXGf8e2qFLmPoixMbaA
-----END CERTIFICATE-----
```

Print help.
```sh
$ openssl req -help
Usage: req [options]
Valid options are:
 -help               Display this summary
 -inform PEM|DER     Input format - DER or PEM
 -outform PEM|DER    Output format - DER or PEM
 -in infile          Input file
 -out outfile        Output file
 -key val            Private key to use
 -keyform format     Key file format
 -pubkey             Output public key
 -new                New request
 -config infile      Request template file
 -keyout outfile     File to send the key to
 -passin val         Private key password source
 -passout val        Output file pass phrase source
 -rand val           Load the file(s) into the random number generator
 -writerand outfile  Write random data to the specified file
 -newkey val         Specify as type:bits
 -pkeyopt val        Public key options as opt:value
 -sigopt val         Signature parameter in n:v form
 -batch              Do not ask anything during request generation
 -newhdr             Output "NEW" in the header lines
 -modulus            RSA modulus
 -verify             Verify signature on REQ
 -nodes              Don't encrypt the output key
 -noout              Do not output REQ
 -verbose            Verbose output
 -utf8               Input characters are UTF8 (default ASCII)
 -nameopt val        Various certificate name options
 -reqopt val         Various request text options
 -text               Text form of request
 -x509               Output a x509 structure instead of a cert request
                     (Required by some CA's)
 -subj val           Set or modify request subject
 -subject            Output the request's subject
 -multivalue-rdn     Enable support for multivalued RDNs
 -days +int          Number of days cert is valid for
 -set_serial val     Serial number to use
 -addext val         Additional cert extension key=value pair (may be given more than once)
 -extensions val     Cert extension section (override value in config file)
 -reqexts val        Request extension section (override value in config file)
 -precert            Add a poison extension (implies -new)
 -*                  Any supported digest
 -engine val         Use engine, possibly a hardware device
 -keygen_engine val  Specify engine to be used for key generation operations
```