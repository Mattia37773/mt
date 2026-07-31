package base

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertFilesEqual(t *testing.T, expectedFile, actualFile string) {
	t.Helper()

	content1, err := os.ReadFile(expectedFile)
	assert.NoError(t, err)

	content2, err := os.ReadFile(actualFile)
	assert.NoError(t, err)

	assert.Equal(t, string(content1), string(content2))
}
