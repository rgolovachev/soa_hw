module mock_statpb

go 1.22.0

require (
	github.com/golang/mock v1.6.0
	google.golang.org/grpc v1.64.0
	stat v1.0.0
)

require (
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace stat v1.0.0 => ../..
