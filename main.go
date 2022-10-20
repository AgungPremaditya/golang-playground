package main

import "movies-golang-api/routers"

func main() {
	PORT := ":8080"

	routers.StartServer().Run(PORT)
}
