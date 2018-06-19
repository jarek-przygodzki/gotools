# gotools
Various small command-line utilities written in Go

## httpl
HTTP server logging all received data

```
httpl -h
Usage of httpl:
  -addr string
        The address to listen to; default is "" (all interfaces).
  -port int
        The port to listen on. (default 8000)
```


```
httpl -port 8080
Listening on :8080.
2018/06/19 14:57:35 [::1]:52578 GET /api/v1/
2018/06/19 14:57:35
00000000  47 45 54 20 2f 61 70 69  2f 76 31 2f 20 48 54 54  |GET /api/v1/ HTT|
00000010  50 2f 31 2e 31 0d 0a 48  6f 73 74 3a 20 6c 6f 63  |P/1.1..Host: loc|
00000020  61 6c 68 6f 73 74 3a 38  30 38 30 0d 0a 41 63 63  |alhost:8080..Acc|
00000030  65 70 74 3a 20 2a 2f 2a  0d 0a 55 73 65 72 2d 41  |ept: */*..User-A|
00000040  67 65 6e 74 3a 20 63 75  72 6c 2f 37 2e 35 35 2e  |gent: curl/7.55.|
00000050  31 0d 0a 0d 0a
```
