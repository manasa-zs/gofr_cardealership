package services

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"strings"
	"training/customers/models"
)

type service struct {
	customer CustomerStore
}

func New(c CustomerStore) *service {
	return &service{customer: c}
}

func (s *service) Create(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error) {
	validCustomer, err := Validate(customer)
	if err != nil {
		return &models.Customer{}, err
	}

	resp, err := s.customer.Create(ctx, validCustomer)
	if err != nil {
		return &models.Customer{}, err
	}

	return resp, nil
}

func (s *service) Get(ctx *gofr.Context, ID uuid.UUID) (*models.Customer, error) {
	customer, err := s.customer.Get(ctx, ID)
	if err != nil {
		return &models.Customer{}, err
	}

	return customer, nil
}

func (s *service) Update(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error) {
	validCustomer, err := Validate(customer)
	if err != nil {
		return &models.Customer{}, err
	}

	resp, err := s.customer.Update(ctx, validCustomer)
	if err != nil {
		return &models.Customer{}, err
	}

	return resp, nil

}

func (s *service) Delete(ctx *gofr.Context, ID uuid.UUID) error {
	err := s.customer.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func Validate(customer *models.Customer) (*models.Customer, error) {
	if customer.Age <= 18 && customer.Age >= 80 {
		return &models.Customer{}, errors.Error("age must be between 18 to 80")
	}

	name := strings.TrimSpace(customer.Name)
	if name == "" {
		return &models.Customer{}, errors.Error("name cannot be empty spaces")
	} else if len(name) < 6 {
		return &models.Customer{}, errors.Error("length of name must be atleast 6")
	}

	return customer, nil
}
