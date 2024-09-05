package libs

var balance = 0

func GetBalance() int {
	return balance
}

func SetBalance(amount int) int {
	balance = amount
	return balance
}
