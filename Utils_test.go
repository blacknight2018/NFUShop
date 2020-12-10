package main

import (
	"NFUShop/Utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	a := Utils.GenerateJWT(123)
	o := Utils.ParseJWT(a)
	assert.Equal(t, o, 123)

	a = Utils.GenerateJWT("123")
	o = Utils.ParseJWT(a)
	assert.Equal(t, o, "123")

}
