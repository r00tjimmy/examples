package main

import (
	"encoding/json"
	"log"
	"strings"

	hello "github.com/micro/examples/puffin/srv/proto/hello"
  ffive "github.com/micro/examples/puffin/srv/proto/ffive"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
)

//---------------------- puffin --------------------------
type Say struct {
	Client hello.SayClient
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.puffin", "Name cannot be blank")
	}

	response, err := s.Client.Hello(ctx, &hello.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}



//-----------------------  ffive --------------------------
type Getapi struct {
  FClient ffive.GetapiClient
}

func (g *Getapi) Apirps(ctx context.Context, req *api.Request, rsp *api.Response) error {
  log.Print("recv getapi.apirps API request")

  ffiveurl, ok := req.Get["ffiveurl"]
  if !ok || len(ffiveurl.Values) == 0 {
    return errors.BadRequest("go.micro.api.buffin", "name cannot be blank")
  }

  response, err := g.FClient.Apirps(ctx, &ffive.Request{Ffiveurl: strings.Join(ffiveurl.Values, " "),
 })

  if err != nil {
    return err
  }

  rsp.StatusCode = 200
  b, _ := json.Marshal(map[string]string{
    "ffiveurl": response.Ffiversp,
  })
  rsp.Body = string(b)

  return nil
}




func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.puffin"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewSayClient("go.micro.srv.puffin", service.Client())},        //puffin
		),

  service.Server().NewHandler(
     &Getapi{FClient: ffive.NewGetapiClient("go.micro.srv.puffin", service.Client())},   //ffive
		),

    //&Getapi{FClient: ffive.NewGetapiClient("go.micro.srv.puffin", service.Client())},   //ffive
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}




