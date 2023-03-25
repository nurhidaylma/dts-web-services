package main

import "github.com/nurhidaylma/dts-web-services/routers"

func main() {
	var PORT = ":8080"

	routers.StartServe().Run(PORT)
}
