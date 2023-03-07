package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Во-первых, глобальная переменная не лучший выбор здесь, так как их можно изменить из любой функции,
//а работа с ними медленнее чем с локальными
// Во-вторых, лучше такую огромную строку записывать в файл и с него читать сколько нужно
// И так как строки в go иммутабельные, при изменении лучше использовать []byte

func createHugeString(str int, r io.Writer) (string, error) {
	x := "1"
	fmt.Println(str)
	for i := 0; i < str; i++ {
		x += x
	}
	_, err := r.Write([]byte(x))
	if err != nil {
		return x, err
	}
	return x[:100], nil
}

func someFunc() string {
	f, err := os.Create("huge_string.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	justString, err := createHugeString(1<<10, f)
	if err != nil {
		log.Fatal(err)
	}
	return justString
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
