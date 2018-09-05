package cgotest

// #include "cfunc.h"
import "C"

import (
	"time"
	"unsafe"
)

//export GoLog
func GoLog(msg *C.struct_msg) {
	ch <- msg
}

type msg struct {
	time time.Time
	body string
}

//export GoLog2
func GoLog2(cm *C.struct_msg) {
	m := msg{
		time.Unix(int64(cm.time.tv_sec), int64(cm.time.tv_nsec)),
		C.GoString(cm.body),
	}

	ch2 <- m
}

//export GoLog3
func GoLog3(msg *C.struct_msg) {
	ch3 <- msg
}

//export GoNop
func GoNop() {}

var count = 0

func CallCFunction() {
	C.logger(nil)
}

func CallCFunction2() {
	C.logger2(nil)
}

func CallCFunction3() {
	C.logger3(nil)
}

func CallCFunctionNop() {
	C.nop()
}

func CallCFunctionNop2() {
	C.gonop()
}

var ch = make(chan *C.struct_msg, 10000)
var ch2 = make(chan msg, 10000)
var ch3 = make(chan *C.struct_msg, 10000)

func init() {
	go func() {
		for _ = range ch {
		}
	}()

	go func() {
		for _ = range ch2 {
		}
	}()

	go func() {
		for m := range ch3 {
			C.free(unsafe.Pointer(m.body))
			C.free(unsafe.Pointer(m))
		}
	}()
}
