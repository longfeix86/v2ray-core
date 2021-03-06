package protocol_test

import (
	"testing"
	"time"

	. "v2ray.com/core/common/protocol"
	"v2ray.com/core/testing/assert"
)

func TestGenerateRandomInt64InRange(t *testing.T) {
	assert := assert.On(t)

	base := time.Now().Unix()
	delta := 100
	generator := NewTimestampGenerator(Timestamp(base), delta)

	for i := 0; i < 100; i++ {
		v := int64(generator())
		assert.Int64(v).AtMost(base + int64(delta))
		assert.Int64(v).AtLeast(base - int64(delta))
	}
}
