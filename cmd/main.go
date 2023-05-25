package main

import (
	"fmt"
	"github.com/DmitryOdintsov/workingWithGit/internal/models"
)

func main() {
	user := models.User{
		ID:   1,
		Name: "Анатолий",
		Age:  30,
	}
	fmt.Println(user)
}
