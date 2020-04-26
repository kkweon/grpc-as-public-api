package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	hello_proto "github.com/kkweon/grpc-as-public-api/server/proto"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".hello", "client", path)
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		serverAddr = flag.String("server-addr", "127.0.0.1:7900", "Hello service address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		log.Fatal(err)
	}

	rawCACert, err := ioutil.ReadFile(*caCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(rawCACert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := hello_proto.NewHelloClient(conn)
	message, err := c.Say(context.Background(), &hello_proto.HelloRequest{"Kelsey"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(message.Message)
}
