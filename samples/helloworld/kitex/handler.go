package main

import (
	"context"
	"time"

	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// Greet implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) Greet(ctx context.Context, req string) (resp string, err error) {
	time.Sleep(time.Millisecond * 10)
	return "Hello " + req, nil
}

// GreetWithStruct implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) GreetWithStruct(ctx context.Context, req *hello.GreetRequest) (resp *hello.GreetResponse, err error) {
	time.Sleep(time.Millisecond * 10)
	return &hello.GreetResponse{Resp: "Hello " + req.Req}, nil
}
