package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"context"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	hello_proto "github.com/kkweon/grpc-as-public-api/server/proto"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".hello", "server", path)
}

type helloWorldServer struct {
}

func (h *helloWorldServer) Say(ctx context.Context, helloRequest *hello_proto.HelloRequest) (*hello_proto.HelloResponse, error) {

	log.Print(fmt.Sprintf("Received a HelloRequest %v", helloRequest))

	return &hello_proto.HelloResponse{
		Message: "Hello " + helloRequest.GetName(),
	}, nil
}

const tcp = "tcp"

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		listenAddr = flag.String("listen-addr", "0.0.0.0:7900", "HTTP listen address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	log.Println("Hello service starting...")

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}

	rawCaCert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCaCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	server := grpc.NewServer(grpc.Creds(creds))
	hello_proto.RegisterHelloServer(server, &helloWorldServer{})

	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.helloservice", 1)
	healthpb.RegisterHealthServer(server, healthServer)

	listener, err := net.Listen(tcp, *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server will listen to %v", *listenAddr)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
