
package main

import (
    "fmt"
    "net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
	// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

}
