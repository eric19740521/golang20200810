package util

import(
	"testing"
	"fmt"
)



func TestAB(t *testing.T) {
	var tName string

	type args struct{
		arg1 int	//第一個參數
		arg2 int	//第二個參數
	
	}
	type targs struct{			////定義 測試參數 struct 
		 
		args args	//參數
		act int		//回傳值
		exp int		//期待值	
	
	
	}

	var nums map[int]targs =make(map[int]targs,3) //型態為struct slice
	//var nums []targs = make([]targs, 1)	//型態為struct slice
	
	
	nums[0] = targs{		//定義 測試參數, 
		 
		args:args{
			arg1:1,
			arg2:2,
		},
		act:3,
		exp:3,
	}
	nums[1] = targs{		//定義 測試參數, 
		 
		args:args{
			arg1:3,
			arg2:2,
		},
		act:5,
		exp:5,
	}
	nums[2] = targs{		//定義 測試參數, 
		 
		args:args{
			arg1:11,
			arg2:33,
		},
		act:44,
		exp:44,
	}	
	

	tName="Sum函數"
	fmt.Printf("測試數據有 %d組\n",len(nums))
	for _,v := range nums {
	
	
		//fmt.Printf("%T",v)
		v.act = Sum(v.args.arg1,v.args.arg2)		//小寫函數,也能測試
		
		if v.act==v.exp {
			t.Logf("%s=>測試成功:回傳值=%d,期望值=%d",tName,v.act,v.exp)
		} else {
		
			t.Errorf("%s=>測試失敗:回傳值=%d,回傳值=%d",tName,v.act,v.exp)
		}		
	}

}