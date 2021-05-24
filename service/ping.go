package service

import (
	"gabrielrsilva/go-ms-scketch/object"
	"net/http"
)

type PingService interface {
	GetPing() object.ResponseTemplate
}

type pingService struct {
}

func GetPingService() PingService {
	return &pingService{}
}

func (svc *pingService) GetPing() object.ResponseTemplate {
	return object.NewResponseTemplate(http.StatusOK, &object.PingResponse{Message: "pong"})
}
