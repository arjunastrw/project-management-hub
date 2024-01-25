package main

import (
	delivery ""

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
