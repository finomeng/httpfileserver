//License: MIT

/*
usage:
$go run httpserver.go
$go run httpserver.go -p 8888 -d /tmp
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var g_port = 9000
var g_pubdir = "."

func init() {
	portFlag := flag.Int("p", g_port, "optional, port number; default: "+strconv.Itoa(g_port))
	pubFlag := flag.String("d", g_pubdir, "optional, dir to publish; default: "+g_pubdir)
	flag.Parse()
	g_port = int(*portFlag)
	g_pubdir = string(*pubFlag)
}

func main() {
	fmt.Printf("Publishing directory \"%s\" at port %v:\n", g_pubdir, g_port)
	http.Handle("/", http.FileServer(http.Dir(g_pubdir)))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(g_port), nil))
}
