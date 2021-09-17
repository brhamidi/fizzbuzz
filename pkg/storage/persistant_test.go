package storage

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

var mock *miniredis.Miniredis

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return s
}

func setup() {
	mock = mockRedis()
}

func teardown() {
	mock.Close()
}

func TestNewPersistant(t *testing.T) {
	t.Run("should return err if redis.ParseUrl fail", func(t *testing.T) {
		setup()
		defer teardown()

		_, err := NewPersistant("linus", "abcd")

		assert.Error(t, err)
	})

	t.Run("should return do nothing if MAX_KEY is present", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Set(MAX_KEY, "EMPTY")

		_, err := NewPersistant(mock.Host(), mock.Port())

		mock.CheckGet(t, MAX_KEY, "EMPTY")
		assert.NoError(t, err)
	})

	t.Run("should return err if redis.Get fail", func(t *testing.T) {
		setup()
		defer teardown()

		mock.SetError("error")

		_, err := NewPersistant(mock.Host(), mock.Port())

		assert.Error(t, err)
	})

	t.Run("should be ok", func(t *testing.T) {
		setup()
		defer teardown()

		_, err := NewPersistant(mock.Host(), mock.Port())
		assert.NoError(t, err)
	})
}

func TestIncrement(t *testing.T) {
	t.Run("should INCR work well and update MAX_HITS", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewPersistant(mock.Host(), mock.Port())

		mock.Incr("test", 42)

		err := c.Increment("test")
		assert.NoError(t, err)

		maxHits, err := mock.Get(MAX_HITS)
		assert.NoError(t, err)
		assert.Equal(t, "43", maxHits)

		testKey, err := mock.Get("test")
		assert.NoError(t, err)
		assert.Equal(t, "43", testKey)
	})

	t.Run("should INCR work well and not update MAX_HITS", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 21)
		mock.Incr("key2", 42)

		c, _ := NewPersistant(mock.Host(), mock.Port())

		err := c.Increment("key2")
		assert.NoError(t, err)

		maxHits, err := mock.Get(MAX_HITS)
		assert.NoError(t, err)
		assert.Equal(t, "43", maxHits)

		err = c.Increment("key3")
		assert.NoError(t, err)

		maxHits, err = mock.Get(MAX_HITS)
		assert.NoError(t, err)
		assert.Equal(t, "43", maxHits)
	})
}

func TestMax(t *testing.T) {
	t.Run("should return err if driver.Get fail", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewPersistant(mock.Host(), mock.Port())

		mock.SetError("error")
		_, _, err := c.Max()
		assert.Error(t, err)
	})

	t.Run("should be ok", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 21)
		mock.Incr("key2", 42)

		c, _ := NewPersistant(mock.Host(), mock.Port())

		c.Increment("key1")
		c.Increment("key2")

		maxKey, maxHits, err := c.Max()

		assert.NoError(t, err)
		assert.Equal(t, "key2", maxKey)
		assert.Equal(t, 43, maxHits)
	})
}

func TestReset(t *testing.T) {
	t.Run("should return err if driver.FlushDB fail", func(t *testing.T) {
		setup()
		defer teardown()

		c, _ := NewPersistant(mock.Host(), mock.Port())

		mock.SetError("error")
		err := c.Reset()
		assert.Error(t, err)
	})

	t.Run("should be ok", func(t *testing.T) {
		setup()
		defer teardown()

		mock.Incr("key1", 21)
		mock.Incr("key2", 42)

		c, _ := NewPersistant(mock.Host(), mock.Port())

		err := c.Reset()
		assert.NoError(t, err)

		c.Increment("key2")
		c.Increment("key2")

		maxKey, maxHits, _ := c.Max()

		assert.Equal(t, "key2", maxKey)
		assert.Equal(t, 2, maxHits)
	})
}
