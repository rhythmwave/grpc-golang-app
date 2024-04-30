package route

import (
	"go-bponline/m/internal/service"
	"go-bponline/m/pb/anggaran"
	"go-bponline/m/pb/health"
	"go-bponline/m/pb/kegiatan"
	
	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server) {
	// Health Service
	healthServer := service.Health{}
	health.RegisterHealthServiceServer(grpcServer, &healthServer)
	
	// Kegiatan Service
	kegiatanServer := service.Kegiatan{}
	kegiatan.RegisterKegiatanServiceServer(grpcServer, &kegiatanServer)
	
	// Anggaran Service
	anggaranServer := service.AnggaranService{}
	anggaran.RegisterAnggaranServiceServer(grpcServer, &anggaranServer)
	
}
