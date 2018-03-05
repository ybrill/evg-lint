package testdata

import (
	"fmt"
	"io"
	"io/ioutil"
)

func tryDefer() {
	for i := 0; i < 10; i += 1 {
		defer func() error { // MATCH /for loop containing defer will not run until end of function/
			fmt.Printf("defer %d", i)
			return nil
		}()

		notADeferStmt := "test"
		fmt.Println(notADeferStmt)
	}
}

func tryDefer2() {
	data := []int{0, 1, 2, 3, 4, 5}
	for _ := range data {
		tempFile, err := ioutil.TempFile("", "mcipatch_")
		if err != nil {
			return err
		}
		defer tempFile.Close() // MATCH /for loop containing defer will not run until end of functio/
		_, err = io.WriteString(tempFile, "hi")
	}
}

func okDefer() {
	for i := 0; i < 10; i += 1 {
		doesTheDefer() // OK
	}
}

func doesTheDefer() {
	defer func() error { // OK
		fmt.Printf("defer %d", i)
		return nil
	}()

	notADeferStmt := "test"
	fmt.Println(notADeferStmt)
}
