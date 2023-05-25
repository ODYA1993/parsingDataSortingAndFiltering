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

	user2 := models.User{
		ID:   2,
		Name: "Сергей",
		Age:  33,
	}

	fmt.Println(user)
	fmt.Println(user2)
}
