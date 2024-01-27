//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go get github.com/snt/dump1090-protobuf@v0.0.4
//go:generate protoc --proto_path=$GOPATH/pkg/mod/github.com/snt/dump1090-protobuf@v0.0.4 --go_out=.. --go_opt=paths=source_relative  --go_opt=Madsb.proto=adsb/ adsb.proto

package build
