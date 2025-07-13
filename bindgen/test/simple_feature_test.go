package test

import (
	"context"
	"fmt"
	"io"
	"sort"
	"testing"
	"unsafe"
)

func TestBasicFeatures(t *testing.T) {
	buf := []byte("He ll o !")
	f := ZeroMallocTrimAllSpace(buf)
	fmt.Println(string(f))
	SortFloat64FastV1(a)
	fmt.Println(a)
	EachUtf8Runes("你好, 世界, 😒👌😘😁🙌")
	// handler := HandlerFunc(func(ctx context.Context, body io.ReadSeekCloser) {})
}

var a = []float64{4.1, 2.3, 5.1, 7.2, 2, 1, 88, 1}

type (
	HttpSeeker interface {
		ReadAll(r io.Reader) []byte
		Close(c io.Closer) error
	}
	HandlerFunc func(ctx context.Context, body io.Reader) []byte
)

// 为函数单独实现接口
var (
	_ HttpSeeker = (HandlerFunc)(nil)
)

// Close implements HttpSeeker.
func (h HandlerFunc) Close(c io.Closer) error {
	return c.Close()
}

// ReadAll implements HttpSeeker.
func (h HandlerFunc) ReadAll(r io.Reader) []byte {
	return h(context.TODO(), r)
}

func ZeroMallocTrimAllSpace(s []byte) []byte {
	// 相当于传递进来切片的引用, len=0, cap=cap(s)
	f := s[:0]
	// 会就地改变原切片
	for _, c := range s {
		if c != ' ' {
			f = append(f, c)
		}
	}

	return f
}

func SortFloat64FastV1(a []float64) {
	// 获取切片的原始引用(长度, 容量一致)
	// 转换为 [1 << 20]int 长度的数组指针
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}

func EachUtf8Runes(s string) {
	for i, r := range s {
		fmt.Printf("rune index: %d, rune is %c\n", i, r)
	}
}
