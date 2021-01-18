package main

import (
	"context"

	"github.com/mrbenosborne/go-http-cache/server"
)

func main() {
	s := server.New()
	s.Serve(context.Background())
}
