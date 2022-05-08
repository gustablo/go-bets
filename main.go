package main

import (
	"fmt"

	"github.com/gustablo/go-bets/internal/transaction"
	"github.com/gustablo/go-bets/internal/user"
)

func main() {
	user1 := user.Register(&user.User{FullName: "gustavo"})
	user2 := user.Register(&user.User{FullName: "joão"})

	var err error

	debt := 10

	user1.Wallet, user2.Wallet, err = transaction.Transaction(user1.Wallet, user2.Wallet, debt)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user1, user2)
}
