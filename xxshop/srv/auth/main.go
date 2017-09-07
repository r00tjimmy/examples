package main

import (
  "encoding/json"
  "errors"
  "log"

  "github.com/micro/examples/xxshop/data"
  "github.com/micro/examples/xxshop/srv/auth/proto"

  "golang.org/x/net/context"
  "golang.org/x/net/trace"

  "github.com/micro/go-micro"
  "github.com/micro/go-micro/metadata"
)

type Auth struct {
  customers map[string]*auth.Customer  //*auth.Customer是用了auth package 里面的 Customer struct
}

func (s *Auth) VerifyToken(ctx context.Context, req *auth.Request, rsp *auth.Result) error {
  // *auth.Request 是引用了 auth package 里面的 Request struct
  md, _ := metadata.FromContext(ctx)
  traceID := md["traceID"]

  if tr, ok := trace.FromContext(ctx); ok {
    tr.LazyPrintf("traceID %s", traceID)
  }

  customer := s.customers[req.AuthToken]
  if customer == nil {
    return  errors.New("Invalid Token")
  }

  rsp.Customer = customer
  return nil
}

//.....
func loadCustomerData(path string) map[string]*auth.Customer {
  file := data.MustAsset(path)
  customers := []*auth.Customer{}

  if err := json.Unmarshal(file, &customers); err != nil {
    log.Fatalf("Failed to umarshal json: %v", err)
  }

  cache := make(map[string]*auth.Customer)
  for _, c := range customers {
    cache[c.AuthToken] = c
  }
  return cache
}


func main() {
  service := micro.NewService(
    micro.Name("xxshop.srv.auth"),
  )

  service.Init()

  auth.RegisterAuthHandler(
    service.Server(),
    &Auth{ customers: loadCustomerData("data/customers.json"), },
  )

  service.Run()
}




















