package main

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

var transactions []Transaction

func AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return errors.New("amount cannot be zero")
	}

	tx.ID = len(transactions) + 1

	transactions = append(transactions, tx)

	return nil
}

func ListTransactions() []Transaction {
	return transactions
}

func main() {
	fmt.Println("Ledger service started")

	tx1 := Transaction{
		Amount:      150.50,
		Category:    "Groceries",
		Description: "Weekly shopping",
		Date:        time.Now(),
	}

	err := AddTransaction(tx1)
	if err != nil {
		fmt.Println("Error adding transaction 1:", err)
	}

	tx2 := Transaction{
		Amount:      -25.0,
		Category:    "Transport",
		Description: "Metro card top-up",
		Date:        time.Now(),
	}

	err = AddTransaction(tx2)
	if err != nil {
		fmt.Println("Error adding transaction 2:", err)
	}

	// Попытка добавить невалидную транзакцию
	txInvalid := Transaction{Amount: 0}
	err = AddTransaction(txInvalid)
	if err != nil {
		fmt.Println("Correctly caught an error:", err)
	}

	allTxs := ListTransactions()
	fmt.Println("\n--- All Transactions ---")
	for _, tx := range allTxs {
		fmt.Printf("%+v\n", tx)
	}
}
