package main

import "api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDrive)
}
