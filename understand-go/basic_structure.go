// 包名
// 只有包名为main的文件，才能包含main函数
package main

// 导入其他包
import "fmt"
import "./learn"
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
type newstruct struct {

}

// 接口的声明
type newinterface interface {

}

// 由main函数作为程序入口点启动；
func main(){
	fmt.Println("基础语法学习")
	learn.Print()
}