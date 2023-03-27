package proxy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	s := &SinaProxy{}
	list, err := s.List("ziguangguowei")
	assert.Nil(t, err)
	fmt.Println(list)
}

func TestGet(t *testing.T) {
	s := &SinaProxy{}
	details, err := s.Get("sz002049")
	assert.Nil(t, err)
	fmt.Println(details)
}

func TestChart(t *testing.T) {
	s := &SinaProxy{}
	chart, err := s.getNewChart("SZ002049")
	assert.Nil(t, err)
	fmt.Println(chart)
}
