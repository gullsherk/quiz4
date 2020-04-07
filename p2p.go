package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
	argument := os.Args
    listener, err := net.Listen("tcp", "127.0.0.1:"+argument[1])

    if err != nil {
        log.Fatal("tcp server listener error:", err)
    }

    //conn, err := net.Dial("tcp", "localhost:"+argument[1])
    //if err != nil {
    // handle error
    //}
    //fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    fmt.Println("The server is dialing on port : "+argument[2])

    fmt.Println("The server is listening on port : "+argument[1])
    fmt.Println("The server can handle multiple requests")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("tcp server accept error", err)
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

    if err != nil {
        log.Println("client left..")
        conn.Close()
        return
    }
    message := string(bufferBytes)
    clientAddr := conn.RemoteAddr().String()
    response := fmt.Sprintf(message + " from " + clientAddr + "\n")

    log.Println(response)



    handleConnection(conn)
}