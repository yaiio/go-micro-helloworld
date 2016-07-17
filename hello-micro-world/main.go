package main

import (
	"github.com/micro/go-micro"

	"fmt"
	"github.com/micro/cli"
	greeterProto "github.com/yaiio/go-micro-helloworld/hello-micro-world/proto"
	"golang.org/x/net/context"
	"log"
	"os"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeterProto.HelloRequest, rsp *greeterProto.HelloResponse) error {
	rsp.Greeting = "Hello, " + req.Name
	return nil
}

func runClient(service micro.Service, c *cli.Context) {
	// Create new greeter client
	greeter := greeterProto.NewGreeterClient("greeter", service.Client())

	name := "World"

	if flagName := c.String("name"); flagName != "" {
		name = flagName
	}

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &greeterProto.HelloRequest{Name: name})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main() {

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("lastest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(
			cli.BoolFlag{
				Name:  "client",
				Usage: "Launch the client",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "Hello Name",
			},
		),
	)

	service.Init( // Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) {
			if c.Bool("client") {
				runClient(service, c)
				os.Exit(0)
			}
		},
		))

	greeterProto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
