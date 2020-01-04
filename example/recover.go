package main

import (
	"errors"
	"fmt"
	"time"
)

func Unmarshal() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}

	}()

	a()
	return nil

}

func a() {
	panic(errors.New("this is a error"))
}

func main() {
	e := Unmarshal()
	if e != nil {
		fmt.Printf("%+v\n", e)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("end")
	//output
	//this is a error
	//end
}
