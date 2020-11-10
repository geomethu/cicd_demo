package addition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T){
	sum:=Addition(2.0, 5.0)
	assert.Equal(t, 7.0, sum, "Sum not the same")
}