package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInmemory_Increment(t *testing.T) {
	mem := NewInmemory()

	mem.Increment("key1")
	mem.Increment("key1")
	mem.Increment("key2")
	mem.Increment("key3")

	assert.Len(t, mem.data, 3)
	assert.Equal(t, 2, mem.data["key1"])
	assert.Equal(t, 1, mem.data["key2"])
	assert.Equal(t, 1, mem.data["key3"])
}

func TestInmemory_Max(t *testing.T) {
	mem := NewInmemory()

	_, hits, err := mem.Max()
	assert.Equal(t, 0, hits)
	assert.NoError(t, err)

	mem.Increment("key1")
	mem.Increment("key1")
	mem.Increment("key2")
	mem.Increment("key2")
	mem.Increment("key3")
	mem.Increment("key3")
	mem.Increment("key3")
	mem.Increment("key3")

	key, hits, _ := mem.Max()

	assert.Equal(t, key, "key3")
	assert.Equal(t, 4, hits)
}

func TestInmemory_Reset(t *testing.T) {
	mem := NewInmemory()

	mem.Increment("key1")
	mem.Increment("key2")
	mem.Increment("key2")
	mem.Increment("key3")

	_ = mem.Reset()

	key, hits, err := mem.Max()

	assert.Equal(t, key, "")
	assert.Equal(t, 0, hits)
	assert.NoError(t, err)
}
