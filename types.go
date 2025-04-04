package main

import "math/rand/v2"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.IntN(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.IntN(1000000)),
	}
}
