package model

import "context"

type (
	PaymentReaderRepository interface{}

	PaymentWriterRepository interface {
		SavePayment(ctx context.Context, payment PaymentRequest) (err error)
	}

	PaymentReaderWriterRepository interface {
		PaymentReaderRepository
		PaymentWriterRepository
	}
)
