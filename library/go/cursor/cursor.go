package cursor

import "encoding/binary"

func CursorGenerate(createdAt int64, id []byte) []byte {
	cursor := make([]byte, 20)
	binary.BigEndian.PutUint64(cursor[:8], uint64(createdAt))
	copy(cursor[8:], id)
	return cursor
}
