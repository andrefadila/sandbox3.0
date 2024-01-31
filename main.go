package main

import (
	"fmt"
	"os"

	// "sandbox3.0/task"
	"github.com/joho/godotenv"
)

func init() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// task.Level1No2()
	// task.Level1No3()

	// nip, err := task.CreateNIP(task.Akhwat, 2024, 6, 10)
	// fmt.Println(nip, err)

	// nips, err := task.GenerateNIPs(task.Akhwat, 2024, 6, 10, 1)
	// fmt.Println(nips, err)

	// nextNip, err := task.CreateNextNIP("ARN242-00000")
	// fmt.Println(nextNip, err)

	// nextNips, err := task.GenerateNextNIPs("ART242-00200", 10)
	// fmt.Println(nextNips, err)

	fmt.Println(os.Getenv("DB_HOST"))
}
