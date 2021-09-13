package main

import "github.com/pcranaway/biblioteka/server"

func main() {
    s := server.NewServer()
    s.Start()
}
