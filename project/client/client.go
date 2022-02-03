package client

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// func newBanner(id int) *b {
// 	b := Banner{id: id}
// 	b.description = ""
// 	return &b
// }

func main() {

	banner := pb.Banner{BannerId: &pb.BannerId{Id: 1}, Description: "test"}
	slot := pb.Slot{SlotId: &pb.SlotId{Id: 324}, Description: "asd"}
	//slotbanner := pb.SlotBanner{BannerId:banner.BannerId,SlotId:slot.SlotId}

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRotatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.AddBanner(ctx, &pb.SlotBanner{BannerId: bannerId, SlotId: slotId})

	//r, err := c.AddBanner(ctx, &pb.SlotBanner{BannerId:&pb.BannerId{Id:34},SlotId:&pb.SlotId{Id:34}})
	r, err := c.AddBanner(ctx, &pb.SlotBanner{BannerId: banner.BannerId, SlotId: slot.SlotId})
	a, err := c.RemoveBanner(ctx, &pb.BannerSlot{SlotId: slot.SlotId, BannerId: banner.BannerId})
	b, err := c.RemoveBanner1(ctx, &pb.SlotBanner{BannerId: banner.BannerId, SlotId: slot.SlotId})
	v, err := c.CountLinkClick(ctx, &pb.SlotBannerUser{SlotBanner: &pb.SlotBanner{BannerId: banner.BannerId, SlotId: slot.SlotId}, UserGroupId: 123})
	f, err := c.SelectBanner(ctx, &pb.UserGroupBanner{SlotId: slot.SlotId, UserGroupsId: &pb.UserGroupsId{Id: 124}})

	if err != nil {
		log.Fatalf("err %v", err)
	}
	log.Printf("BannerID %v, SlotId %v added, %v", banner.BannerId, slot.SlotId, r.ProtoReflect())
	log.Printf("BannerID %v, SlotId %v added, %v", banner.BannerId, slot.SlotId, a.ProtoReflect())
	log.Printf("BannerID %v, SlotId %v added, %v", banner.BannerId, slot.SlotId, b.ProtoReflect())
	log.Printf("BannerID %v, SlotId %v added, %v", banner.BannerId, slot.SlotId, v.ProtoReflect())
	log.Printf("BannerID %v, SlotId %v added, %v", banner.BannerId, slot.SlotId, f.ProtoReflect())
}
