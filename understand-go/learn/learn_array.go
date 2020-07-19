package learn

import "fmt"

func Print_Array() {
	var varra [10] int
	varra[0] = 1
	// 直接初始化数组
	var varrb = [3] int {1,2,3}
	varrb[2] = 4
	var vinta = varra[1]
	fmt.Println(vinta) // 进行输出对应的变量的值

	// 多维数组
	var arr [1][2] int
	arr[0][1] = 0
	// Go语言中对于数组和分片进行区分; array 和 slice
	// 在进行函数参数传递时，分片传送的是地址，数组传送的是数值；
}

func Slice() int {
	// go 语言支持切片形式的动态数组
	arr :=  [4] int {1,2,3,4}
	var vslicea [] int =  arr[:3]
	//var vsliceb [] int= make([]int, 10)
	var vslicec [] int = make([]int, 10, 20) // length 表示切片的初始长度， capacity表示数组容量
	fmt.Println(len(arr), cap(arr))
	fmt.Println(len(vslicec), cap(vslicec))
	vslicea = append(vslicea, 2, 3, 4)
	fmt.Println(vslicea)
	copy(vslicec, vslicea)
	fmt.Println(vslicec)
	// foreach 和 range 结合， range关键字可用于 for循环，slice, channel 和 map集合元素
	for i, v := range vslicec{
		fmt.Println(i,v)
	}
	return 0
}