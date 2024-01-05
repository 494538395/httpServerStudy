package utlis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)

	testcases := []struct {
		name     string
		arg      any
		expected map[string]interface{}
	}{
		{
			name: "struct to map",
			arg: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{Name: "jerry", Age: 21},
			expected: map[string]interface{}{
				"name": "jerry",
				"age":  21,
			},
		},
	}

	for _, tt := range testcases {
		res := Map(tt.arg)
		assert.Equal(tt.expected, res)
	}

}
