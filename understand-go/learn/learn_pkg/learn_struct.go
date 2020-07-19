package learn_pkg

import (
	"fmt"
)

type Person struct {
	name string // 结构体成员为name，类型为string
}

func Print_Struct() int {
	// 作为函数参数时; 进行传指针需要使用 × 操作符
	fmt.Println(Person{name: "yaowenxu"})
	var xiaom Person
	xiaom.name = "xiaoming"
	return 0
}

func (p Person) Vmap() int {
	var vmapa = make(map[string]int)
	vmapa["a"] = 1;
	fmt.Println(vmapa)
	delete(vmapa, "a") // 进行删除vmapa中的对象;
	fmt.Println(vmapa)
	return 0 // 进行返回值
}