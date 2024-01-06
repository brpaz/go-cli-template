package hello_test

import (
	"bytes"
	"testing"

	"github.com/brpaz/go-cli-template/cmd/hello"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloCmd(t *testing.T) {

	t.Parallel()

	t.Run("Prints hello world", func(t *testing.T) {
		helloCmd := hello.NewCommand()
		out := bytes.NewBuffer([]byte{})
		helloCmd.SetOut(out)

		err := helloCmd.RunE(helloCmd, []string{})
		require.NoError(t, err)

		assert.Equal(t, "Hello World\n", out.String())
	})
}
