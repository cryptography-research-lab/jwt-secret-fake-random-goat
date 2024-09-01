package main

import "github.com/cryptography-research-lab/jwt-secret-fake-random-goat/pkg/web"

func main() {
	port := 10086
	web.Run(port)
}
