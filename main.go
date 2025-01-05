package main

import "log"

// my interview question program
type C struct {
	Name    string
	Balance int
}

type B struct {
	Name    string
	Balance int
}

func main() {
	isValidBalance := BoolFunc(C{Name: "Bank", Balance: 10})
	log.Println(isValidBalance)

	AccountOperations(&InstagramAccount{})
}

func BoolFunc(data interface{}) bool {
	val, _ := data.(B)
	return val.Balance == 10

}

type MethodsList interface {
	CreateAccountID()
	GetAccountID() int
}

type FacebookAccount struct {
	ID int
}

type InstagramAccount struct {
	ID int
}

func (f *FacebookAccount) CreateAccountID() {
	f.ID = 5
}

func (i InstagramAccount) GetAccountID() int {
	return i.ID
}

func AccountOperations(m MethodsList) {
	m.CreateAccountID()
	m.GetAccountID()
}
