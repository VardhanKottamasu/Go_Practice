package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var a1 [...] int 
	a1:= [...] int {1,2,3,4,5,6}
	// arr := [...]string{"a","b","c","d","e","f"}
	for i:=0;i<len(a1);i++ {
		fmt.Println(a1[i])
	}

	arr1:=append(a1[0:],1,2)
	fmt.Println(arr1)

	v1:="1"
	ans,i:=strconv.Atoi(v1)
	fmt.Println(ans)
	print(i)


}