package update

// UpdateTransactionRequest representa el request para actualizar una transacción
type UpdateTransactionRequest struct {
	TypeID      string  `json:"type_id,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Payed       bool    `json:"payed,omitempty"`
	ExpiryDate  string  `json:"expiry_date,omitempty"`
	Category    string  `json:"category,omitempty"`
}

// UpdateTransactionResponse representa la respuesta al actualizar una transacción
type UpdateTransactionResponse struct {
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

// UpdateTransactionResponseBuilder patrón builder para la respuesta
type UpdateTransactionResponseBuilder struct {
	response UpdateTransactionResponse
}

func NewUpdateTransactionResponseBuilder() *UpdateTransactionResponseBuilder {
	return &UpdateTransactionResponseBuilder{}
}

func (b *UpdateTransactionResponseBuilder) SetID(id string) *UpdateTransactionResponseBuilder {
	b.response.ID = id
	return b
}

func (b *UpdateTransactionResponseBuilder) SetTypeID(typeID string) *UpdateTransactionResponseBuilder {
	b.response.TypeID = typeID
	return b
}

func (b *UpdateTransactionResponseBuilder) SetDescription(description string) *UpdateTransactionResponseBuilder {
	b.response.Description = description
	return b
}

func (b *UpdateTransactionResponseBuilder) SetAmount(amount float64) *UpdateTransactionResponseBuilder {
	b.response.Amount = amount
	return b
}

func (b *UpdateTransactionResponseBuilder) SetPayed(payed bool) *UpdateTransactionResponseBuilder {
	b.response.Payed = payed
	return b
}

func (b *UpdateTransactionResponseBuilder) SetExpiryDate(expiryDate string) *UpdateTransactionResponseBuilder {
	b.response.ExpiryDate = expiryDate
	return b
}

func (b *UpdateTransactionResponseBuilder) SetCategory(category string) *UpdateTransactionResponseBuilder {
	b.response.Category = category
	return b
}

func (b *UpdateTransactionResponseBuilder) SetCreatedAt(createdAt string) *UpdateTransactionResponseBuilder {
	b.response.CreatedAt = createdAt
	return b
}

func (b *UpdateTransactionResponseBuilder) SetUpdatedAt(updatedAt string) *UpdateTransactionResponseBuilder {
	b.response.UpdatedAt = updatedAt
	return b
}

func (b *UpdateTransactionResponseBuilder) Build() *UpdateTransactionResponse {
	return &b.response
}
