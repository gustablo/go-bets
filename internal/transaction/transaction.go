package transaction

import (
	"errors"

	"github.com/gustablo/go-bets/internal/wallet"
)

func Transaction(receiver wallet.Wallet, debtor wallet.Wallet, amount int) (wallet.Wallet, wallet.Wallet, error) {
	if receiver.Amount-amount < 0 {
		return receiver, debtor, errors.New("Transaction not authorized because debt would be less than 0")
	}

	receiver.Amount -= amount
	debtor.Amount += amount

	return receiver, debtor, nil
}
