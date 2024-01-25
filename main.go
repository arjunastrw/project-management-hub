package main

import (
	delivery "enigma.com/projectmanagementhub/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
