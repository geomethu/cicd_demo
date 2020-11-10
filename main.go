package main

import (
	"fmt"

	"github.com/geomethu/cicd_demo/addition"
)


func main(){
	sum:= addition.Addition(40.5, 3.5)
	fmt.Println(sum)
}