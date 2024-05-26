module auth

go 1.22.0

require (
	github.com/gorilla/mux v1.8.1
	github.com/lib/pq v1.10.9
	posts v1.0.0
	auth v1.0.0
)

replace posts v1.0.0 => ../posts
replace auth v1.0.0 => ./go

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.6.0
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0 // indirect
)
