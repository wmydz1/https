package main

//  生成秘钥
// openssl genrsa -out server.key 2048
//openssl req -new -x509 -key server.key -out server.crt -days 365
import (

    "net/http"
    "fmt"
    "log"
)

const (
    SERVER_PORT = 8080
    SERVER_DOMAIN = "127.0.0.1"
    RESPONSE_TEMPLATE = "hello"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
    w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
    http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
    err := http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "server.crt", "server.key", nil)
    if err!=nil {
        log.Fatal(err)
    }

}