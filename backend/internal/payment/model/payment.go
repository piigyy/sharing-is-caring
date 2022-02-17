package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/mail"

	"github.com/google/uuid"
	pb "github.com/piigyy/sharing-is-caring/internal/payment/proto"
)

type Validator interface {
	Validate() (err error)
}

func Validate(e Validator) error {
	return e.Validate()
}

var (
	ErrInvalidTransactionAmount = errors.New("invalid gross_amount; should be greater than 0.01 idr")
	ErrInvalidCustomerEmail     = errors.New("invalid customer email")
	ErrInvalidCustomerPhone     = errors.New("invalid customer phone")
)

type PaymentResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type Transaction struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float32 `json:"gross_amount"`
}

func (t *Transaction) Validate() error {
	if t.GrossAmount < 0.01 {
		return ErrInvalidTransactionAmount
	}
	return nil
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (c *Customer) Validate() error {
	_, err := mail.ParseAddress(c.Email)
	if err != nil {
		return ErrInvalidCustomerEmail
	}

	if len(c.Phone) < 8 {
		return ErrInvalidCustomerPhone
	}

	return nil
}

type PaymentRequest struct {
	PaymentType        string      `json:"payment_type"`
	TransactionDetails Transaction `json:"transaction_details"`
	ItemDetails        []Item      `json:"item_details"`
	CustomerDetails    Customer    `json:"customer_details"`
}

func (pr *PaymentRequest) Validate() error {
	if err := pr.TransactionDetails.Validate(); err != nil {
		return err
	}

	if err := pr.CustomerDetails.Validate(); err != nil {
		return err
	}

	return nil
}

func (pr *PaymentRequest) PaymentToIOReader() (io.Reader, error) {
	var buff bytes.Buffer
	if err := json.NewEncoder(&buff).Encode(pr); err != nil {
		return nil, err
	}

	return &buff, nil
}

func MapPaymentRequestProto(request *pb.PaymentRequest) (string, PaymentRequest) {
	orderID := uuid.New().String()
	return orderID, PaymentRequest{
		PaymentType: "qris",
		TransactionDetails: Transaction{
			OrderID:     orderID,
			GrossAmount: request.GetGrossAmount(),
		},
		ItemDetails: []Item{
			{
				ID:       request.GetItem().GetId(),
				Name:     request.GetItem().GetName(),
				Price:    request.GetItem().GetPrice(),
				Quantity: int(request.GetItem().GetQuantity()),
			},
		},
		CustomerDetails: Customer{
			ID:        request.GetCustomer().GetId(),
			FirstName: request.GetCustomer().GetFirstName(),
			LastName:  request.GetCustomer().GetLastName(),
			Email:     request.GetCustomer().GetEmail(),
			Phone:     request.GetCustomer().GetPhone(),
		},
	}
}
