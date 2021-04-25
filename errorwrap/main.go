package main

import (
	"errorWrap/dao"
	"fmt"
	"log"
)

func main() {
	// init db client
	client, err := dao.NewAutoDao()
	if err != nil {
		panic(err)
	}

	user := client.QueryUser()

	if user.Err != nil {
		// handler err
		log.Print(user.Err)
		return
	}
	fmt.Println(user)
}
