package user

import (
	"github.com/gustablo/go-bets/internal/wallet"
)

type User struct {
	FullName string
	Wallet   wallet.Wallet
}
