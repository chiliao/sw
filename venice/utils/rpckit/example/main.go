//+build !test

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
	exproto "github.com/pensando/sw/venice/utils/rpckit/example/proto"
	"github.com/pensando/sw/venice/utils/rpckit/tlsproviders"
)

// Usage:
// 0. Compile the example: go install
// 1. Run server in a window: $GOPATH/bin/example -server -url :443
// 2. Run client in another window: $GOPATH/bin/example -client -url localhost:443

// ExampleRPCHandler is the grpc handler
type ExampleRPCHandler struct {
	srvMsg string
}

// ExampleRPC is example rpc call handler
func (es *ExampleRPCHandler) ExampleRPC(ctx context.Context, req *exproto.ExampleReq) (*exproto.ExampleResp, error) {
	log.Infof("Example server got request: %+v", req)
	exampleResp := &exproto.ExampleResp{
		RespMsg: es.srvMsg,
	}

	return exampleResp, nil
}

// runServer runs the server
func runServer(url string, stopChannel chan bool, tlsProvider rpckit.TLSProvider) {
	// create an rpc handler object
	exampleHandler := &ExampleRPCHandler{
		srvMsg: "Example Server Response",
	}

	// create an RPC server
	rpcServer, err := rpckit.NewRPCServer("grpc.local", url, rpckit.WithTLSProvider(tlsProvider))
	if err != nil {
		log.Fatalf("Error creating rpc server. Err; %v", err)
	}

	// register the RPC handler
	exproto.RegisterExampleServer(rpcServer.GrpcServer, exampleHandler)
	defer func() { rpcServer.Stop() }()

	// wait forever
	<-stopChannel
}

// runClient runs the client
func runClient(url string, count int, tlsProvider rpckit.TLSProvider) {
	// create an RPC client
	rpcClient, err := rpckit.NewRPCClient("grpc.local", url, rpckit.WithTLSProvider(tlsProvider))
	if err != nil {
		log.Errorf("Error connecting to server. Err: %v", err)
		return
	}

	exampleClient := exproto.NewExampleClient(rpcClient.ClientConn)
	req := exproto.ExampleReq{ReqMsg: "example request"}

	// make calls `count` times
	for i := 0; i < count; i++ {
		// make a call
		resp, err := exampleClient.ExampleRPC(context.Background(), &req)
		if err != nil {
			log.Errorf("Got RPC error: %v", err)
			return
		}

		log.Infof("Got RPC response: %+v", resp)
	}

	log.Infof("RPC stats: %+v", rpckit.Stats())

	// close client connection and stop the server
	rpcClient.Close()
	time.Sleep(time.Second)
}

// main function
func main() {
	var clientFlag, serverFlag bool
	var url string
	var certFile, keyFile, caFile string
	var count int

	// parse all flags
	flag.BoolVar(&clientFlag, "client", false, "Run in client mode")
	flag.BoolVar(&serverFlag, "server", false, "Run in server mode")
	flag.StringVar(&url, "url", "localhost:9100", "URL for server or client")
	flag.StringVar(&certFile, "cert", "", "Certificate file")
	flag.StringVar(&keyFile, "key", "", "Key file")
	flag.StringVar(&caFile, "ca", "", "Root CA file")
	flag.IntVar(&count, "count", 1, "Number of times to make rpc call")

	flag.Parse()

	var tlsProvider rpckit.TLSProvider
	var err error
	if certFile != "" {
		tlsProvider, err = tlsproviders.NewFileBasedProvider(certFile, keyFile, caFile)
		if err != nil {
			log.Infof("Error initializing TLS provider: %v", err)
			os.Exit(1)
		}
		log.Info("Initialized TLS provider")
	}

	// make sure necessary flags are specified
	if !serverFlag && !clientFlag {
		fmt.Printf("Need to run as either -client or -server\n  See example --help for more details\n")
		os.Exit(1)
	} else if serverFlag && clientFlag {
		fmt.Printf("Can not specify both -client and -server flags\n")
		os.Exit(1)
	} else if serverFlag {
		runServer(url, make(chan bool), tlsProvider)
	} else {
		runClient(url, count, tlsProvider)
	}
}
