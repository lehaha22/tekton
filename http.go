package main

import (
        "fmt"
        "net/http"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello golang!")
}

func main() {
        http.HandleFunc("/", sayHello)
        fmt.Println("start 127.0.0.1:8080...")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                fmt.Printf("http server failed, err:%v\n", err)
                return
        }
}
