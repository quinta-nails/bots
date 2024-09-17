module github.com/quinta-nails/telegram-backend

go 1.23.0

require (
	github.com/amacneil/dbmate/v2 v2.20.0
	github.com/caarlos0/env/v11 v11.2.2
	github.com/joho/godotenv v1.5.1
	github.com/quinta-nails/protobuf v0.0.1
	google.golang.org/grpc v1.66.2
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240717164558-a6c49f84cc0f.2 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/bufbuild/protovalidate-go v0.6.5 // indirect
	github.com/google/cel-go v0.21.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20240808152545-0cdaa3abc0fa // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/quinta-nails/protobuf v0.0.1 => ../protobuf
