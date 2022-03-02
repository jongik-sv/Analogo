package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNormal(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", NORMAL)
	p.Add("A", "B", NORMAL)
	p.Add("B", "A", NORMAL)
	a := p.Get("A")
	assert.Equal(t, 1, len(a))
	assert.Equal(t, "B", a[0])

	b := p.Get("B")
	assert.Equal(t, 1, len(b))
	assert.Equal(t, "A", b[0])
	c := p.Get("C")
	assert.Equal(t, 0, len(c))

}

func TestAddSet(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", SET)
	p.Add("A", "B", SET)
	p.Add("A", "B", SET)
	p.Add("A", "B", SET)
	p.Add("A", "B", SET)
	p.Add("B", "A", SET)
	a := p.Get("A")
	assert.Equal(t, 2, len(a))
	b := p.Get("B")
	assert.Equal(t, 1, len(b))
	assert.Equal(t, "A", b[0])
	c := p.Get("C")
	assert.Equal(t, 0, len(c))

}

func TestAddAndGetFromProperties(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("B", "A", ARRAY)
	a := p.Get("A")
	assert.Equal(t, 2, len(a))
	assert.Equal(t, "A", a[0])
	assert.Equal(t, "B", a[1])

	b := p.Get("B")
	assert.Equal(t, 1, len(b))
	assert.Equal(t, "A", a[0])
	c := p.Get("C")
	assert.Equal(t, 0, len(c))

}

func TestGetJoinStringFromProperties(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("B", "A", ARRAY)
	a := p.GetJoinString("A")
	assert.Equal(t, "A, B", a)
	b := p.GetJoinString("B")
	assert.Equal(t, "A", b)
	c := p.GetJoinString("C")
	assert.Equal(t, "", c)
}

func TestGetFirstFromProperties(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("B", "B", ARRAY)
	a := p.GetFirst("A")
	assert.Equal(t, "A", a)
	b := p.GetFirst("B")
	assert.Equal(t, "B", b)
	c := p.GetFirst("C")
	assert.Equal(t, "", c)

}

func TestGetLastFromProperties(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "X", ARRAY)
	p.Add("B", "B", ARRAY)
	a := p.GetLast("A")
	assert.Equal(t, "X", a)
	b := p.GetLast("B")
	assert.Equal(t, "B", b)
	c := p.GetLast("c")
	assert.Equal(t, "", c)
}

func TestGetSliceThroughSetMethod(t *testing.T) {
	p := NewProperties()
	p.Add("A", "A", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "B", ARRAY)
	p.Add("A", "X", ARRAY)
	p.Add("B", "B", ARRAY)
	a := strings.Join(p.GetSliceThroughSetMethod("A"), ",")
	assert.Equal(t, "A,B,X", a)

}

func TestSplit(t *testing.T) {
	fmt.Println(strings.Split("target1,target2", ","))
	for _, target := range strings.Split("target1,target2", ",") {
		fmt.Println(target)
	}
}
