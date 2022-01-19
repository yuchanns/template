package utils

import "fmt"

func SafeAsync(fn func()) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("panic: %+v", msg)
		}
	}()
	fn()
}

func SafeAsyncFunc(fn func()) func() {
	return func() {
		SafeAsync(fn)
	}
}
