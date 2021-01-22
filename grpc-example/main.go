package main

import (
	"fmt"
	"grpc-example/model"
)

func main() {
	jufi := &model.User{
		Id:       "123",
		Name:     "Jufianto",
		Password: "jufiisteh",
		Gender:   model.UserGender_MALE,
	}

	fmt.Println(jufi)
}
