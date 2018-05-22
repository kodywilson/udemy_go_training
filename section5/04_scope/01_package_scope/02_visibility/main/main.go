package main

import (
	"fmt"

	"github.com/kodywilson/udemy_go_training/04_scope/01_package_scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
	vis.SuperCool()
}
