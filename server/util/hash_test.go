package util_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.nhnent.com/godo/cfo/server/util"
)

func TestHashPassword(t *testing.T) {
	input := "123"
	expected := true

	// pwd, err := util.HashPassword(input)
	// fmt.Println(pwd, err)
	hashedPassword := "$2a$14$bf8U7C6OQvmUZX.ciNSXJOlBOgyP91p9mOBn4B8udibgzbSt0KF0."

	actual := util.CheckPasswordHash(input, hashedPassword)
	fmt.Println(actual)

	assert.Equal(t, expected, actual)
}
