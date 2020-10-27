# Installation
Before running the Go server locally, you need to download Go on your OS.

On Ubuntu you easilly install Go using:
```sh
$ sudo snap install go --classic
```

Windows and MacOS instructions can be found [here](https://golang.org/doc/install).

After installing Go, check its version:
```sh
$ go version
1.15.3 linux/amd64
```
# Run the server
First cd to the main directory:
```sh
$ cd news-sentiment/go-server/main
```

Run the server:
```sh
$ go run main.go
```

Check the server is running by pinging it:
```sh
$ curl localhost:8090/hello
Hello world!
```

Done!
