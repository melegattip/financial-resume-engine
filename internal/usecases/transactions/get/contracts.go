package get

// GetTransactionResponse representa la respuesta al obtener una transacción
type GetTransactionResponse struct {
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

// GetTransactionResponseBuilder patrón builder para la respuesta
type GetTransactionResponseBuilder struct {
	response GetTransactionResponse
}

func NewGetTransactionResponseBuilder() *GetTransactionResponseBuilder {
	return &GetTransactionResponseBuilder{}
}

func (b *GetTransactionResponseBuilder) SetID(id string) *GetTransactionResponseBuilder {
	b.response.ID = id
	return b
}

func (b *GetTransactionResponseBuilder) SetTypeID(typeID string) *GetTransactionResponseBuilder {
	b.response.TypeID = typeID
	return b
}

func (b *GetTransactionResponseBuilder) SetDescription(description string) *GetTransactionResponseBuilder {
	b.response.Description = description
	return b
}

func (b *GetTransactionResponseBuilder) SetAmount(amount float64) *GetTransactionResponseBuilder {
	b.response.Amount = amount
	return b
}

func (b *GetTransactionResponseBuilder) SetPayed(payed bool) *GetTransactionResponseBuilder {
	b.response.Payed = payed
	return b
}

func (b *GetTransactionResponseBuilder) SetExpiryDate(expiryDate string) *GetTransactionResponseBuilder {
	b.response.ExpiryDate = expiryDate
	return b
}

func (b *GetTransactionResponseBuilder) SetCategory(category string) *GetTransactionResponseBuilder {
	b.response.Category = category
	return b
}

func (b *GetTransactionResponseBuilder) SetCreatedAt(createdAt string) *GetTransactionResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *GetTransactionResponseBuilder) SetUpdatedAt(updatedAt string) *GetTransactionResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *GetTransactionResponseBuilder) Build() *GetTransactionResponse {
	return &b.response
}
