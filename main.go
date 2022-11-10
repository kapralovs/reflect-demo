package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Age    int
	Weight float64
	Items  []string
	Keys   map[string]string
}

func main() {
	var f float64 = 3.4           // объявляем переменную f с типом float64 и значением 3.4
	p := reflect.ValueOf(&f)      // получем значение по указателю
	v := p.Elem()                 // с помощью метода Elem() получаем значение, которое лежит внутри интерфейса или на которое ссылается указатель
	fmt.Printf("Before: %v\n", v) // выводим значение переменной f на экран
	v.SetFloat(7.1)               // меняем значение f на 7.1
	fmt.Printf("After: %v\n", v)  // выводим значение переменной f на экран

}
