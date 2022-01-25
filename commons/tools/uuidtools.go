package tools

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	uid := uuid.New()
	return uid
}

func GenerateUUIDBytes() []byte {
	uid, _ := uuid.NewUUID()
	return []byte(uid.String())
}

func UUIDStringByBytes(bytes []byte) (string, error) {
	uuid, err := uuid.FromBytes(bytes)
	return uuid.String(), err
}

func GenerateUUIDShortString() string {
	return UUID2ShortString(GenerateUUID())
}

func UUID2ShortString(uuid uuid.UUID) string {
	mostBits := make([]byte, 8)
	leastBits := make([]byte, 8)
	for i := 0; i < 8; i++ {
		mostBits[i] = uuid[i]
		fmt.Print(uuid[i])
		fmt.Print(",")
	}
	for i := 8; i < 16; i++ {
		leastBits[i-8] = uuid[i]
		fmt.Print(uuid[i])
		fmt.Print(",")
	}
	fmt.Println()
	mostLong := BytesToUInt64(mostBits)
	leastLong := BytesToUInt64(leastBits)
	fmt.Println("most:", mostLong, "\tleast:", leastLong)
	return strings.Join([]string{toIdString(BytesToUInt64(mostBits)), toIdString(BytesToUInt64(leastBits))}, "")
}

var DIGITS64 []byte = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")

func toIdString(l uint64) string {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	var length int = 11
	var least uint64 = 63 //0x3f

	for {
		length--
		buf[length] = DIGITS64[int(l&least)]
		l = l >> 6
		if l == 0 {
			break
		}
	}
	return string(buf)
}
