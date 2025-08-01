package bindgen

/*
#cgo CFLAGS: -march=native -mavx2 -I${SRCDIR}/lib
#cgo LDFLAGS: -static-libgcc -static-libstdc++ -lpthread
#include "mat.c"
#include "call.c"
#include "counter.c"
*/
import "C"
import (
	"fmt"
	"log"
	"sync"
	"unsafe"
)

type (
	ObjectID int32
)

var (
	Refer struct {
		sync.RWMutex
		objMap    map[ObjectID]any
		nextObjID ObjectID
	}
)

func init() {
	Refer.Lock()
	defer Refer.Unlock()
	Refer.nextObjID = 1000
	Refer.objMap = map[ObjectID]any{}
}

func NewObjectID(obj any) ObjectID {
	Refer.Lock()
	defer Refer.Unlock()
	Refer.nextObjID++
	Refer.objMap[Refer.nextObjID] = obj
	return Refer.nextObjID
}

func GetObj(id ObjectID) (obj any, ok bool) {
	Refer.RLock()
	defer Refer.RUnlock()
	obj, ok = Refer.objMap[id]

	return
}

func Release(id ObjectID) any {
	Refer.Lock()
	defer Refer.Unlock()
	obj := Refer.objMap[id]
	delete(Refer.objMap, id)

	return obj
}

//export printSomething
func printSomething() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			var space = ""
			if i != j {
				space = "\t"
			}
			fmt.Printf("%d * %d = %d%s", j, i, i*j, space)
		}

		fmt.Println()
	}
}

//export cgoBridgeHandler
func cgoBridgeHandler(handlerID C.int, user C.User) {
	if obj, ok := GetObj(ObjectID(handlerID)); ok {
		if handler, ok := obj.(func(C.User)); ok {
			handler(user)
		} else {
			log.Println("can not cast handler to matched type")
		}
	}
}

func PrintGOFuncToCExtern() {
	C.printSomething()
}

func newUser(name string) *C.user {
	u := (*C.user)(C.malloc(C.size_t(unsafe.Sizeof(C.user{}))))
	u.username = C.CString(name)
	u.age = C.int(30)
	u.gender = C.int(1) // male

	return u
}

type User struct {
	username string
	age      int
	male     int
}

func PrintCStructInGO() {
	u := newUser("Jon")
	defer func() {
		C.free(unsafe.Pointer(u.username))
		C.free(unsafe.Pointer(u))
	}()
	u2 := C.print_user(u)
	tmp := User{
		username: C.GoString(u2.username),
		age:      int(u2.age),
		male:     int(u.gender),
	}
	fmt.Println(tmp)
	// cleanup before
	C.free(unsafe.Pointer(u.username))
	// setup new
	u.username = C.CString("Fox")

	objID := NewObjectID(func(user C.User) {
		fmt.Println("------------- C TO GO START---------------->")
		tmp2 := User{
			username: C.GoString(u2.username),
			age:      int(u2.age),
			male:     int(u.gender),
		}
		fmt.Println(tmp2)
		fmt.Println("------------- C TO GO END---------------->")
	})

	C.callFromGo(C.int(objID))
}

func AVX2MatAdd(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic(fmt.Errorf("len(a) != len(b)"))
	}

	out := make([]float32, len(a))
	C.mat_add(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&out[0])),
		C.int(len(out)),
	)

	return out
}

func AVX2MatAddWithRecv(a, b, result []float32) {
	if len(a) != len(b) || len(a) != len(result) {
		panic(fmt.Errorf("len(a) != len(b) != len(result)"))
	}

	C.mat_add(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&result[0])),
		C.int(len(result)),
	)
}

func AVX2MatSub(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic(fmt.Errorf("len(a) != len(b)"))
	}
	out := make([]float32, len(a))

	C.mat_sub(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&out[0])),
		C.int(len(out)),
	)
	return out
}

// GetCGONextID 基于 CGO 获取原子递增 ID
func GetCGONextID() uint64 {
	return uint64(C.next_id())
}
