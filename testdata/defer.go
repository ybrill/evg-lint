package testdata

import "fmt"

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
