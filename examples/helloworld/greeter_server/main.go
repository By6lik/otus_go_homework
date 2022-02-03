/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	//"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedRotatorServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddBanner(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
	log.Printf("Received: bannerId: %v slotId: %v", in.GetBannerId() , in.GetSlotId())
	//time.Sleep(time.Second * 5)
	return &empty.Empty{}, nil
}

func (s *server) RemoveBanner(ctx context.Context, in *pb.BannerSlot) (*empty.Empty, error) {
	log.Printf("Removed: slotId: %v bannerId: %v", in.GetSlotId(), in.GetBannerId() )
	return &empty.Empty{}, nil
}

func (s *server) RemoveBanner1(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
	log.Printf("Removed: bannerId: %v slotId: %v", in.GetBannerId() , in.GetSlotId() )
	return &empty.Empty{}, nil
}

func (s *server) CountLinkClick(ctx context.Context, in *pb.SlotBannerUser) (*empty.Empty, error) {
	
	log.Printf("CountLinkClick: SlotBanner: %v UserGroupsId: %v", in.GetSlotBanner() , in.GetUserGroupId() )
	return &empty.Empty{}, nil
}

func (s *server) SelectBanner(ctx context.Context, in *pb.UserGroupBanner) (*pb.BannerId, error) {

	log.Printf("Select banner: SlotId: %v UserGroupsId: %v", in.GetSlotId() , in.GetUserGroupsId())
	return &pb.BannerId{}, nil
}


func main() {
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
