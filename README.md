# Go TCP price query server

Insert timestamp + price, query average price between range of times

https://protohackers.com/problem/2

### run the tcp server
```
go build -o server/main server/main.go
./server/main
```

### test with a client
```
go build -o client/main client/main.go
./client/main
```

## protocol

```
In this example, "-->" denotes messages from the server to the client, and "<--" denotes messages from the client to the server.

    Hexadecimal:                 Decoded:
<-- 49 00 00 30 39 00 00 00 65   I 12345 101
<-- 49 00 00 30 3a 00 00 00 66   I 12346 102
<-- 49 00 00 30 3b 00 00 00 64   I 12347 100
<-- 49 00 00 a0 00 00 00 00 05   I 40960 5
<-- 51 00 00 30 00 00 00 40 00   Q 12288 16384
--> 00 00 00 65                  101
```