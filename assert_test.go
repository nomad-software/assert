package assert

import (
	"errors"
	"testing"
	"unsafe"
)

func TestTrue(t *testing.T) {
	True(t, true)
	True(t, !false)
}

func TestFalse(t *testing.T) {
	False(t, false)
	False(t, !true)
}

func TestEqual(t *testing.T) {
	type Foo struct {
		foo string
		bar int
	}

	Eq(t, 1, 1)
	Eq(t, "foo", "foo")
	Eq(t, Foo{foo: "foo", bar: 1337}, Foo{foo: "foo", bar: 1337})
}

func TestGreaterThan(t *testing.T) {
	Gt(t, 2, 1)
	Gt(t, 20, 10)
}

func TestGreaterThanOrEqual(t *testing.T) {
	Gte(t, 2, 1)
	Gte(t, 1, 1)
	Gte(t, 20, 10)
	Gte(t, 10, 10)
}

func TestLessThan(t *testing.T) {
	Lt(t, 1, 2)
	Lt(t, 10, 20)
}

func TestLessThanOrEqual(t *testing.T) {
	Lte(t, 1, 2)
	Lte(t, 1, 1)
	Lte(t, 10, 20)
	Lte(t, 10, 10)
}

func TestLen(t *testing.T) {
	Len(t, []int{1, 2, 3, 4, 5}, 5)
	Len(t, []string{"foo", "bar"}, 2)
}

func TestEmpty(t *testing.T) {
	Empty(t, []int{})
	Empty(t, []string{})
}

func TestContains(t *testing.T) {
	Contains(t, []int{1, 2, 3, 4, 5}, 4)
	Contains(t, []string{"foo", "bar", "baz", "qux"}, "baz")

	c1 := make(chan string)
	Contains(t, []chan string{c1}, c1)
}

func TestRegex(t *testing.T) {
	Match(t, "Mary had a little lamb.", "^Mary")
	Match(t, "Mary had a little lamb.", ".*little")
	Match(t, "Mary had a little lamb.", "Mary.*lamb\\.")
}

func TestApprox(t *testing.T) {
	Approx(t, 0.0, -0.0, 0.00001)
	Approx(t, -0.0, 0.0, 0.00001)

	Approx(t, float32(0.0), float32(-0.0), 0.00001)
	Approx(t, float32(-0.0), float32(0.0), 0.00001)

	Approx(t, 2.0, 1.99, 0.01)
	Approx(t, 1.99, 2.0, 0.01)

	Approx(t, float32(2.0), float32(1.99), 0.01)
	Approx(t, float32(1.99), float32(2.0), 0.01)
}

func TestError(t *testing.T) {
	Error(t, errors.New("an error occured"), "an error occured")
}

func TestNil(t *testing.T) {
	Nil(t, nil)

	var foo *int
	Nil(t, foo)

	var bar map[string]int
	Nil(t, bar)

	var baz []int
	Nil(t, baz)

	var qux chan int
	Nil(t, qux)

	var quux func()
	Nil(t, quux)

	var corge any
	Nil(t, corge)

	grault := unsafe.Pointer(uintptr(0))
	Nil(t, grault)
}

func TestNotNil(t *testing.T) {
	foo := 123
	NotNil(t, &foo)

	bar := make(map[string]int)
	NotNil(t, bar)

	baz := []int{1, 2, 3, 4, 5}
	NotNil(t, baz)

	qux := make(chan int)
	NotNil(t, qux)

	quux := func() {}
	NotNil(t, quux)

	corge := "corge"
	NotNil(t, any(corge))

	grault := unsafe.Pointer(uintptr(1337))
	NotNil(t, grault)
}
