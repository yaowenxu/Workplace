package learn

import "fmt"

// go 语言中，使用大小写，来决定常量，变量，类型，接口，结构或者函数是否可以被外部包调用；
func Print() {
	fmt.Println("导入Learn 包")
	version()
}

func version(){
	fmt.Println("version 1.0.0")
}