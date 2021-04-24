package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	b := protojson.MarshalOptions{
		UseEnumNumbers: false,
		Indent:         "  ",
		UseProtoNames:  true,
	}

	res, err := b.Marshal(message)
	return string(res), err
}
