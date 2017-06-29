package main

import (
	"log"

	"github.com/emicklei/go-restful"

	hello "github.com/micro/examples/puffin/srv/proto/hello"
	ffive "github.com/micro/examples/puffin/srv/proto/ffive"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"golang.org/x/net/context"
)

//------------------------- puffin -----------------------
type Say struct{}

var (
	cl hello.SayClient
)

func (s *Say) Anything(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Anything API request")
	rsp.WriteEntity(map[string]string{
		"message": "Hi, this is the puffin API",
	})
}

func (s *Say) Hello(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Hello API request")

	name := req.PathParameter("name")

	response, err := cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}


//------------------------- ffive -----------------------
type Getapi struct{}

var (
	fcl ffive.GetapiClient
)


func (g *Getapi) Apirps(req *restful.Request, rsp *restful.Response) {
	log.Print("Received ffvie-getapi.apirps API request")

	ffiveurl := req.PathParameter("ffiveurl")

	response, err := fcl.Apirps(context.TODO(), &ffive.Request{
		Ffiveurl: ffiveurl,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}



func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.puffin"),
	)

	service.Init()

	// setup puffin Server Client
	cl = hello.NewSayClient("go.micro.srv.puffin", client.DefaultClient)

  fcl = ffive.NewGetapiClient("go.micro.srv.puffin", client.DefaultClient)

	// Create RESTful handler
	say := new(Say)
  getapi := new(Getapi)

	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/puffin")

	ws.Route(ws.GET("/").To(say.Anything))
	ws.Route(ws.GET("/{name}").To(say.Hello))

  ws.Route(ws.GET("/{ffiveurl}").To(getapi.Apirps))

	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}




