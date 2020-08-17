// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hello is a trivial package for gomobile bind example.
package hello

import (
	"Java/android/app"
	"Java/android/os"
	"Java/android/widget"
	jpkg "Java/hello"
	"Java/hello/RBinding"
	"Java/java/lang"
	"Java/java/lang/Object"
	"Java/java/lang/System"
	rid "Java/org/golang/example/bind/R/id"
	rlayout "Java/org/golang/example/bind/R/layout"
	"fmt"
	"log"
	"unsafe"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type Counter struct {
	Value int
}

func (c *Counter) Inc() { c.Value++ }

func NewCounter() *Counter { return &Counter{5} }

type Printer interface {
	Print(s string)
}

func PrintHello(p Printer) {
	p.Print("Hello, World!")
}

type GoObject struct {
	lang.Object
}

func (o *GoObject) ToString(this jpkg.GoObject) string {
	return this.Super().ToString()
}

func NewGoObject() *GoObject {
	return &GoObject{}
}

func NewObject() lang.Object {
	//return lang.Object.New() // not work
	return Object.New()
}

type GoRBinding struct {
	jpkg.RBinding

	Age int

	// Age int32 // gomobile didinot support `int32 as constructor param`
	// gen.java explain:
	// Skip constructors taking a single int32 argument
	// since they clash with the proxy constructors that
	// take a refnum
 }

func NewGoRBinding(age int) *GoRBinding {
	return &GoRBinding{Age: age}
}

func JavaCall() {
	t := System.CurrentTimeMillis()
	log.Println("Init", "[Test] Called java function return value is: ", t)
}

func GetStringFromJavaByGo() string {
	return RBinding.GetStringFromJava()
}

type MainActivityByGo struct {
	app.Activity

	//Tv widget.TextView
}

func (o *MainActivityByGo) ToString(this jpkg.MainActivityByGo) string {
	return this.Super().ToString()
}

func (o *MainActivityByGo) OnCreate(this jpkg.MainActivityByGo, b os.Bundle) {
	this.Super().OnCreate(b)
	RBinding.SetContentView(this, rlayout.Activity_main)
	var tv widget.TextView = RBinding.FindTextViewById(this, rid.Mytextview)
	greetings := Greetings("Android and Gopher")
	RBinding.SetText(tv, greetings)

	//RBinding.M_I(9876)
	//RBinding.M_Ljava_lang_String_2(greetings)
	var i32 int32 = 9876
	RBinding.M(i32)
	var i64 int64 = 987654321
	RBinding.M(i64)
	var i64_2 int = 987654321
	// maybe the `gomobile`'s bug
	// it will get wrong on android platform although `int` and `int64` has the same bit size -_-||
	RBinding.M(i64_2)
	RBinding.M(greetings)

	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))
}

func (o *MainActivityByGo) OnDestroy(this jpkg.MainActivityByGo) {
	this.Super().OnDestroy()
}

func FloatDoubleValue(f lang.Float) float64 {
	return f.DoubleValue()
}
