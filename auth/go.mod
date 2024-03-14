module auth

go 1.22.0

require sw v1.0.0

require github.com/google/uuid v1.6.0 // indirect

replace sw v1.0.0 => ./go

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lib/pq v1.10.9
)
