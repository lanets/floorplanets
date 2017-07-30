package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lanets/floorplanets/backend/cmd/floorplanets/internal/cli"
)

func TestHelp(t *testing.T) {
	var writer bytes.Buffer

	cli.Run([]string{"--help"}, &writer)
	assert.True(t, strings.Contains(writer.String(), "Usage"))

	writer.Reset()

	cli.Run([]string{"-h"}, &writer)
	assert.True(t, strings.Contains(writer.String(), "Usage"))
}
