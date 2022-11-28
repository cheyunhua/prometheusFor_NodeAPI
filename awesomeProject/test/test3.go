package main

import (
	"awesomeProject/logger"
	"fmt"
)

func YiDisk(str string) error {
	if str == "" {
		return fmt.Errorf("this is %+v a erros", str)
	} else {
		return nil
	}

}

func UU() {
	fmt.Println("testx")
}

func main() {
	if err := YiDisk(""); err != nil {
		logger.DefaultLogger.Errorf(" %v", err)
	}

	UU()

}
