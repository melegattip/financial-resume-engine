package transactions

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	Date        string  `json:"date"`
}
