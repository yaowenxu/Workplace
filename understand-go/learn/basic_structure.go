// 包名
// 只有包名为main的文件，才能包含main函数
package main

// 导入其他包
import (
	"fmt"
	"main/learn_pkg"
)

//import (
//	"math"
//	"io"
//	)

// 定义常量
const PI = 3.14

// 定义全局变量声明和赋值
var name = "go"

// 一般类型的声明
type new int

// 结构体的声明
type newstruct struct {}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

// 接口的声明
type newinterface interface {}
var vglobal int = 3 // 全局变量
// 由main函数作为程序入口点启动；
func main(){
	var va, vb int = 1, 2 // 局部变量
	var ptr *int
	ptr = nil
	fmt.Println(va, vb) // 进行输出值
	fmt.Println("基础语法学习")
	learn_pkg.Print()
	learn_pkg.Print_Array()
	fmt.Println("指针地址为：", &va)
	fmt.Printf("数值 %x\n", ptr)
	learn_pkg.Print_Struct()
	learn_pkg.Slice()
	var VPersiona learn_pkg.Person
	fmt.Println("OK")
	VPersiona.Vmap() // 调用结构体私有函数
	for i := 0; i < 10; i++{
		fmt.Println(fib(i))
	}
	// 类型转换
	var vsum int = 10
	var vfload float32 = 1.0
	vfload = float32(vsum)
	fmt.Println(vfload)
	learn_pkg.Print_Interface()
	learn_pkg.Print_Goroutine()
}