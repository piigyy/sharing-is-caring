package repository

import (
	"context"

	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	"github.com/piigyy/sharing-is-caring/pkg/logger"
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
	const caller = "repository.payment.SavePayment"
	logger.Info(ctx, caller, "saving payment %s", payment.TransactionDetails.OrderID)

	_, err = r.collection.InsertOne(ctx, &payment)
	if err != nil {
		logger.Error(ctx, caller, "SavePayment.r.collection.InsertOne return an error: %v", err)
		return
	}

	return
}
