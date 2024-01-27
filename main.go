package main

import (
	"fmt"

	"sandbox3.0/task"
)

func main() {
	// task.Level1No2()
	// task.Level1No3()

	// nip, err := task.CreateNIP(task.Akhwat, 2024, 6, 10)
	// fmt.Println(nip, err)

	// nips, err := task.GenerateNIPs(task.Akhwat, 2024, 6, 10, 1)
	// fmt.Println(nips, err)

	nextNip, err := task.CreateNextNIP("ARN242-00000")
	fmt.Println(nextNip, err)

	// nextNips, err := task.GenerateNextNIPs("ART242-00200", 10)
	// fmt.Println(nextNips, err)

}
