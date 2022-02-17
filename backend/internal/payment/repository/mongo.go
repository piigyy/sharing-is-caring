package repository

import (
	"context"
	"log"

	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type payment struct {
	collection *mongo.Collection
}

func NewPaymentRepository(
	mongoDB *mongo.Database,
) *payment {
	return &payment{
		collection: mongoDB.Collection("payment"),
	}
}

func (r *payment) SavePayment(ctx context.Context, payment model.PaymentRequest) (err error) {
	log.Printf("saving payment %s", payment.TransactionDetails.OrderID)
	_, err = r.collection.InsertOne(ctx, &payment)
	if err != nil {
		log.Printf("error SavePayment.r.collection.InsertOne: %v", err)
		return
	}

	return
}
