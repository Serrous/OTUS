package main

import (
	"errors"
	"fmt"
	"runtime"
)

type myType = func() error

func aaa() error {

	sl := []int{1, 2, 3, 4, 5}
	for _, i := range sl {
		if i == 3 {
			return errors.New("Произошла ошибка в функции aaa")
		}
		fmt.Println("aaa", i)
	}
	return nil
}

func bbb() error {

	sl := []int{6, 7, 8, 9, 10}
	for _, i := range sl {
		if i == 8 {
			return errors.New("Произошла ошибка в функции bbb")
		}
		fmt.Println("bbb", i)
	}
	return nil
}

//Run is our function
func Run(task []myType, n int, m int) error {
	var countGorutines = make(chan bool, n) //канал для ограничения конкрурентного запуска горутин
	var errs = make(chan error)             //канал для получения ошибок из горутин
	var quit = make(chan bool)              //канал для сигнала завершения
	errCount := 0                           //счетчик ошибок

	defer fmt.Println("Сейчас запущено горутин:", runtime.NumGoroutine()) // для проверки оставшихся горутин

	defer func() { //обрабатываем панику, которая будет вызвана попыткой записать в закрытый канал.
		if err := recover(); err != nil {
			fmt.Println("panic!", err)
		}
	}()

	go func() { //запускаем WatchDog, который остлеживает прилетающие ошибки из канала errs
		for {
			msg := <-errs
			fmt.Println(msg)
			errCount++

			if errCount >= m {
				fmt.Println("Произошло", errCount, "ошибок, посылаю сигнал завершения")
				close(quit)           //посылаем сигнал завершения путем закрытия канала quit
				close(countGorutines) //закрываем канал
				return
			}

		}
	}()

	for _, i := range task {
		countGorutines <- true
		go func(n myType) {
			for {
				select {
				case <-quit:
					fmt.Println("Получен сигнал завершения")
					return
				default:
					result := n()
					if result != nil {
						errs <- result
					}
					<-countGorutines
					return
				}
			}
		}(i)

	}

	return nil
}

func main() {

	myTasks := []myType{aaa, bbb, aaa, bbb, aaa, bbb, aaa, bbb, aaa, bbb}

	fmt.Println(Run(myTasks, 2, 4))
}
