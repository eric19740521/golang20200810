package main

import(
	"testing"
	//"fmt"
	"t001/util"
)

func Test_All(t *testing.T){	//注意,Tes後面,第一個若為字母,要大寫

	var d=0
	var (
		a=1
		b=2
		c=3
		
	)

   d = sum(a, b)	
   if d == c {
        t.Logf("sum(%d,%d)=%d ,期望值=%d,測試通過",a,b,d,c)
    } else {
        t.Errorf("sum(%d,%d)=%d ,期望值=%d,測試失敗",a,b,d,c)
    }
	
	
   d = Sum(a, b)	
   if d == c {
        t.Logf("Sum(%d,%d)=%d ,期望值=%d,測試通過",a,b,d,c)
    } else {
        t.Errorf("Sum(%d,%d)=%d ,期望值=%d,測試失敗",a,b,d,c)
    }	
	
   d = Add(a, b)	
   if d == c {
        t.Logf("Add(%d,%d)=%d ,期望值=%d,測試通過",a,b,d,c)
    } else {
        t.Errorf("Add(%d,%d)=%d ,期望值=%d,測試失敗",a,b,d,c)
    }	
	
}



func Test_Util(t *testing.T) {
	var (
		a=1
		b=2
		c=13
		
	)

   d := util.Ab(a, b)	
   if d == c {
        t.Logf("util.Ab(%d,%d)=%d ,期望值=%d,測試通過",a,b,d,c)
    } else {
        t.Errorf("util.Ab(%d,%d)=%d ,期望值=%d,測試失敗",a,b,d,c)
    }
}
