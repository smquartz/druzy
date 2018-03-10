package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/smquartz/druzy/handler"
	"github.com/smquartz/druzy/subscriber"

	example "github.com/smquartz/druzy/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("video.quartz.srv.druzy"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.srv.druzy", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.srv.druzy", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
