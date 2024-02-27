module github.com/kirychukyurii/protoc-gen-go-webitel

go 1.22.0

replace github.com/webitel/protos/wfm => ../protos/wfm

require (
	github.com/webitel/protos/wfm v0.0.0-00010101000000-000000000000
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.32.0-20231115204500-e097f827e652.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240116215550-a9fa1716bcac // indirect
)
