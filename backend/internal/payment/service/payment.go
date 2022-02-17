package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	pb "github.com/piigyy/sharing-is-caring/internal/payment/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type payment struct {
	http       *http.Client
	paymentURL string
	paymentKey string
}

func NewPayment(paymentURL, paymentKey string) *payment {
	key := fmt.Sprintf("%s:", paymentKey)
	encodedKey := base64.StdEncoding.EncodeToString([]byte(key))
	return &payment{
		http:       &http.Client{},
		paymentURL: paymentURL,
		paymentKey: encodedKey,
	}
}

func (s *payment) CreatePayment(ctx context.Context, request *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	orderID, paymentRequest := model.MapPaymentRequestProto(request)
	log.Printf("new payment with order id: %sv", orderID)

	if err := paymentRequest.Validate(); err != nil {
		log.Printf("error paymentRequest.Validate(): %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	paymentResponse, err := s.requestToMidtrans(paymentRequest)
	if err != nil {
		log.Printf("failed sending request to midtrans: %v", err)
		return nil, status.Errorf(codes.Internal, "failed sending request to midtrans: %v", err)
	}

	return &pb.PaymentResponse{
		Success:     true,
		Token:       paymentResponse.Token,
		RedirectUrl: paymentResponse.RedirectURL,
		OrderId:     orderID,
	}, nil
}

func (s *payment) requestToMidtrans(paymentRequest model.PaymentRequest) (response *model.PaymentResponse, err error) {
	var paymentResponse model.PaymentResponse
	log.Println("requesting SNAP Payment To Midtrans...")

	paymentRequestReader, err := paymentRequest.PaymentToIOReader()
	if err != nil {
		log.Printf("failed convert model.PaymentRequest to io.Reader: %v", err)
		return nil, status.Errorf(codes.Internal, "failed convert model.PaymentRequest to io.Reader: %v", err)
	}

	req, reqErr := http.NewRequest(http.MethodPost, s.paymentURL, paymentRequestReader)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", s.paymentKey))
	req.Header.Set("Content-Type", "application/json")
	if reqErr != nil {
		log.Printf("failed creeating http.NewRequest: %v", reqErr)
		return nil, status.Errorf(codes.Internal, "failed creeating http.NewRequest: %v", reqErr)
	}

	resp, respErr := s.http.Do(req)
	if respErr != nil {
		log.Printf("failed request to midtrans: %v", respErr)
		return nil, status.Errorf(codes.Internal, "failed request do midtrans: %v", respErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
			log.Printf("failed decode http response to model.PaymentResponse: %v", err)
			return nil, status.Errorf(codes.Internal, "failed decode http response to model.PaymentResponse: %v", err)
		}

		log.Printf("SNAP Payment Request to Midtrans success, token: %s", paymentResponse.Token)
		return &paymentResponse, nil
	}

	log.Println("SNAP Payment Request to Midtrans failed")
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return nil, errors.New(bodyString)
}
