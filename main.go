package main

import (
	"flag"

	_ "tz.com/m/docs"
	"tz.com/m/routes"
	"tz.com/m/utils"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

// @title Car Management API
// @version 1.0
// @description This API allows you to manage cars in a database.
// @host localhost:8000
// @BasePath /v1
func main() {

	flag.Parse()
	utils.InitLogger()
	routes.Routes(addr)

}
