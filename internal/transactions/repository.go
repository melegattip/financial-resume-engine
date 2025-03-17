package transactions

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *Transaction) error {
	return r.db.QueryRow("INSERT INTO transactions (amount, description, category_id, date) VALUES ($1, $2, $3, $4) RETURNING id",
		transaction.Amount, transaction.Description, transaction.CategoryID, transaction.Date).Scan(&transaction.ID)
}

func (r *TransactionRepository) ListTransactions() ([]Transaction, error) {
	rows, err := r.db.Query("SELECT id, amount, description, category_id, date FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []Transaction{}
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Description, &transaction.CategoryID, &transaction.Date); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction *Transaction) error {
	return r.db.QueryRow("UPDATE transactions SET amount = $1, description = $2, category_id = $3, date = $4 WHERE id = $5",
		transaction.Amount, transaction.Description, transaction.CategoryID, transaction.Date, transaction.ID).Scan()
}
