
.PHONY: proto-build
proto-build:
	protoc -I proto server/proto/hello_service.proto --go_out=plugins=grpc:server/proto


.PHONY: server-build
server-build:
	cd server && go build


.PHONY: client-build
client-build:
	cd client && go build


.PHONY: gen-cert
gen-cert:
	sh gen-cert.sh

.PHONY: clean-cert
clean-cert:
	rm ca.* client.* server.*
