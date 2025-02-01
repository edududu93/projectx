package main

import (
	//"fmt"
	"projectx/network" // YT: "github.com/anthdm/projectx/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")

	opts := network.ServerOpts {
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}