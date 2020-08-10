# golang&testing單元測試

#### 參考資料:
>https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html  
>https://ithelp.ithome.com.tw/articles/10204692



# package main
1. func 在同一個main.go,
   go build main.go 或是go build  
2. func 在同一個main.go,跟main2.go 
   go build 
 
func名稱 大小寫是為不同函數



testing 为 Go 语言 package 提供自动化测试的支持。通过 go test 命令，
能够自动执行如下形式的任何函数：

func TestXxx(*testing.T)
注意：Xxx 可以是任何字母数字字符串，但是第一个字母不能是小写字母。


## 1.package main 測試

 1)func 在同一個main.go,
   go build main.go 或是go build  
 2)func 在同一個main.go,跟main2.go 
   go build 

   go test -v 	

   go test -v -cover=true      測試涵蓋率為何不是100%	

   ==> package u001	


## 2.package util包 測試

go test -v 
go test -v -cover=true      測試涵蓋率




### git github上傳 操作
1. 在github建立倉庫 golang20200810
2. 電腦建立資料夾 
   git clone 
   https://github.com/eric19740521/golang20200810.git

3. cd golang20200810
	開始寫程式

4. git add .
5. git status 查看加入要上傳的的檔案

6. git commit -m "init commit" 提交上傳

7. git push 上傳吧

### 如果檔案有修改後,才上傳
1. git add readme.md
2. git commit -m "修改"
3. git push

