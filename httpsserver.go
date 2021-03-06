package main
import (
    "net"
    "log"
    "io"
    "crypto/tls"
    "time"
    "crypto/rand"
)

func handleClient(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 512)
    for {
        log.Print("server: conn :waiting")
        n, err := conn.Read(buf)
        if err !=nil {
            if err != io.EOF {
                log.Printf("server:conn: read %s", err)
            }
            break
        }
        log.Printf("server:conn : echo %q\n", string(buf[:n]))
        n, err =conn.Write(buf[:n])
        log.Printf("server: conn :wrote %d bytes", n)
        if err !=nil {
            log.Printf("server: write : %s", err)
            break
        }

    }
    log.Println("server: conn:closed")
}


func main() {
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err !=nil {
        log.Fatalf("server : loadkeys :%s", err)
    }
    config := tls.Config{Certificates:[]tls.Certificate{cert}}
    config.Time =time.Now
    config.Rand =rand.Reader

    service := "127.0.0.1:10000"
    listener, err := tls.Listen("tcp", service, &config)
    if err !=nil {
        log.Fatalf("server : listen: %s", err)
    }
    log.Print("server: listening")
    for {
        conn, err := listener.Accept()
        if err !=nil {
            log.Panicf("server:accept :%s", err)
            break
        }
        log.Printf("server : accepted from %s", conn.RemoteAddr())
        go handleClient(conn)
    }
}