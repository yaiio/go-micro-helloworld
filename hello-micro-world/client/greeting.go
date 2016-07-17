package main

import (
	"github.com/micro/go-micro"

	"fmt"
	greeterProto "github.com/yaiio/go-micro-helloworld/hello-micro-world/proto"
	"golang.org/x/net/context"
	"os"
)

func runClient(service micro.Service, name string) {
	// Create new greeter client
	greeter := greeterProto.NewGreeterClient("greeter", service.Client())

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
	)

	if len(os.Args) != 2 {
		fmt.Println("greeting {name}")
		return
	}

	runClient(service, os.Args[1])

}
