package test

import (
	"bindgen"
	"fmt"
	"testing"
)

const (
	arrLen = 4096
)

func BenchmarkCGOAndSimpleGOMat(b *testing.B) {
	b.Run("CgoAVX256", func(b *testing.B) {
		mat1 := make([]float32, arrLen)
		mat2 := make([]float32, arrLen)
		result := make([]float32, arrLen)
		for i := range mat1 {
			mat1[i] = 1.2
			mat2[i] = 1.2
		}
		for b.Loop() {
			bindgen.AVX2MatAddWithRecv(mat1, mat2, result)
		}
	})
	b.Run("NativeGO", func(b *testing.B) {
		mat1 := make([]float32, arrLen)
		mat2 := make([]float32, arrLen)
		result := make([]float32, arrLen)
		for i := range mat1 {
			mat1[i] = 1.2
			mat2[i] = 1.2
		}
		for b.Loop() {
			if len(mat1) != len(mat2) && len(mat2) != len(result) {
				panic("")
			}
			for i := range len(mat1) { // 遍历行
				result[i] = mat1[i] + mat2[i]
			}
		}
	})
	b.Run("NativeGOWith8Wrap", func(b *testing.B) {
		mat1 := make([]float32, arrLen)
		mat2 := make([]float32, arrLen)
		result := make([]float32, arrLen)
		for i := range mat1 {
			mat1[i] = 1.2
			mat2[i] = 1.2
		}
		for b.Loop() {
			if len(mat1) != len(mat2) && len(mat2) != len(result) {
				panic(fmt.Errorf("len(a) != len(b) != len(result)"))
			}
			for i := 0; i+7 < arrLen; i += 8 {
				result[i+0] = mat1[i+0] + mat2[i+0]
				result[i+1] = mat1[i+1] + mat2[i+1]
				result[i+2] = mat1[i+2] + mat2[i+2]
				result[i+3] = mat1[i+3] + mat2[i+3]
				result[i+4] = mat1[i+4] + mat2[i+4]
				result[i+5] = mat1[i+5] + mat2[i+5]
				result[i+6] = mat1[i+6] + mat2[i+6]
				result[i+7] = mat1[i+7] + mat2[i+7]
			}
			for i := arrLen - (arrLen % 8); i < arrLen; i++ {
				result[i] = mat1[i] + mat2[i]
			}
		}
	})
}
