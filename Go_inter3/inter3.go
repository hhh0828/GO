package main

import (
	"fmt"
)

// reader interface
type Remotecontrol interface {
	Readwriter
	Closer
}

type Writer interface {
	Write()
}

type Reader interface {
	Read()
}

// intergrated interface.
type Readwriter interface {
	Writer
	Reader
}

// closer interface
type Closer interface {
	Close()
}

type Info struct {
	modelname string
	madedate  string
	madeby    string
}

type Audio struct {
	Info
	Bios   string
	serial string
}
type TV struct {
	Info
	Bios   string
	serial string
}

func (t TV) Write() {
	fmt.Println("TV Writer")
}

func (aud Audio) write() {
	fmt.Println("Audio writer")
}

func (t *TV) Read() {
	fmt.Println("TV Read")
}
func (aud *Audio) Read() {
	fmt.Println("Audio Read")
}

//인터페이스에서 예외처리가 가능한가...
/*
func (t TV) Close() {
	fmt.Println("TV close")
}
*/
func (aud Audio) Close() {
	fmt.Println("Audio Close")
}

func Readfile(reader Reader) {
	//interface nil exception.
	/*
		c := reader.(Closer)
		c.Close()
	*/

	if c, ok := reader.(Closer); ok {
		c.Close()
	} else {
		fmt.Println("cannot run")
	}
}

func main() {
	/*
		var Remotecontrol Remotecontrol = TV{}
		Remotecontrol.Close()
		Remotecontrol.Read()
	*/
	Readfile(&TV{})
}
