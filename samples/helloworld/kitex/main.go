package main

import (
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	hello "github.com/kitex-contrib/codec-dubbo/samples/helloworld/kitex/kitex_gen/hello/greetservice"
	"github.com/kitex-contrib/registry-nacos/nacos"
	"github.com/kitex-contrib/registry-nacos/registry"
	xdsmanager "github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// WithAuth sets the username .
func WithAuth() nacos.Option {
	return nacos.Option{F: func(param *vo.NacosClientParam) {
		if param != nil && param.ClientConfig != nil {
			param.ClientConfig.Username = os.Getenv("NACOS_ENV_USERNAME")
			param.ClientConfig.Password = os.Getenv("NACOS_ENV_PASSWARD")
		}
	}}
}

func main() {
	cli, err := nacos.NewDefaultNacosClient(WithAuth())
	if err != nil {
		klog.Fatal(err)
	}
	if err := xdsmanager.Init(); err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ":9102")
	if err != nil {
		klog.Fatal(err)
	}
	go func() {
		svr := hello.NewServer(new(GreetServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
				ServiceName: "org.cloudwego.kitex.samples.api.GreetProvider",
				Tags: map[string]string{
					"sidecar.mesh.io/protocol": "dubbo",
				},
			}),
			server.WithServiceAddr(addr),
			server.WithReusePort(true),
			server.WithRegistry(registry.NewNacosRegistry(cli)),
			server.WithSuite(xdssuite.NewServerSuite()),
			server.WithCodec(dubbo.NewDubboCodec(
				dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
			)),
		)

		err = svr.Run()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	svr := hello.NewServer(new(GreetServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "hello",
			Tags: map[string]string{
				"sidecar.mesh.io/protocol": "dubbo",
			},
		}),
		server.WithServiceAddr(addr),
		server.WithReusePort(true),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithSuite(xdssuite.NewServerSuite()),
		server.WithCodec(dubbo.NewDubboCodec(
			dubbo.WithJavaClassName("org.cloudwego.kitex.samples.api.GreetProvider"),
		)),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
