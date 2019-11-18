GOPATH = /root/go
build:
	@go mod vendor
	@go build -mod vendor -v -o messenger main/main.go
proto:
	# If build proto failed, make sure you have protoc installed and:
	# go get -u github.com/google/protobuf
	# go get -u github.com/golang/protobuf/protoc-gen-go
	# go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	# mkdir -p $GOPATH/src/github.com/googleapis && git clone git@github.com:googleapis/googleapis.git $GOPATH/src/github.com/googleapis/
	@mkdir -p pb
	protoc \
	 -I . \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
		--proto_path=${GOPATH}/src/github.com/googleapis/googleapis \
		--proto_path=. \
		--go_out=plugins=grpc:$(PWD)/pb \
		--govalidators_out=$(PWD)/pb \
 		messenger.proto
	$(call color_out,$(CL_ORANGE),"Done")

all:
	build
