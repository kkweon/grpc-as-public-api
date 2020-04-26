
.PHONY: proto-build
proto-build:
	protoc -I proto server/proto/hello_service.proto --go_out=plugins=grpc:server/proto


server-build:
	cd server && go build

.PHONY: server-run
server-run: server-build
	./server/server -ca-cert cert/ca.crt -tls-cert cert/server.crt -tls-key cert/server.key


client-build:
	cd client && go build

.PHONY: client-run
client-run: client-build
	./client/client -ca-cert cert/ca.crt -tls-cert cert/client.crt -tls-key cert/client.key

.PHONY: client-run-mkube
client-run-mkube: client-build
	./client/client -server-addr=mkube:30000 -tls-cert cert/client.crt -tls-key cert/client.key -ca-cert cert/ca.crt

.PHONY: gen-cert
gen-cert: clean-cert
	sh gen-cert.sh

.PHONY: clean-cert
clean-cert:
	rm -rf cert ca.* client.* server.*
