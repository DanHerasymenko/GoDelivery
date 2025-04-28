package services

import (
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/services/auth"
)

type Services struct {
	Auth *auth.Service
}

func NewServices(cfg *config.Config, clnts *clients.Clients) *Services {
	return &Services{
		Auth: auth.NewService(cfg, clnts),
	}
}
