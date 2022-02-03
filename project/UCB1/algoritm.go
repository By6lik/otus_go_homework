package algoritm

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"strconv"
)

// var (
// 	port = flag.Int("port", 50051, "The server port")
// )

// // server is used to implement helloworld.GreeterServer.
// type server struct {
// 	pb.UnimplementedRotatorServer
// }

// // SayHello implements helloworld.GreeterServer
// func (s *server) AddBanner(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
// 	log.Printf("Received: bannerId: %v slotId: %v", in.GetBannerId() , in.GetSlotId())
// 	//time.Sleep(time.Second * 5)
// 	return &empty.Empty{}, nil
// }

// func (s *server) RemoveBanner(ctx context.Context, in *pb.BannerSlot) (*empty.Empty, error) {
// 	log.Printf("Removed: slotId: %v bannerId: %v", in.GetSlotId(), in.GetBannerId() )
// 	return &empty.Empty{}, nil
// }

// func (s *server) RemoveBanner1(ctx context.Context, in *pb.SlotBanner) (*empty.Empty, error) {
// 	log.Printf("Removed: bannerId: %v slotId: %v", in.GetBannerId() , in.GetSlotId() )
// 	return &empty.Empty{}, nil
// }

// func (s *server) CountLinkClick(ctx context.Context, in *pb.SlotBannerUser) (*empty.Empty, error) {
// 	log.Printf("CountLinkClick: SlotBanner: %v UserGroupsId: %v", in.GetSlotBanner() , in.GetUserGroupId() )
// 	return &empty.Empty{}, nil
// }

// func (s *server) SelectBanner(ctx context.Context, in *pb.UserGroupBanner) (*pb.BannerId, error) {
// 	log.Printf("Select banner: SlotId: %v UserGroupsId: %v", in.GetSlotId() , in.GetUserGroupsId())
// 	return &pb.BannerId{}, nil
// }
type EventType int32

const (
	CLICK EventType = 0
	SHOW  EventType = 1
)

type clickevent struct {
	SlotId      int64
	BannerId    int64
	UserGroupId int64
	ActionType  EventType
}

type banner struct {
	id     int64
	slotid int64
}

//var BannerList = []banner{}
var BannerDict map[int64]interface{}

//BannerDict =  make(map[int64]interface{})

var EventList = []clickevent{}

// func CountClick(SlotId,BannerId,UserGroupId int64) {
// 	M := clickevent{SlotId:SlotId,BannerId:BannerId,UserGroupId:UserGroupId}
// 	Vector = append(Vector, M)
// }
func AddBanner(bannerId int64, slotid int64) {
	//b := banner{id: bannerId, slotid: slotid}

	//BannerList = append(BannerList, b)
	if BannerDict[slotid] == nil {
		BannerDict[slotid] = []int64{}
	}
	BannerDict[slotid] = append(BannerDict[slotid].([]int64), bannerId)
}

func CountClick(SlotId, BannerId, UserGroupId int64) {
	M := clickevent{SlotId: SlotId, BannerId: BannerId, UserGroupId: UserGroupId, ActionType: CLICK}
	EventList = append(EventList, M)
}

func SelectBanner(SlotId, UserGroupId int64) int64 {
	//func SelectBanner() int64 {
	// if len(BannerList) == 0 {
	// 	return -1
	// }
	// resultBannerID := BannerList[0].id
	max_profitability := float64(0)

	BannerList := BannerDict[SlotId].([]int64)
	resultBannerID := BannerList[0]
	for _, bannerId := range BannerList {
		clicksCount := 0
		shoWbanner := 0
		shoWAllbanners := 0
		for _, i := range EventList {
			if bannerId == i.BannerId && i.ActionType == CLICK && i.SlotId == SlotId && i.UserGroupId == UserGroupId {
				clicksCount += 1
			}
			if bannerId == i.BannerId && i.ActionType == SHOW && i.SlotId == SlotId && i.UserGroupId == UserGroupId {
				shoWbanner += 1
			}
			if i.ActionType == SHOW && i.SlotId == SlotId && i.UserGroupId == UserGroupId {
				shoWAllbanners += 1
			}
		}
		profitability := float64(clicksCount) + math.Sqrt(2*math.Log(float64(shoWAllbanners))/float64(shoWbanner))
		fmt.Printf("profitability: %v \n", profitability)

		if math.IsNaN(profitability) {
			profitability = math.Inf(1)
		}
		if profitability > max_profitability {
			resultBannerID = bannerId
			max_profitability = profitability
		}
	}

	//расчет в результате которого у нас появляется id banner, который мы показываем

	N := clickevent{BannerId: resultBannerID, ActionType: SHOW, SlotId: SlotId, UserGroupId: UserGroupId}
	EventList = append(EventList, N)
	return resultBannerID
}

func qwe() {
	BannerDict = make(map[int64]interface{})

	AddBanner(1, 0)
	AddBanner(2, 0)
	AddBanner(3, 1)
	AddBanner(2, 1)
	fmt.Println(BannerDict)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Provide Usergroup:\n")
		scanner.Scan()
		usergroup, _ := strconv.Atoi(scanner.Text())
		for slotid := 0; slotid <= 1; slotid++ {

			a := SelectBanner(int64(slotid), int64(usergroup))

			fmt.Printf("You see baneerID: %v, slotId: %v, groupID: %v do you want to click banner: Y/n\n", a, slotid, usergroup)
			scanner.Scan()
			if scanner.Text() == "y" {
				CountClick(int64(slotid), a, int64(usergroup))
				fmt.Printf("You clicked %v\n", a)
			}
		}
		fmt.Println(EventList)
	}

}
