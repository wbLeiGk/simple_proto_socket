# simple_proto_socket
just implement simple proto in socket 


proto:

1st byte : version

2nd - 4th: data length

5th ~ : data content




test:
    project root path :  go test ./...

run :

    project root path ： go run server/server_simple_proto.go   and   go run client/client_simple_proto.go