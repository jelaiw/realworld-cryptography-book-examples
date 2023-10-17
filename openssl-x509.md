```sh
$ openssl x509 -in _.google.crt -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            6f:74:d0:38:20:cb:1b:45:10:cb:4c:1e:a2:7d:a2:eb
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, O = Google Trust Services LLC, CN = GTS CA 1C3
        Validity
            Not Before: Sep 18 08:19:26 2023 GMT
            Not After : Dec 11 08:19:25 2023 GMT
        Subject: CN = *.google.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:4a:46:13:89:87:0e:95:78:27:42:02:80:02:69:
                    39:68:1b:bc:d6:9e:cb:d7:86:ae:23:ee:b3:97:bf:
                    4b:2e:06:57:c0:b4:e5:32:de:6b:09:e0:46:a2:99:
                    d2:2d:f8:a5:31:d9:9c:8b:fa:16:43:e3:e9:3c:a0:
                    92:6d:17:60:4f
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
                33:3E:CA:90:E9:65:4A:95:78:D4:F0:25:6F:55:49:ED:02:1A:79:DB
            X509v3 Authority Key Identifier: 
                keyid:8A:74:7F:AF:85:CD:EE:95:CD:3D:9C:D0:E2:46:14:F3:71:35:1D:27

            Authority Information Access: 
                OCSP - URI:http://ocsp.pki.goog/gts1c3
                CA Issuers - URI:http://pki.goog/repo/certs/gts1c3.der

            X509v3 Subject Alternative Name: 
                DNS:*.google.com, DNS:*.appengine.google.com, DNS:*.bdn.dev, DNS:*.origin-test.bdn.dev, DNS:*.cloud.google.com, DNS:*.crowdsource.google.com, DNS:*.datacompute.google.com, DNS:*.google.ca, DNS:*.google.cl, DNS:*.google.co.in, DNS:*.google.co.jp, DNS:*.google.co.uk, DNS:*.google.com.ar, DNS:*.google.com.au, DNS:*.google.com.br, DNS:*.google.com.co, DNS:*.google.com.mx, DNS:*.google.com.tr, DNS:*.google.com.vn, DNS:*.google.de, DNS:*.google.es, DNS:*.google.fr, DNS:*.google.hu, DNS:*.google.it, DNS:*.google.nl, DNS:*.google.pl, DNS:*.google.pt, DNS:*.googleadapis.com, DNS:*.googleapis.cn, DNS:*.googlevideo.com, DNS:*.gstatic.cn, DNS:*.gstatic-cn.com, DNS:googlecnapps.cn, DNS:*.googlecnapps.cn, DNS:googleapps-cn.com, DNS:*.googleapps-cn.com, DNS:gkecnapps.cn, DNS:*.gkecnapps.cn, DNS:googledownloads.cn, DNS:*.googledownloads.cn, DNS:recaptcha.net.cn, DNS:*.recaptcha.net.cn, DNS:recaptcha-cn.net, DNS:*.recaptcha-cn.net, DNS:widevine.cn, DNS:*.widevine.cn, DNS:ampproject.org.cn, DNS:*.ampproject.org.cn, DNS:ampproject.net.cn, DNS:*.ampproject.net.cn, DNS:google-analytics-cn.com, DNS:*.google-analytics-cn.com, DNS:googleadservices-cn.com, DNS:*.googleadservices-cn.com, DNS:googlevads-cn.com, DNS:*.googlevads-cn.com, DNS:googleapis-cn.com, DNS:*.googleapis-cn.com, DNS:googleoptimize-cn.com, DNS:*.googleoptimize-cn.com, DNS:doubleclick-cn.net, DNS:*.doubleclick-cn.net, DNS:*.fls.doubleclick-cn.net, DNS:*.g.doubleclick-cn.net, DNS:doubleclick.cn, DNS:*.doubleclick.cn, DNS:*.fls.doubleclick.cn, DNS:*.g.doubleclick.cn, DNS:dartsearch-cn.net, DNS:*.dartsearch-cn.net, DNS:googletraveladservices-cn.com, DNS:*.googletraveladservices-cn.com, DNS:googletagservices-cn.com, DNS:*.googletagservices-cn.com, DNS:googletagmanager-cn.com, DNS:*.googletagmanager-cn.com, DNS:googlesyndication-cn.com, DNS:*.googlesyndication-cn.com, DNS:*.safeframe.googlesyndication-cn.com, DNS:app-measurement-cn.com, DNS:*.app-measurement-cn.com, DNS:gvt1-cn.com, DNS:*.gvt1-cn.com, DNS:gvt2-cn.com, DNS:*.gvt2-cn.com, DNS:2mdn-cn.net, DNS:*.2mdn-cn.net, DNS:googleflights-cn.net, DNS:*.googleflights-cn.net, DNS:admob-cn.com, DNS:*.admob-cn.com, DNS:googlesandbox-cn.com, DNS:*.googlesandbox-cn.com, DNS:*.safenup.googlesandbox-cn.com, DNS:*.gstatic.com, DNS:*.metric.gstatic.com, DNS:*.gvt1.com, DNS:*.gcpcdn.gvt1.com, DNS:*.gvt2.com, DNS:*.gcp.gvt2.com, DNS:*.url.google.com, DNS:*.youtube-nocookie.com, DNS:*.ytimg.com, DNS:android.com, DNS:*.android.com, DNS:*.flash.android.com, DNS:g.cn, DNS:*.g.cn, DNS:g.co, DNS:*.g.co, DNS:goo.gl, DNS:www.goo.gl, DNS:google-analytics.com, DNS:*.google-analytics.com, DNS:google.com, DNS:googlecommerce.com, DNS:*.googlecommerce.com, DNS:ggpht.cn, DNS:*.ggpht.cn, DNS:urchin.com, DNS:*.urchin.com, DNS:youtu.be, DNS:youtube.com, DNS:*.youtube.com, DNS:youtubeeducation.com, DNS:*.youtubeeducation.com, DNS:youtubekids.com, DNS:*.youtubekids.com, DNS:yt.be, DNS:*.yt.be, DNS:android.clients.google.com, DNS:developer.android.google.cn, DNS:developers.android.google.cn, DNS:source.android.google.cn
            X509v3 Certificate Policies: 
                Policy: 2.23.140.1.2.1
                Policy: 1.3.6.1.4.1.11129.2.5.3

            X509v3 CRL Distribution Points: 

                Full Name:
                  URI:http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl

            CT Precertificate SCTs: 
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : 7A:32:8C:54:D8:B7:2D:B6:20:EA:38:E0:52:1E:E9:84:
                                16:70:32:13:85:4D:3B:D2:2B:C1:3A:57:A3:52:EB:52
                    Timestamp : Sep 18 09:19:30.353 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:46:02:21:00:86:FB:D9:B6:D6:79:CB:43:09:8C:3F:
                                B8:34:E1:84:8D:E3:94:14:A1:88:2B:01:14:38:06:22:
                                34:4B:BC:EE:F6:02:21:00:CF:FB:BA:29:93:26:44:06:
                                8B:AD:CB:A0:02:24:2C:C7:B9:0C:58:14:92:12:62:81:
                                23:AE:20:BD:C1:D6:2B:7C
                Signed Certificate Timestamp:
                    Version   : v1 (0x0)
                    Log ID    : AD:F7:BE:FA:7C:FF:10:C8:8B:9D:3D:9C:1E:3E:18:6A:
                                B4:67:29:5D:CF:B1:0C:24:CA:85:86:34:EB:DC:82:8A
                    Timestamp : Sep 18 09:19:30.364 2023 GMT
                    Extensions: none
                    Signature : ecdsa-with-SHA256
                                30:45:02:20:04:98:7D:03:1B:92:7B:6C:28:91:3B:C3:
                                4F:57:B9:FD:05:05:59:B2:DE:63:EF:C2:AF:28:25:7C:
                                56:F7:B6:77:02:21:00:9B:22:80:DE:02:C9:33:CD:83:
                                06:EA:10:7A:6C:06:2F:CE:66:90:8A:3F:09:D0:CB:A4:
                                7C:E2:7D:A9:01:0E:25
    Signature Algorithm: sha256WithRSAEncryption
         50:70:b6:f9:a0:b6:38:16:8c:75:6d:4a:33:ba:6a:c6:47:e3:
         6d:9c:8c:8c:e4:70:07:5a:eb:d9:68:f9:68:de:85:f0:e1:19:
         eb:04:de:6f:a1:7d:0e:d9:16:12:5f:ae:73:2d:d7:dc:50:9d:
         b3:ec:52:bb:4a:c8:77:15:e2:0f:43:24:c1:7a:5b:15:39:c4:
         79:eb:7a:93:a0:4c:b2:2f:3c:79:d6:34:a3:4f:58:88:e8:66:
         56:f6:59:39:2b:a9:39:05:b5:32:7a:64:08:27:ab:db:6b:5d:
         fd:93:ab:05:bd:73:ec:ee:92:5f:07:6b:b2:12:c0:14:59:6c:
         e3:8e:51:d4:49:40:52:db:75:b5:19:e6:bd:1f:68:c2:e2:c4:
         dc:44:e9:75:25:26:9f:56:4b:a1:5a:96:e9:8a:6e:92:c2:aa:
         16:91:9a:b9:a1:ff:ba:35:b3:e0:0e:84:ef:07:91:fa:a6:d3:
         f3:f8:50:08:83:d2:06:e1:71:93:37:1a:58:5d:3a:d4:10:3b:
         5f:cb:be:09:0a:37:d1:96:02:14:6a:e5:27:16:e5:7a:47:9c:
         7e:29:43:c4:a7:bb:a4:a9:ed:f8:38:5a:b5:ce:10:db:5f:54:
         a5:70:54:e1:88:da:7f:e0:99:ec:56:4f:dd:f3:07:ad:a5:15:
         17:f2:02:89
-----BEGIN CERTIFICATE-----
MIIOPDCCDSSgAwIBAgIQb3TQOCDLG0UQy0weon2i6zANBgkqhkiG9w0BAQsFADBG
MQswCQYDVQQGEwJVUzEiMCAGA1UEChMZR29vZ2xlIFRydXN0IFNlcnZpY2VzIExM
QzETMBEGA1UEAxMKR1RTIENBIDFDMzAeFw0yMzA5MTgwODE5MjZaFw0yMzEyMTEw
ODE5MjVaMBcxFTATBgNVBAMMDCouZ29vZ2xlLmNvbTBZMBMGByqGSM49AgEGCCqG
SM49AwEHA0IABEpGE4mHDpV4J0ICgAJpOWgbvNaey9eGriPus5e/Sy4GV8C05TLe
awngRqKZ0i34pTHZnIv6FkPj6Tygkm0XYE+jggweMIIMGjAOBgNVHQ8BAf8EBAMC
B4AwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQU
Mz7KkOllSpV41PAlb1VJ7QIaedswHwYDVR0jBBgwFoAUinR/r4XN7pXNPZzQ4kYU
83E1HScwagYIKwYBBQUHAQEEXjBcMCcGCCsGAQUFBzABhhtodHRwOi8vb2NzcC5w
a2kuZ29vZy9ndHMxYzMwMQYIKwYBBQUHMAKGJWh0dHA6Ly9wa2kuZ29vZy9yZXBv
L2NlcnRzL2d0czFjMy5kZXIwggnNBgNVHREEggnEMIIJwIIMKi5nb29nbGUuY29t
ghYqLmFwcGVuZ2luZS5nb29nbGUuY29tggkqLmJkbi5kZXaCFSoub3JpZ2luLXRl
c3QuYmRuLmRldoISKi5jbG91ZC5nb29nbGUuY29tghgqLmNyb3dkc291cmNlLmdv
b2dsZS5jb22CGCouZGF0YWNvbXB1dGUuZ29vZ2xlLmNvbYILKi5nb29nbGUuY2GC
CyouZ29vZ2xlLmNsgg4qLmdvb2dsZS5jby5pboIOKi5nb29nbGUuY28uanCCDiou
Z29vZ2xlLmNvLnVrgg8qLmdvb2dsZS5jb20uYXKCDyouZ29vZ2xlLmNvbS5hdYIP
Ki5nb29nbGUuY29tLmJygg8qLmdvb2dsZS5jb20uY2+CDyouZ29vZ2xlLmNvbS5t
eIIPKi5nb29nbGUuY29tLnRygg8qLmdvb2dsZS5jb20udm6CCyouZ29vZ2xlLmRl
ggsqLmdvb2dsZS5lc4ILKi5nb29nbGUuZnKCCyouZ29vZ2xlLmh1ggsqLmdvb2ds
ZS5pdIILKi5nb29nbGUubmyCCyouZ29vZ2xlLnBsggsqLmdvb2dsZS5wdIISKi5n
b29nbGVhZGFwaXMuY29tgg8qLmdvb2dsZWFwaXMuY26CESouZ29vZ2xldmlkZW8u
Y29tggwqLmdzdGF0aWMuY26CECouZ3N0YXRpYy1jbi5jb22CD2dvb2dsZWNuYXBw
cy5jboIRKi5nb29nbGVjbmFwcHMuY26CEWdvb2dsZWFwcHMtY24uY29tghMqLmdv
b2dsZWFwcHMtY24uY29tggxna2VjbmFwcHMuY26CDiouZ2tlY25hcHBzLmNughJn
b29nbGVkb3dubG9hZHMuY26CFCouZ29vZ2xlZG93bmxvYWRzLmNughByZWNhcHRj
aGEubmV0LmNughIqLnJlY2FwdGNoYS5uZXQuY26CEHJlY2FwdGNoYS1jbi5uZXSC
EioucmVjYXB0Y2hhLWNuLm5ldIILd2lkZXZpbmUuY26CDSoud2lkZXZpbmUuY26C
EWFtcHByb2plY3Qub3JnLmNughMqLmFtcHByb2plY3Qub3JnLmNughFhbXBwcm9q
ZWN0Lm5ldC5jboITKi5hbXBwcm9qZWN0Lm5ldC5jboIXZ29vZ2xlLWFuYWx5dGlj
cy1jbi5jb22CGSouZ29vZ2xlLWFuYWx5dGljcy1jbi5jb22CF2dvb2dsZWFkc2Vy
dmljZXMtY24uY29tghkqLmdvb2dsZWFkc2VydmljZXMtY24uY29tghFnb29nbGV2
YWRzLWNuLmNvbYITKi5nb29nbGV2YWRzLWNuLmNvbYIRZ29vZ2xlYXBpcy1jbi5j
b22CEyouZ29vZ2xlYXBpcy1jbi5jb22CFWdvb2dsZW9wdGltaXplLWNuLmNvbYIX
Ki5nb29nbGVvcHRpbWl6ZS1jbi5jb22CEmRvdWJsZWNsaWNrLWNuLm5ldIIUKi5k
b3VibGVjbGljay1jbi5uZXSCGCouZmxzLmRvdWJsZWNsaWNrLWNuLm5ldIIWKi5n
LmRvdWJsZWNsaWNrLWNuLm5ldIIOZG91YmxlY2xpY2suY26CECouZG91YmxlY2xp
Y2suY26CFCouZmxzLmRvdWJsZWNsaWNrLmNughIqLmcuZG91YmxlY2xpY2suY26C
EWRhcnRzZWFyY2gtY24ubmV0ghMqLmRhcnRzZWFyY2gtY24ubmV0gh1nb29nbGV0
cmF2ZWxhZHNlcnZpY2VzLWNuLmNvbYIfKi5nb29nbGV0cmF2ZWxhZHNlcnZpY2Vz
LWNuLmNvbYIYZ29vZ2xldGFnc2VydmljZXMtY24uY29tghoqLmdvb2dsZXRhZ3Nl
cnZpY2VzLWNuLmNvbYIXZ29vZ2xldGFnbWFuYWdlci1jbi5jb22CGSouZ29vZ2xl
dGFnbWFuYWdlci1jbi5jb22CGGdvb2dsZXN5bmRpY2F0aW9uLWNuLmNvbYIaKi5n
b29nbGVzeW5kaWNhdGlvbi1jbi5jb22CJCouc2FmZWZyYW1lLmdvb2dsZXN5bmRp
Y2F0aW9uLWNuLmNvbYIWYXBwLW1lYXN1cmVtZW50LWNuLmNvbYIYKi5hcHAtbWVh
c3VyZW1lbnQtY24uY29tggtndnQxLWNuLmNvbYINKi5ndnQxLWNuLmNvbYILZ3Z0
Mi1jbi5jb22CDSouZ3Z0Mi1jbi5jb22CCzJtZG4tY24ubmV0gg0qLjJtZG4tY24u
bmV0ghRnb29nbGVmbGlnaHRzLWNuLm5ldIIWKi5nb29nbGVmbGlnaHRzLWNuLm5l
dIIMYWRtb2ItY24uY29tgg4qLmFkbW9iLWNuLmNvbYIUZ29vZ2xlc2FuZGJveC1j
bi5jb22CFiouZ29vZ2xlc2FuZGJveC1jbi5jb22CHiouc2FmZW51cC5nb29nbGVz
YW5kYm94LWNuLmNvbYINKi5nc3RhdGljLmNvbYIUKi5tZXRyaWMuZ3N0YXRpYy5j
b22CCiouZ3Z0MS5jb22CESouZ2NwY2RuLmd2dDEuY29tggoqLmd2dDIuY29tgg4q
LmdjcC5ndnQyLmNvbYIQKi51cmwuZ29vZ2xlLmNvbYIWKi55b3V0dWJlLW5vY29v
a2llLmNvbYILKi55dGltZy5jb22CC2FuZHJvaWQuY29tgg0qLmFuZHJvaWQuY29t
ghMqLmZsYXNoLmFuZHJvaWQuY29tggRnLmNuggYqLmcuY26CBGcuY2+CBiouZy5j
b4IGZ29vLmdsggp3d3cuZ29vLmdsghRnb29nbGUtYW5hbHl0aWNzLmNvbYIWKi5n
b29nbGUtYW5hbHl0aWNzLmNvbYIKZ29vZ2xlLmNvbYISZ29vZ2xlY29tbWVyY2Uu
Y29tghQqLmdvb2dsZWNvbW1lcmNlLmNvbYIIZ2dwaHQuY26CCiouZ2dwaHQuY26C
CnVyY2hpbi5jb22CDCoudXJjaGluLmNvbYIIeW91dHUuYmWCC3lvdXR1YmUuY29t
gg0qLnlvdXR1YmUuY29tghR5b3V0dWJlZWR1Y2F0aW9uLmNvbYIWKi55b3V0dWJl
ZWR1Y2F0aW9uLmNvbYIPeW91dHViZWtpZHMuY29tghEqLnlvdXR1YmVraWRzLmNv
bYIFeXQuYmWCByoueXQuYmWCGmFuZHJvaWQuY2xpZW50cy5nb29nbGUuY29tghtk
ZXZlbG9wZXIuYW5kcm9pZC5nb29nbGUuY26CHGRldmVsb3BlcnMuYW5kcm9pZC5n
b29nbGUuY26CGHNvdXJjZS5hbmRyb2lkLmdvb2dsZS5jbjAhBgNVHSAEGjAYMAgG
BmeBDAECATAMBgorBgEEAdZ5AgUDMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly9j
cmxzLnBraS5nb29nL2d0czFjMy96ZEFUdDBFeF9Gay5jcmwwggEFBgorBgEEAdZ5
AgQCBIH2BIHzAPEAdwB6MoxU2LcttiDqOOBSHumEFnAyE4VNO9IrwTpXo1LrUgAA
AYqnlnYxAAAEAwBIMEYCIQCG+9m21nnLQwmMP7g04YSN45QUoYgrARQ4BiI0S7zu
9gIhAM/7uimTJkQGi63LoAIkLMe5DFgUkhJigSOuIL3B1it8AHYArfe++nz/EMiL
nT2cHj4YarRnKV3PsQwkyoWGNOvcgooAAAGKp5Z2PAAABAMARzBFAiAEmH0DG5J7
bCiRO8NPV7n9BQVZst5j78KvKCV8Vve2dwIhAJsigN4CyTPNgwbqEHpsBi/OZpCK
PwnQy6R84n2pAQ4lMA0GCSqGSIb3DQEBCwUAA4IBAQBQcLb5oLY4Fox1bUozumrG
R+NtnIyM5HAHWuvZaPlo3oXw4RnrBN5voX0O2RYSX65zLdfcUJ2z7FK7Ssh3FeIP
QyTBelsVOcR563qToEyyLzx51jSjT1iI6GZW9lk5K6k5BbUyemQIJ6vba139k6sF
vXPs7pJfB2uyEsAUWWzjjlHUSUBS23W1Gea9H2jC4sTcROl1JSafVkuhWpbpim6S
wqoWkZq5of+6NbPgDoTvB5H6ptPz+FAIg9IG4XGTNxpYXTrUEDtfy74JCjfRlgIU
auUnFuV6R5x+KUPEp7ukqe34OFq1zhDbX1SlcFThiNp/4JnsVk/d8wetpRUX8gKJ
-----END CERTIFICATE-----
```