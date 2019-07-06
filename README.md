# go-wasm

WebAssembly implementation by Golang.  
[Official tutorial](https://github.com/golang/go/wiki/WebAssembly)

## build

```
GOOS=js GOARCH=wasm go build -o main.wasm
```

## run
```go
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```