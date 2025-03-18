package models

// TransactionRequest representa el request para crear una transacción
type TransactionRequest struct {
	TypeID      string  `json:"type_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Payed       bool    `json:"payed,omitempty"`
	ExpiryDate  string  `json:"expiry_date,omitempty"`
	Category    string  `json:"category,omitempty"`
}

// TransactionResponse representa la respuesta al crear una transacción
type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	CreatedAt     string `json:"created_at"`
}

// TransactionResponseBuilder patrón builder para la respuesta
type TransactionResponseBuilder struct {
	response TransactionResponse
}

func NewTransactionResponseBuilder() *TransactionResponseBuilder {
	return &TransactionResponseBuilder{}
}

func (b *TransactionResponseBuilder) SetTransactionID(id string) *TransactionResponseBuilder {
	b.response.TransactionID = id
	return b
}

func (b *TransactionResponseBuilder) SetCreatedAt(createdAt string) *TransactionResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *TransactionResponseBuilder) Build() *TransactionResponse {
	return &b.response
}
