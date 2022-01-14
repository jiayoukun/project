package main

import (
	"test10/conf"
	"test10/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)

}
