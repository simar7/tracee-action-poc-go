package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	//d, _ := os.MkdirTemp("", "")
	err1 := os.MkdirAll("/tmp/tracee2", 0755)
	fmt.Println("err1: ", err1)

	fi, err2 := os.Stat("/tmp/tracee2")
	fmt.Println("err: ", err2)
	fmt.Println("fi.isDir(): ", fi.IsDir())
}
