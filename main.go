package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil // Интерфейс под капотом это структура из 2 полей
	return err                  // type = *os.PathError  data = nil
}

func main() {
	err := Foo()
	fmt.Println(err)        // nil     т.к data == nil
	fmt.Println(err == nil) // false
	// Мы сравниваем тип и data, т.к только data у нас == nil, а type != nil, выведет false
}
