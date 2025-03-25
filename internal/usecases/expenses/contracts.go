package expenses

// CreateExpenseRequest representa el request para crear un gasto
type CreateExpenseRequest struct {
	UserID      string  `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category,omitempty"`
	Paid        bool    `json:"paid"`
	DueDate     string  `json:"due_date,omitempty"`
}

// CreateExpenseResponse representa la respuesta al crear un gasto
type CreateExpenseResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category,omitempty"`
	Paid        bool    `json:"paid"`
	DueDate     string  `json:"due_date,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// GetExpenseResponse representa la respuesta al obtener un gasto
type GetExpenseResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category,omitempty"`
	Paid        bool    `json:"paid"`
	DueDate     string  `json:"due_date,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// ListExpensesResponse representa la respuesta al listar gastos
type ListExpensesResponse struct {
	Expenses []GetExpenseResponse `json:"expenses"`
}

// UpdateExpenseRequest representa el request para actualizar un gasto
type UpdateExpenseRequest struct {
	Amount      float64 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
	Category    string  `json:"category,omitempty"`
	Paid        bool    `json:"paid,omitempty"`
	DueDate     string  `json:"due_date,omitempty"`
}

// UpdateExpenseResponse representa la respuesta al actualizar un gasto
type UpdateExpenseResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category,omitempty"`
	Paid        bool    `json:"paid"`
	DueDate     string  `json:"due_date,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// MarkAsPaidResponse representa la respuesta al marcar un gasto como pagado
type MarkAsPaidResponse struct {
	ID        string `json:"id"`
	Paid      bool   `json:"paid"`
	UpdatedAt string `json:"updated_at"`
}
