package main

import (
	"log"
	"time"

	hello "github.com/micro/examples/puffin/srv/proto/hello"
  ffive "github.com/micro/examples/puffin/srv/proto/ffive"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)


//---------------- puffin --------------------
type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

//---------------- ffive ---------------------
type Getapi struct{}

func (g *Getapi) Apirps(ctx context.Context, req *ffive.Request, rsp *ffive.Response) error {
	log.Print("get api from ffive")
	rsp.Ffiversp = "F5 api url is: " + req.Ffiveurl
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.puffin"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

  //register f5 handlers
  ffive.RegisterGetapiHandler(service.Server(), new(Getapi))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
