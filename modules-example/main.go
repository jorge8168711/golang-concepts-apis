package main

import (
	"fmt"

	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
	utils "github.com/jorge8168711/golang-concepts-apis/my-module-test"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Some hello")

	fmt.Println(utils.HelloMyModule())
}
