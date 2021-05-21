package service

type PingService interface {
}

type pingService struct {
}

func GetPingService() PingService {
	return &pingService{}
}

func (svc *pingService) getPing() {

}
