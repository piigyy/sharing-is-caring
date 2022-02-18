package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	pb "github.com/piigyy/sharing-is-caring/internal/payment/proto"
	"github.com/piigyy/sharing-is-caring/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type payment struct {
	http       *http.Client
	paymentURL string
	paymentKey string
	repository model.PaymentReaderWriterRepository
}

func NewPayment(
	paymentURL, paymentKey string,
	repository model.PaymentReaderWriterRepository,
) *payment {
	key := fmt.Sprintf("%s:", paymentKey)
	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))
	return &payment{
		http:       &http.Client{},
		paymentURL: paymentURL,
		paymentKey: encodedKey,
		repository: repository,
	}
}

func (s *payment) CreatePayment(ctx context.Context, request *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	const caller = "payment.CreatePayment"
	ctx = s.setContext(ctx)

	orderID := primitive.NewObjectID()
	paymentRequest := model.MapPaymentRequestProto(request, orderID.Hex())
	paymentRequest.ID = orderID
	logger.Info(ctx, caller, "new create payment request")

	if err := paymentRequest.Validate(); err != nil {
		logger.Error(ctx, caller, "paymentRequest.Validate(): %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.repository.SavePayment(ctx, paymentRequest)
	if err != nil {
		logger.Error(ctx, caller, "s.repository.SavePayment return an error: %v", err)
		return nil, status.Errorf(codes.Internal, "error s.repository.SavePayment: %v", err)
	}

	paymentResponse, err := s.requestToMidtrans(ctx, paymentRequest)
	if err != nil {
		logger.Error(ctx, caller, "s.requestToMidtrans return an error: %v", err)
		return nil, status.Errorf(codes.Internal, "failed sending request to midtrans: %v", err)
	}

	logger.Info(ctx, caller, "SNAP payment link created")
	return &pb.PaymentResponse{
		Success:     true,
		Token:       paymentResponse.Token,
		RedirectUrl: paymentResponse.RedirectURL,
		OrderId:     orderID.Hex(),
	}, nil
}

func (s *payment) requestToMidtrans(ctx context.Context, paymentRequest model.PaymentRequest) (response *model.PaymentResponse, err error) {
	const caller = "payment.CreatePayment"
	var paymentResponse model.PaymentResponse
	logger.Info(ctx, caller, "request creating SNAP Payment to midtrans")

	paymentRequestReader, err := paymentRequest.PaymentToIOReader()
	if err != nil {
		logger.Error(ctx, caller, "failed convert model.PaymentRequest to io.Reader: %v", err)
		return nil, status.Errorf(codes.Internal, "failed convert model.PaymentRequest to io.Reader: %v", err)
	}

	req, reqErr := http.NewRequest(http.MethodPost, s.paymentURL, paymentRequestReader)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", s.paymentKey))
	req.Header.Set("Content-Type", "application/json")
	if reqErr != nil {
		logger.Error(ctx, caller, "failed creeating http.NewRequest: %v", reqErr)
		return nil, status.Errorf(codes.Internal, "failed creeating http.NewRequest: %v", reqErr)
	}

	resp, respErr := s.http.Do(req)
	if respErr != nil {
		logger.Error(ctx, caller, "failed request to midtrans: %v", respErr)
		return nil, status.Errorf(codes.Internal, "failed request do midtrans: %v", respErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
			logger.Error(ctx, caller, "failed decode http response to model.PaymentResponse: %v", err)
			return nil, status.Errorf(codes.Internal, "failed decode http response to model.PaymentResponse: %v", err)
		}

		logger.Info(ctx, caller, "SNAP Payment Request to Midtrans success")
		return &paymentResponse, nil
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	logger.Error(ctx, caller, "SNAP Payment Request to Midtrans failed; error response: %s", bodyString)
	return nil, errors.New(bodyString)
}

func (s *payment) setContext(parent context.Context) context.Context {
	correlationID := uuid.New().String()
	return context.WithValue(parent, logger.RequestIDKey, correlationID)
}
