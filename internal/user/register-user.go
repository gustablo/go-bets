package user

import (
	"github.com/gustablo/go-bets/internal/wallet"
)

func Register(user *User) User {
	userInstance := User{}
	userInstance.FullName = user.FullName
	userInstance.Wallet = wallet.Register(0)

	return userInstance
}
