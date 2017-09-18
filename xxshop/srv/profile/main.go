package main

import (
  "github.com/micro/examples/booking/srv/profile/proto"
  "context"
  "github.com/hashicorp/consul/vendor/google.golang.org/grpc/metadata"
  "golang.org/x/net/trace"
  "github.com/micro/examples/booking/data"
  "encoding/json"
  "github.com/prometheus/common/log"
  "github.com/micro/go-micro"
)

type Profile struct {
  hotels map[string]*profile.Hotel
}


func (s *Profile) GetProfiles(ctx context.Context, req *profile.Request, rsp *profile.Result) error {
  md, _ := metadata.FromContext(ctx)
  traceID := md["traceID"]
  if tr, ok := trace.FromContext(ctx); ok {
    tr.LazyPrintf("traceID %s", traceID)
  }

  for _, i := range req.HotelIds {
    rsp.Hotels = append(rsp.Hotels, s.hotels[i])
  }

  return  nil
}



func loadProfiles(path string) map[string]*profile.Hotel {
  file := data.MustAsset(path)

  hotels := []*profile.Hotel{}
  if err := json.Unmarshal(file, &hotels); err != nil {
    log.Fatalf("failed to load json: %v", err)
  }

  profiles := make(map[string]*profile.Hotel)

  for _, hotel := range hotels {
    profiles[hotel.Id] = hotel
  }

  return  profiles
}



func main() {
  service := micro.NewService(
    micro.Name("go.micro.srv.profile"),
  )

  service.Init()

  /**
  这里注意，下面的说明：
  1. Profile 对应 proto 文件的 service Profile
  2. hotels: loadProfiles 作用是装备数据到hotels, 对应 proto 文件
     Profile return 的类型是 Result， 而 Result 返回的又是 Hotel 类型的列表
     (repeated 就是列表的意思)， 这样下面的逻辑就清晰了

   */
  profile.RegisterProfileHandler(service.Server(), &Profile{
    hotels: loadProfiles("data/profiles.json"),
  })

  service.Run()
}










