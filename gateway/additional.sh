go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:$HOME/go/bin

source ~/.bashrc

which protoc-gen-go
which protoc-gen-go-grpc
