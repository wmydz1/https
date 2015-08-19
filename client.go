package main
import (

    "crypto/x509"
    "io/ioutil"
    "fmt"
    "net/http"
    "crypto/tls"
)

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"
    caCrt, err := ioutil.ReadFile(caCertPath)
    if err !=nil {
        fmt.Println("ReadFile err", err)
        return
    }
    pool.AppendCertsFromPEM(caCrt)

    tr := &http.Transport{
        TLSClientConfig:&tls.Config{RootCAs:pool},
    }
    client := &http.Client{Transport:tr}
    resp, err := client.Get("https://localhost:10000")
    if err !=nil {
        fmt.Println("Get error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

}