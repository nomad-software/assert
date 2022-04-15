package assert

import "testing"

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

func TestRegex(t *testing.T) {
	Regex(t, "Mary had a little lamb.", "^Mary")
	Regex(t, "Mary had a little lamb.", ".*little")
	Regex(t, "Mary had a little lamb.", "Mary.*lamb\\.")
}
