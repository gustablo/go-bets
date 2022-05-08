package wallet

func Register(Amount int) Wallet {
	wallet := Wallet{}

	wallet.Amount = Amount

	return wallet
}
