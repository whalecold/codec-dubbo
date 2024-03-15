package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello"
	"github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	xdsmanager "github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
)

func main() {
	if err := xdsmanager.Init(); err != nil {
		klog.Fatal(err)
	}

	cli, err := greetservice.NewClient("org.cloudwego.kitex.samples.api.GreetProvider",
		//client.WithHostPorts("127.0.0.1:21001"),
		client.WithSuite(xdssuite.NewClientSuite(xdssuite.WithRouterMetaExtractor(tracing.ExtractFromPropagator))),
		client.WithCodec(
			dubbo.NewDubboCodec(
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
			),
		),
	)
	if err != nil {
		panic(err)
	}

	for {
		resp, err := cli.Greet(context.Background(), "world")
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Infof("resp: %s", resp)

		respWithStruct, err := cli.GreetWithStruct(context.Background(), &hello.GreetRequest{Req: "world"})
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Infof("respWithStruct: %s", respWithStruct.Resp)

		time.Sleep(time.Millisecond * 100)
	}
}
