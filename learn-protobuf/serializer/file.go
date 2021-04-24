package serializer

import (
	"fmt"
	"io/ioutil"

	// "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can not marshal prote message to binary : %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can not write binary data to file: %w", err)
	}
	return nil
}

func ReadProteobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can not read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can not unmarshal data from binray file: %w", err)
	}
	return nil
}

// write protobuf to json file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	// convert protobuf message to json
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("can not marshal prote message to json : %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0664)
	if err != nil {
		return fmt.Errorf("cant write JSON data to file : %w", err)
	}
	return nil
}
