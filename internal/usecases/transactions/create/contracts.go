package create

// CreateTransactionRequest representa el request para crear una transacción
type CreateTransactionRequest struct {
	UserID      string  `json:"user_id"`
	TypeID      string  `json:"type_id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Payed       bool    `json:"payed,omitempty"`
	ExpiryDate  string  `json:"expiry_date,omitempty"`
	CategoryID  string  `json:"category,omitempty"`
}

// CreateTransactionResponse representa la respuesta al crear una transacción
type CreateTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	CreatedAt     string `json:"created_at"`
}

// CreateTransactionResponseBuilder patrón builder para la respuesta
type CreateTransactionResponseBuilder struct {
	response CreateTransactionResponse
}

func NewCreateTransactionResponseBuilder() *CreateTransactionResponseBuilder {
	return &CreateTransactionResponseBuilder{}
}

func (b *CreateTransactionResponseBuilder) SetTransactionID(id string) *CreateTransactionResponseBuilder {
	b.response.TransactionID = id
	return b
}

func (b *CreateTransactionResponseBuilder) SetCreatedAt(createdAt string) *CreateTransactionResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *CreateTransactionResponseBuilder) Build() *CreateTransactionResponse {
	return &b.response
}

func (r *CreateTransactionResponse) ToContract() *CreateTransactionResponse {
	return r
}
