package main

import (
	"fmt"
	"ltools/util/enum"
)

func main() {
	fmt.Println(enum.StatusAuthFaild)
	fmt.Println(enum.StatusText(enum.StatusAuthFaild))
}
