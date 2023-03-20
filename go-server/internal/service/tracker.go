package service

import (
	"context"
	pb "go-server/api/tracker/v1"
	"go-server/internal/data"
)

type TrackerService struct {
	pb.UnimplementedTrackerServer
	repo *data.TrackerRepo 
}

func NewTrackerService(g *data.TrackerRepo) *TrackerService {
	return &TrackerService{repo:g}
}

func (s *TrackerService) CreateBlockFunc(ctx context.Context, req *pb.CreateBlock) (*pb.BlockResp, error) {
	return &pb.BlockResp{}, nil
}
func (s *TrackerService) DeletBlockFunc(ctx context.Context, req *pb.DeletBlock) (*pb.Status, error) {
	return &pb.Status{}, nil
}
func (s *TrackerService) UpdateBlockFunc(ctx context.Context, req *pb.UpdateBlock) (*pb.UpdateResponse, error) {
	return &pb.UpdateResponse{}, nil
}
func (s *TrackerService) ListBlockFunc(req *pb.GetListReq, conn pb.Tracker_ListBlockFuncServer) error {
	for {
		err := conn.Send(&pb.BlockResp{})
		if err != nil {
			return err
		}
	}
}
