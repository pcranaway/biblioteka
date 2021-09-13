package main

import "github.com/pcranaway/biblioteka/env"
import "github.com/pcranaway/biblioteka/server"

func main() {
    // load environment
    environment := env.Load()

    // create server
    s := server.NewServer(environment)
    s.Start()
}
