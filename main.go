package main

import(
	"fmt"
 
)



func Sum1(a,b int) int{
 
	c:=a+b
	return c
}
/* func Sum2(c,d int) int{


	return c+d
}
 */
func main(){


	//c:=util.Ab(1,2)

	//fmt.Println("c=",c)
	fmt.Println("回傳值=",Sum1(1,4))	//main.go
	//fmt.Println("回傳值=",Sum2(1,4))	//main.go
	
	//fmt.Println("回傳值=",Add(11,22))	//main2.go ,go build main.go會錯誤
	
	
	
}