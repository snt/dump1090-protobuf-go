//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go get github.com/snt/dump1090-protobuf@v0.0.3
//go:generate protoc --proto_path=$GOPATH/pkg/mod/github.com/snt/dump1090-protobuf@v0.0.3 --go_out=../.. --go_opt=Madsb.proto=dump1090-protobuf-go/ adsb.proto

package build
