package service

import (
	"github.com/golang/protobuf/proto"
)

// gRPC codec options
type ProtoCodec struct{}

func (ProtoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (ProtoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (ProtoCodec) Name() string {
	return "proto"
}
