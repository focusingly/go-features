package test

import (
	"bindgen"
	"testing"
)

func TestFuncExport2CGO(t *testing.T) {
	bindgen.PrintGOFuncToCExtern()
	bindgen.PrintCStructInGO()
}
