# gotools
Various small command-line utilities written in Go

## goexec 
Execute command with timeout

```
goexec -h
Usage of goexec:
  -command string
        Command
  -timeout duration
        Command timeout (default 1m0s)
  -verbose
        Verbose output
```


```
goexec --timeout=5s --command=sleep -- 10
Command timed out
```
