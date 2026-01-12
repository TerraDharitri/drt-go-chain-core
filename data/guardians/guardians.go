//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/dharitri/protobuf/protobuf --gogoslick_out=. guardians.proto
package guardians
