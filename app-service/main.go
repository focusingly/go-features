package main

import (
	"bindgen"
)

func main() {
	bindgen.PrintGOFuncToCExtern()
	bindgen.PrintCStructInGO()
}
