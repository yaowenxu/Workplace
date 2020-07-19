package learn
// 学习go语言包组织功能

import "fmt"

// go 变量声明, 布尔值，整型，浮点型，字符串; 四种主要基本数据类型，其拷贝是传值拷贝；
var vboolt bool = true
var vboolf bool = false
var vint1 int = 1
var vfloat float32 = 1.0
var vinta, vintb = 2, 3 // 当不写类型的时候，编译器可以根据后面赋值类型进行推断；
var vint0 int
var vstringa string = "a string"

// go 语言常量声明
const vconsta bool = true
const vconstb int = 1
const vconstc = "str"
const (
	i = iota
	j = iota
	x = iota
)

// 算术运算符 + - × / % ++ --
// 关系运算符 == != > < >= <=
// 逻辑运算符 && || ！
// 位运算符 & | ^ << >>
// 赋值运算符 = += -= ×= /= %= <<= >>= &= ^= |=
// 地址运算符 & * , 用于指针操作；
// 针对于运算符优先级，最好使用括号操作来保证代码的稳定性和可读性；

// go 语言中，使用大小写，来决定常量，变量，类型，接口，结构或者函数是否可以被外部包调用；
func Print() {
	fmt.Println("导入Learn 包")
	fmt.Println(vstringa)
	fmt.Println(vint0)
	vinta, vintb, _ =  numbers()
	fmt.Println(vinta, vintb)
	version()
}

func version(){
	fmt.Println("version 1.0.0")
}

func numbers() (int, int, string) {
	a, b, c :=  1, 2, "3"
	return a,b,c
}

func condition() int {
	var cond bool = true
	var vval int = 1
	var vinterface interface {}
	switch vval { // switch 后面不用添加break了;
	case 1:
		break
	case 2:
		fallthrough // 执行后面的句子
	case 3:
	default:
		return 6
	}
	switch vinterface.(type) {
	case int:
	case bool:
	case string:
	default:
		return 7
	}
	if cond {
		return 1
		if cond {
			return 4
		}
	}
	if cond {
		return 2
	} else if cond {
		return 5
	}else {
		return 3
	}
	// select 操作，真对于channel 使用;
}
// 函数，包括 函数名称，返回值，和参数类型 函数声明 函数体
func loop(sum int) (int, int) {
	sum = 0
	strings := []string{"google", "windows"}
	for i, s := range strings{ // for-each 循环类型
		fmt.Println(i, s)
	}
	for i := 1; i < 10; i++ {
		sum++
	}
	for sum < 20 {
		sum++
	}
	for true {
		sum++
		for true {
			sum++
			if sum<100 {
				break;
			}
		}
		break
	}
	return 0, 1
}

// 传值和传引用的区别
func swap(x *int, y *int){
	// 函数变量
	afunc := func() int{
		return 0
	}
	fmt.Println(useafunc(afunc)) // 调用函数并输出值
}

func useafunc(f func() int) func() int {
	i := 0
	// 函数闭包
	return func () int {
		i++
		fmt.Println("b")
		return f()
	}
}

type Circle struct {
	r float32
}

// 结构体的方法
func (c Circle) get() float32 {
	return 3.14*c.r*c.r
}