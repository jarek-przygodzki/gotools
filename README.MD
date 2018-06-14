# gotools
Various small command-line utilities written in Go

## tcpl
TCP server logging all received data

```
tcpl -h

Usage of tcpl:
  -addr string
        The address to listen on; default is "" (all interfaces).
  -port int
        The port to listen on. (default 8000)
```


```
tcpl --port 8080
Listening on :8080.
2018/06/14 20:34:24 Accepted connection from [::1]:54901
00000000  47 45 54 20 2f 20 48 54  54 50 2f 31 2e 31 0d 0a  |GET / HTTP/1.1..|
00000010  48 6f 73 74 3a 20 6c 6f  63 61 6c 68 6f 73 74 3a  |Host: localhost:|
00000020  38 30 38 30 0d 0a 43 6f  6e 6e 65 63 74 69 6f 6e  |8080..Connection|
00000030  3a 20 6b 65 65 70 2d 61  6c 69 76 65 0d 0a 55 70  |: keep-alive..Up|
00000040  67 72 61 64 65 2d 49 6e  73 65 63 75 72 65 2d 52  |grade-Insecure-R|
00000050  65 71 75 65 73 74 73 3a  20 31 0d 0a 55 73 65 72  |equests: 1..User|
00000060  2d 41 67 65 6e 74 3a 20  4d 6f 7a 69 6c 6c 61 2f  |-Agent: Mozilla/|
```
