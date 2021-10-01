package utils

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc/metadata"
)

var (
	sf        *sonyflake.Sonyflake
	startTime = time.Date(2010, 11, 21, 0, 0, 0, 0, time.UTC)
)

func init() {
	var st sonyflake.Settings
	st.StartTime = startTime
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func NextNumID() uint64 {
	id, _ := sf.NextID()
	return id
}

const IdPrefixKey = "p-id"

func NextID(ctx context.Context) string {
	nextId, err := sf.NextID()
	if err != nil {
		return err.Error()
	}
	// uit64 转成 str
	id := strconv.FormatUint(nextId, 10)

	if ctx == nil {
		ctx = context.Background()
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ps := md.Get(IdPrefixKey); len(ps) != 0 {
			return ps[0] + id
		}
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if ps := md.Get(IdPrefixKey); len(ps) != 0 {
			return ps[0] + id
		}
	}
	return id
}

func UUID() string {
	return uuid.New().String()
}

func UUIDShort() string {
	return strings.Replace(UUID(), "-", "", -1)
}
