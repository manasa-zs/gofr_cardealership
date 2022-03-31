package stores

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"fmt"
	"github.com/google/uuid"
	"training/customers/db"
	"training/customers/models"
)

const (
	create = "INSERT INTO customers(id,name,age) values (?,?,?)"
	get    = "select * from customers where id=?"
	update = "update customers set name=?,age=? where id=?"
	del    = "delete from customers where id=?"
)

type customer struct {
	db []models.Customer
}

func New() *customer {
	return &customer{db.Customer}
}

func (c *customer) Create(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error) {
	customer.ID = uuid.New()
	_, err := ctx.DB().ExecContext(ctx, create, customer.ID, customer.Name, customer.Age)
	if err != nil {
		return &models.Customer{}, errors.DB{Err: err}
	}

	return customer, nil
}

func (c *customer) Get(ctx *gofr.Context, ID uuid.UUID) (*models.Customer, error) {
	var customer = &models.Customer{}
	rows := ctx.DB().QueryRowContext(ctx, "SELECT * FROM customers where id=?", ID)
	err := rows.Scan(&customer.ID, &customer.Name, &customer.Age)

	if err != nil {
		return &models.Customer{}, errors.DB{Err: err}
	}
	return customer, nil
}

func (c *customer) Update(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error) {
	fmt.Println(customer.ID)
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET name=?,age=? where id=?",
		customer.Name, customer.Age, customer.ID)
	if err != nil {
		return &models.Customer{}, errors.DB{Err: err}
	}

	return customer, nil
}

func (c *customer) Delete(ctx *gofr.Context, ID uuid.UUID) error {
	_, err := ctx.DB().ExecContext(ctx, del, ID)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
