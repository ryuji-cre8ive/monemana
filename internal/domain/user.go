package domain

import ()

type User struct {
	ID           string
	Name         string
	Transactions *[]Transaction
}
