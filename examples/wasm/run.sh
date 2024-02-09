#!/bin/bash
GOOS=js GOARCH=wasm go build -o main.wasm
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'