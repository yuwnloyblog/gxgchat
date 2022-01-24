package tools

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

func String2Bytes(s string) []byte {
	// reader := strings.NewReader(s)
	// bytes := make([]byte, reader.Size())
	// reader.ReadAt(bytes, 0)
	bytes := []byte(s)
	return bytes
}

func Bytes2String(bytes []byte) string {
	// sb := strings.Builder{}
	// sb.Write(bytes)
	// return sb.String()
	return string(bytes)
}

func Bytes2Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

func Int2Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint32(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

func PbMarshal(obj proto.Message) ([]byte, error) {
	bytes, err := proto.Marshal(obj)
	return bytes, err
}
func PbUnMarshal(bytes []byte, typeScope proto.Message) error {
	err := proto.Unmarshal(bytes, typeScope)
	return err
}

func JsonMarshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func JsonUnMarshal(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, obj)
}
