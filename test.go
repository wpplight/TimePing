package main

import "fmt"


func main(){
	var i uint16
	i=1025
	b:=make([]byte,2)
	b[0]=byte(i>>8)
	b[1]=byte(i)
	fmt.Println(b)
	fmt.Println(uint16(b[0])<<8 | uint16(b[1]))
}