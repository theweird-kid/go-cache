package main

import "github.com/theweird-kid/go-cache/cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}
	server := NewServer(opts, cache.NewCache())
	server.Start()
}
