package main

import (
	"context"
	"crypto/tls"
	"flag"
	"os"
	"path/filepath"
	"time"

	hello_proto "github.com/kkweon/grpc-as-public-api/server/proto"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func withConfigDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".hello", "client", path)
}

func main() {
	var (
		caCert     = flag.String("ca-cert", withConfigDir("ca.pem"), "Trusted CA certificate.")
		serverAddr = flag.String("server-addr", "localhost:7900", "Hello service address.")
		tlsCert    = flag.String("tls-cert", withConfigDir("cert.pem"), "TLS server certificate.")
		tlsKey     = flag.String("tls-key", withConfigDir("key.pem"), "TLS server key.")
	)
	flag.Parse()

	log.WithFields(log.Fields{
		"caCert":    *caCert,
		"serverAdd": *serverAddr,
		"tlsCert":   *tlsCert,
		"tlsKey":    *tlsKey,
	}).Info("flag.Parse() called")

	//cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	//if err != nil {
	//log.Fatal(err)
	//}

	//log.Print("Loaded X509KeyPair")

	//rawCACert, err := ioutil.ReadFile(*caCert)
	//if err != nil {
	//log.Fatal(err)
	//}

	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(rawCACert)

	creds := credentials.NewTLS(&tls.Config{})
	log.Print("Created a new TLS")

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("gRPC connected to the server")
	defer conn.Close()

	c := hello_proto.NewHelloClient(conn)

	for {

		log.Print("Sending a HelloRequest")
		message, err := c.Say(context.Background(), &hello_proto.HelloRequest{Name: "Kelsey"})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(message.Message)

		time.Sleep(time.Second)
	}

}
