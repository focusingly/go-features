package test

import (
	"bindgen"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicPef(b *testing.B) {
	b.Run("CgoFastAtomicU64", func(b *testing.B) {
		for b.Loop() {
			bindgen.GetCGONextID()
		}
	})

	b.Run("GoNativeAtomicU64", func(b *testing.B) {
		var a atomic.Uint64

		for b.Loop() {
			a.Add(1)
		}
	})
}
