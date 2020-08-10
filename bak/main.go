package main

import(
	"fmt"
	//"t001/util"
)



func Sum(a,b int) int{
 
	c:=a+b
	return c
}
func sum(c,d int) int{


	return c+d
}

func main(){


	//c:=util.Ab(1,2)

	//fmt.Println("c=",c)
	fmt.Println("回傳值=",sum(1,4))	//main.go
	
	
	fmt.Println("回傳值=",Add(11,22))	//main2.go ,go build main.go會錯誤
	
	
	
}