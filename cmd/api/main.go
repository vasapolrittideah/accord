package main

import "github.com/vasapolrittideah/accord/server"

const VERSION = "1.0.0"

func main() {
	server.NewServer(VERSION).Run()
}
