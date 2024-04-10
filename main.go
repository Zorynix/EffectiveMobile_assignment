package main

import (
	"flag"

	"tz.com/m/routes"
	"tz.com/m/utils"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {

	flag.Parse()
	utils.InitLogger()
	routes.Routes(addr)

}
