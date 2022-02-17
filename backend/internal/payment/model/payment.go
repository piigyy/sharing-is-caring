package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/mail"

	pb "github.com/piigyy/sharing-is-caring/internal/payment/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	OrderID     string  `json:"order_id" bson:"orderID"`
	GrossAmount float32 `json:"gross_amount" bson:"grossAmount"`
}

func (t *Transaction) Validate() error {
	if t.GrossAmount < 0.01 {
		return ErrInvalidTransactionAmount
	}
	return nil
}

type Item struct {
	ID       string  `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Price    float32 `json:"price" bson:"price"`
	Quantity int     `json:"quantity" bson:"quantity"`
}

type Customer struct {
	ID        string `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"firstName"`
	LastName  string `json:"last_name" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
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
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Status             string             `bson:"status"`
	PaymentType        string             `json:"payment_type" bson:"paymentType"`
	TransactionDetails Transaction        `json:"transaction_details" bson:"transactionDetails"`
	ItemDetails        []Item             `json:"item_details" bson:"itemDetails"`
	CustomerDetails    Customer           `json:"customer_details" bson:"customerDetails"`
}

func (pr *PaymentRequest) UpdateOrderID(orderID string) {
	pr.TransactionDetails.OrderID = orderID
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

func MapPaymentRequestProto(request *pb.PaymentRequest, orderID string) PaymentRequest {
	return PaymentRequest{
		Status:      PAYMENT_PENDING,
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
