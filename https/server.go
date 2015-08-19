package main
import (

    "net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hi,this is golang https service")
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":10000", "server.crt", "server.key", nil)
}
