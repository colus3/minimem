package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_listStoreFunc(t *testing.T) {
	store := NewListStore()
	store.rightPush("key", "value3")
	store.rightPush("key", "value4")
	store.rightPush("key", "value5")
	store.leftPush("key", "value2")
	store.leftPush("key", "value1")
	store.leftPush("key", "value0")


	assert.Equal(t, "value0", store.leftPeek("key"))
	assert.Equal(t, "value5", store.rightPeek("key"))
	assert.Equal(t, "value0", store.leftPop("key"))
	assert.Equal(t, "value5", store.rightPop("key"))

	list := store.rangeGet("key", 0, 2)

	assert.Equal(t, "value1", list[0])
	assert.Equal(t, "value2", list[1])
}