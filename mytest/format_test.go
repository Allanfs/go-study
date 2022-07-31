package mytest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatName(t *testing.T) {
	formated := formatName("Allan", "Silva")
	assert.Equal(t, "Dear Allan Silva", formated)
}

func TestG(t *testing.T) {
	assert.Equal(t, 4, g())
}
