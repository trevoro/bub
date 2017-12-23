package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSanitizeBranchName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "55gf66-sf9-3", Git().sanitizeBranchName("55gf66     sf9#3     "))
	assert.Equal(t, "55gf66-sf9-3", Git().sanitizeBranchName("55gf66 sf9#3"))
}

func TestExtractIssueKey(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "PL-2345", Git().extractIssueKeyFromName("PL-2345-asfsd-asfsf-sffff"))
}
