package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	//"time"

	pb ""
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedRotatorServer
}

func (s *server) AddBanner(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
	log.Printf("Received: bannerId: %v slotId: %v", in.GetBannerId(), in.GetSlotId())
	//time.Sleep(time.Second * 5)
	return &empty.Empty{}, nil
}

func (s *server) RemoveBanner(ctx context.Context, in *pb.BannerSlot) (*empty.Empty, error) {
	log.Printf("Removed: slotId: %v bannerId: %v", in.GetSlotId(), in.GetBannerId())
	return &empty.Empty{}, nil
}

func (s *server) RemoveBanner1(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
	log.Printf("Removed: bannerId: %v slotId: %v", in.GetBannerId(), in.GetSlotId())
	return &empty.Empty{}, nil
}

func (s *server) CountLinkClick(ctx context.Context, in *pb.SlotBannerUser) (*empty.Empty, error) {

	log.Printf("CountLinkClick: SlotBanner: %v UserGroupsId: %v", in.GetSlotBanner(), in.GetUserGroupId())
	return &empty.Empty{}, nil
}

func (s *server) SelectBanner(ctx context.Context, in *pb.UserGroupBanner) (*pb.BannerId, error) {

	log.Printf("Select banner: SlotId: %v UserGroupsId: %v", in.GetSlotId(), in.GetUserGroupsId())
	return &pb.BannerId{}, nil
}

func server() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRotatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
