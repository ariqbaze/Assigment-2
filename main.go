package main

import (
	"Assigment-2/routers"
)

func main() {

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
