package main

import (
	"fmt"
	"log"
	"net/http"

	"unipro-proxy/internal/proxy"
)

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/", proxy.NewTcpProxy())
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("listen err:%s", err)
	}
}
