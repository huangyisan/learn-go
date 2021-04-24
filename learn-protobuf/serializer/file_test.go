package serializer_test

import (
	"learn-protobuf/sample"
	"learn-protobuf/serializer"
	"testing"

	pb "learn-protobuf/pb"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	jsonfile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProteobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	// 确认laptop1和laptop2是否一致
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonfile)
	require.NoError(t, err)

}
