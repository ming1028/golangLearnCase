package main

import (
	"errors"
	"github.com/elliotchance/testify-stats/assert"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 4, Add(1, 3))

	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	sl3 := []int{2, 3, 4}
	assert.Equal(t, sl1, sl2, "sl1 should equal to sl2 ")

	p1 := &sl1
	p2 := &sl2
	assert.Equal(t, p1, p2, "the content which p1 point to should equal to which p2 point to")

	err := errors.New("demo error")
	assert.EqualError(t, err, "demo error")

	// 布尔断言
	assert.True(t, 1+1 == 2, "1+1 == 2 should be true")
	assert.Contains(t, "Hello World", "World")
	assert.Contains(t, []string{"Hello", "World"}, "World")
	assert.Contains(t, map[string]string{"Hello": "World"}, "Hello")
	assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})

	// 反向断言
	assert.NotEqual(t, 4, Add(2, 3), "The result should not be 4")
	assert.NotEqual(t, sl1, sl3, "sl1 should not equal to sl3 ")
	assert.False(t, 1+1 == 3, "1+1 == 3 should be false")
	// assert.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond) //1秒之内condition参数都不为true，每10毫秒检查一次
	assert.NotContains(t, "Hello World", "Go")
}
