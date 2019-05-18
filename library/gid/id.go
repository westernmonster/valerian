package gid

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/bwmarrin/snowflake"
)

var WorkerID = int64(1)

var generator *snowflake.Node

func NextID() (ts int64, err error) {
	if generator == nil {
		if iw, err := snowflake.NewNode(WorkerID); err != nil {
			return 0, err
		} else {
			generator = iw
		}
	}

	return generator.Generate().Int64(), nil
}

func NewID() (ts int64) {
	if generator == nil {
		generator, _ = snowflake.NewNode(WorkerID)
	}

	return generator.Generate().Int64()
}

func EncodeInt64ToString(id int64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(id))
	return base64.RawURLEncoding.EncodeToString(b)
}

func DecodeStringToInt64(str string) (id int64, err error) {
	bytes, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return
	}
	id = int64(binary.LittleEndian.Uint64(bytes))
	return
}
