<<<<<<< HEAD
package main

import (
	delivery "enigma.com/projectmanagementhub/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
=======
package main

import (
	delivery ""

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
>>>>>>> 4e0c22da47a4350c832b379f60379ec7e899024d
