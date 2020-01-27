package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	var inpt = flag.String("envdir", "", "Path to envarement directory")

	flag.Parse()

	if len(*inpt) == 0 {
		fmt.Println("Флаг не задан")
		return
	}

	if err := os.Chdir(*inpt); err != nil {
		fmt.Println("Нет такой директории", *inpt)
		return
	}

	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Ошибка чтения содержимого")
		return
	}

	os.Clearenv() //очищаем переменные

	for _, i := range dirs {
		if !i.IsDir() {
			eval, err := ioutil.ReadFile(i.Name())
			if err != nil {
				fmt.Println("Невозможно прочитать содержимое файла", i.Name())
				return
			}
			os.Setenv(strings.TrimSpace(i.Name()), strings.TrimSpace(string(eval)))
		}

	}

	fmt.Println("Какая-то программа запустится с перемеными:")
	for _, r := range os.Environ() {
		fmt.Println(r)
	}

}
