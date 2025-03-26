package incomes

import (
	"context"
	"time"

	"github.com/melegattip/financial-resume-engine/internal/core/domain"
)

// IncomeService define las operaciones disponibles para el servicio de ingresos
type IncomeService interface {
	CreateIncome(ctx context.Context, request *CreateIncomeRequest) (*CreateIncomeResponse, error)
	GetIncome(ctx context.Context, userID string, incomeID string) (*GetIncomeResponse, error)
	ListIncomes(ctx context.Context, userID string) (*ListIncomesResponse, error)
	UpdateIncome(ctx context.Context, userID string, incomeID string, request *UpdateIncomeRequest) (*UpdateIncomeResponse, error)
	DeleteIncome(ctx context.Context, userID string, incomeID string) error
}

// IncomeServiceImpl implementa IncomeService
type IncomeServiceImpl struct {
	repository IncomeRepository
}

// IncomeRepository define las operaciones necesarias para el repositorio de ingresos
type IncomeRepository interface {
	Create(income *domain.Income) error
	Get(userID string, id string) (*domain.Income, error)
	List(userID string) ([]*domain.Income, error)
	Update(income *domain.Income) error
	Delete(userID string, id string) error
}

func NewIncomeService(repository IncomeRepository) IncomeService {
	return &IncomeServiceImpl{
		repository: repository,
	}
}

func (s *IncomeServiceImpl) CreateIncome(ctx context.Context, request *CreateIncomeRequest) (*CreateIncomeResponse, error) {
	income := domain.NewIncomeBuilder().
		SetID(domain.NewID()).
		SetUserID(request.UserID).
		SetAmount(request.Amount).
		SetDescription(request.Description).
		SetCategory(request.Category).
		SetSource(request.Source).
		Build()

	if err := s.repository.Create(income); err != nil {
		return nil, err
	}

	return &CreateIncomeResponse{
		ID:          income.ID,
		UserID:      income.UserID,
		Amount:      income.Amount,
		Description: income.Description,
		Category:    income.Category,
		Source:      income.Source,
		CreatedAt:   income.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   income.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *IncomeServiceImpl) GetIncome(ctx context.Context, userID string, incomeID string) (*GetIncomeResponse, error) {
	income, err := s.repository.Get(userID, incomeID)
	if err != nil {
		return nil, err
	}

	return &GetIncomeResponse{
		ID:          income.ID,
		UserID:      income.UserID,
		Amount:      income.Amount,
		Description: income.Description,
		Category:    income.Category,
		Source:      income.Source,
		CreatedAt:   income.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   income.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *IncomeServiceImpl) ListIncomes(ctx context.Context, userID string) (*ListIncomesResponse, error) {
	incomes, err := s.repository.List(userID)
	if err != nil {
		return nil, err
	}

	var response []GetIncomeResponse
	for _, income := range incomes {
		response = append(response, GetIncomeResponse{
			ID:          income.ID,
			UserID:      income.UserID,
			Amount:      income.Amount,
			Description: income.Description,
			Category:    income.Category,
			Source:      income.Source,
			CreatedAt:   income.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   income.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &ListIncomesResponse{
		Incomes: response,
	}, nil
}

func (s *IncomeServiceImpl) UpdateIncome(ctx context.Context, userID string, incomeID string, request *UpdateIncomeRequest) (*UpdateIncomeResponse, error) {
	income, err := s.repository.Get(userID, incomeID)
	if err != nil {
		return nil, err
	}

	if request.Amount != 0 {
		income.Amount = request.Amount
	}
	if request.Description != "" {
		income.Description = request.Description
	}
	if request.Category != "" {
		income.Category = request.Category
	}
	if request.Source != "" {
		income.Source = request.Source
	}
	income.UpdatedAt = time.Now()

	if err := s.repository.Update(income); err != nil {
		return nil, err
	}

	return &UpdateIncomeResponse{
		ID:          income.ID,
		UserID:      income.UserID,
		Amount:      income.Amount,
		Description: income.Description,
		Category:    income.Category,
		Source:      income.Source,
		CreatedAt:   income.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   income.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *IncomeServiceImpl) DeleteIncome(ctx context.Context, userID string, incomeID string) error {
	_, err := s.repository.Get(userID, incomeID)
	if err != nil {
		return err
	}

	return s.repository.Delete(userID, incomeID)
}
