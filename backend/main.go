package main

import (
	"gift/boot"
)

func main() {

	err := boot.Run()
	if err != nil {
		panic(err)
	}

}
