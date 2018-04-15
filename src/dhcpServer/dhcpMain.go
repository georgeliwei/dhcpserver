package main

import "fmt"

import "dhcpconfig"

func main() {

	var x dhcpconfig.TdhcpCfg
	err := x.Init("test.txt")
	fmt.Println(x, err)
}
