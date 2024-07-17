package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/theweird-kid/go-cache/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", opts.ListenAddr)
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET Foo Bar 250000000000"))

		time.Sleep(time.Second * 2)
		conn.Write([]byte("GET Foo"))

		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))

	}()

	server := NewServer(opts, cache.NewCache())
	server.Start()
}
