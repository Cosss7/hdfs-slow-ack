# hdfs-slow-ack
排查HDFS Slow ack问题的最小复现程序

## BUILD
```go
go build
```

## USAGE
```bash
# server
./hello s > out.txt 2>&1
# client
./hello c SERVER_HOST > out.txt 2>&1
```

