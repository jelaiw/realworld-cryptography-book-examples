```sh
$ echo -n "Alice""Bob""100""15" | openssl dgst -sha3-256
(stdin)= 34d6b397c7f2e8a303fc8e39d283771c0397dad74cef08376e27483efc29bb02
$ echo -n "Alice""Bob""1001""5" | openssl dgst -sha3-256
(stdin)= 34d6b397c7f2e8a303fc8e39d283771c0397dad74cef08376e27483efc29bb02
```
