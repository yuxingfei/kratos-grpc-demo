package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pb "kratos-server/api"
	"kratos-server/internal/dao"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.NewsServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// GetNews demo .
func (s *Service) GetNews(ctx context.Context,in *emptypb.Empty) (reply *pb.NewsResp,err error) {
	new01 := pb.NewsResp_Info{
		Id:      1,
		Title:   "new01",
		Content: "first news",
	}

	new02 := pb.NewsResp_Info{
		Id:      2,
		Title:   "new02",
		Content: "second news",
	}

	new03 := pb.NewsResp_Info{
		Id:      3,
		Title:   "new03",
		Content: "three news",
	}

	reply = &pb.NewsResp{
		News: []*pb.NewsResp_Info{
			&new01,
			&new02,
			&new03,
		},
	}

	err = nil

	fmt.Println("GetNews GRPC 服务 调用", time.Now().Format("2006-01-02:15:04:05"))
	return
}

// GetNewsInfoById demo.
func (s *Service) GetNewsInfoById(ctx context.Context, req *pb.IdReq) (reply *pb.NewsInfoResp,err error) {
	idStr := strconv.Itoa(int(req.Id))
	reply = &pb.NewsInfoResp{
		Id:req.Id,
		Title:"news" + idStr,
		Content:"get news info by id " + idStr,
	}
	err = nil
	fmt.Println("GetNewsInfoById GRPC 服务 调用", time.Now().Format("2006-01-02:15:04:05"))
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
