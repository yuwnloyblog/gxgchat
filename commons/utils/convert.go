package utils

import (
	"bytes"
	"encoding/binary"
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
