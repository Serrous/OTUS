package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type fArgs struct {
	sourceFile, destFile string
	offset, limit        int
}

func main() {
	inputArgs := os.Args
	fmt.Scan(inputArgs)

	//pntrArgs := &inputArgs

	// limit, err := strconv.Atoi(inputArgs[4])
	// if err != nil { //проверяем, что offset представлен числом
	// 	fmt.Println(errors.New("Некорректное значение для limit"))
	// 	return
	// }

	//finalArgs := &fArgs{inputArgs[1], inputArgs[1], offset, limit}

	//fmt.Println(finalArgs.sourceFile, finalArgs.offset, finalconv
	// 	count := 100000
	// 	// create and start new bar
	// 	bar := pb.StartNew(count)

	// 	for i := 0; i < count; i++ {
	// 		bar.Increment()
	// 		time.Sleep(time.Millisecond)
	// 	}
	// 	bar.Finish()

	// fi, err := os.Stat(finalArgs.sourceFile)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fi.Size())
	// }
	asdf := createArgsStruct(inputArgs)
	copy(asdf)
}

func checkArgs(inputVal []string) error { //проверяем количество аргументов

	if len(inputVal) > 5 || len(inputVal) < 3 {
		return errors.New("Недопустимое число параметров. Число входных параметров: 2, 3, 4")
	}

	// for _, i := range inputVal[3:] {
	// 	if err := convOffsetLimit(i); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func createArgsStruct(inputVal []string) fArgs {
	var finalStruct fArgs
	if err := checkArgs(inputVal); err == nil {

		switch len(inputVal) {
		case 3:
			finalStruct = fArgs{inputVal[1], inputVal[2], 0, 0}
		case 4:
			if offs, err := strconv.Atoi(inputVal[3]); err == nil {
				finalStruct = fArgs{inputVal[1], inputVal[2], offs, 0}
			}
		}

		finalStruct = fArgs{inputVal[1], inputVal[2], 0, 2048}
	}
	return finalStruct
}

func copy(asdf fArgs) {
	fmt.Println(asdf)

	lim := asdf.limit
	buf := make([]byte, lim)

	file, err := os.Open(asdf.sourceFile)
	if err != nil {
		fmt.Println(err)
	}

	ofst := asdf.offset

	for ofst < lim {
		read, err := file.Read(buf[ofst:])
		ofst += read
		if err == io.EOF {
			break
		}
		file2, _ := os.Create(asdf.destFile)
		writen, err := file.Write(read)

		if err != nil {
			fmt.Println(err)
		}

	}

}
