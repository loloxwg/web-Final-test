package main

import (
	_ "gin_demo/boot"
	"gin_demo/router"
)

func main() {
	router.Init()
}
