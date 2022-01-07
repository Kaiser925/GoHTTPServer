# fileserve

fileserve is a simple file server.

Usage:

install

```bash
$ go install github.com/Kaiser925/fileserve@latest
```

```bash
$ fileserve --help

Usage of ./fileserve:
  -bind string
        bind address (default "127.0.0.1")
  -directory string
        root directory (default "./")
  -port int
        bind port (default 8765)
```

Run in docker

```bash
$ docker run -v $YOUR_DATA_DIR:/data -p $PORT:8765 tricker1996/simple-http-server  
```
