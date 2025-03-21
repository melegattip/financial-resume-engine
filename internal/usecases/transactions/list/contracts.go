package list

// ListTransactionsResponse representa la respuesta al listar transacciones
type ListTransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

// TransactionResponse representa una transacción en la respuesta
type TransactionResponse struct {
	ID          string  `json:"id"`
	TypeID      string  `json:"type_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Payed       bool    `json:"payed"`
	ExpiryDate  string  `json:"expiry_date,omitempty"`
	Category    string  `json:"category,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// ListTransactionsResponseBuilder patrón builder para la respuesta
type ListTransactionsResponseBuilder struct {
	response ListTransactionsResponse
}

func NewListTransactionsResponseBuilder() *ListTransactionsResponseBuilder {
	return &ListTransactionsResponseBuilder{}
}

func (b *ListTransactionsResponseBuilder) SetTransactions(transactions []TransactionResponse) *ListTransactionsResponseBuilder {
	b.response.Transactions = transactions
	return b
}

func (b *ListTransactionsResponseBuilder) Build() *ListTransactionsResponse {
	return &b.response
}
