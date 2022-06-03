package main

import (
	"gift/boot"
	"gift/pkg"
)

func main() {

	err := boot.Run()
	if err != nil {
		panic(err)
	}

	pkg.Start()
}
