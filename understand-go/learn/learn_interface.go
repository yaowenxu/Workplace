package learn

import (
	"errors"
	"fmt"
)

func Print_Interface() int {
	var phone Phone
	phone = new(Nokia)
	phone.call()
	phone = new(IPhone)
	phone.call()
	_, e := check(1)
	fmt.Println(e)
	return 0
}

func check(cond int	) (int, error) {
	if cond < 0 {
		return 0, errors.New("negative")
	}
	return 0,errors.New("")
}

type Phone interface {
	call() int
}

type Nokia struct {

}

func (nokia Nokia) call() int {
	fmt.Println("nokia call")
	return 0
}

type IPhone struct {

}

func (iphone IPhone) call() int {
	fmt.Println("iphone call")
	return 1
}