.PHONY: proto

proto:
	protoc proto/payment.proto --go_out=plugins=grpc:.

run-auth:
	go run cmd/auth/http/main.go

run-payment:
	go run cmd/payment/grpc/main.go

