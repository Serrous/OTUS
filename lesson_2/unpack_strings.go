// Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
// * "a4bc2d5e" => "aaaabccddddde"
// * "abcd" => "abcd"
// * "45" => "" (некорректная строка)

package main

import (
	"fmt"
	"strconv"
)

func main() {

	anyWord := "AZAp9ef4ЯЗЪ9"
	var tempElem rune
	var newString []rune

	for _, val := range anyWord {
		if counter, err := strconv.Atoi(string(val)); err != nil { // мне совсем не нравится данный способ - определять, является ли элемент числом, но я не смог придумать как иначе.
			tempElem = val
			newString = append(newString, val)
			continue
		} else {
			for i := 1; i < counter; i++ {
				newString = append(newString, tempElem)
			}
		}
	}
	if tempElem == 0 {
		fmt.Println("Некорректная строка")
	} else {
		fmt.Printf("%v\n", string(newString))
	}
}
