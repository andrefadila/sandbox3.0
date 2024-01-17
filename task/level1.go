package task

import (
	"errors"
	"fmt"
	"time"
)

func Level1No2() {
	
	sl1 := fmt.Sprintln("Hello,", "Sprintln!")
	fmt.Print(sl1)
	
	fmt.Println("Hello, Println!")

	sl2 := fmt.Sprintln("Hello", "again", "Sprintln!")
	fmt.Print(sl2)
}

func Level1No3() {

    errf := fmt.Errorf("error occurred at: %v", time.Now())
    fmt.Println(errf)

	err := errors.New("error occurred (text only)") 
    fmt.Println(err)
}