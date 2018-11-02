package servicedef

// generate gRPC code (invoked with 'go generate')
//go:generate sh -c "docker run --rm -v`pwd`:`pwd` -w`pwd` znly/protoc:0.3.0 -I. --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. servicedef.proto"
