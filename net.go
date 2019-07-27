//__author__ = "YaoYao"
//Date: 2019-07-23
package main

import (
	"fmt"
	"math/rand"
)

type I interface {
	Test(s string)
	Hello(s string)
}

type T struct {
}

func (t *T) Test(s string) {
	fmt.Println("Test")
}

func (t *T) Hello(s string) {
	fmt.Println("Hello")
}

func II() {
	var n int
	n = rand.Intn(100)
	fmt.Println(n)
	//s := []func(i I, s string){I.Test, I.Hello}
	//t := &T{}
	//for _, f := range s {
	//	f(t, "hello")
	//}
}
