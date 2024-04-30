package service

import (
	"context"
	"go-bponline/m/pb/health"
)
type Health struct {
	health.UnimplementedHealthServiceServer
}

func (m *Health) GetHealth(ctx context.Context, req *health.HealthReq) (*health.HealthRes, error) {
	return &health.HealthRes{
		Status: "OK",
	}, nil
}