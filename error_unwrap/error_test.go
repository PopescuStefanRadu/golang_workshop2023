package error_unwrap

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestErrorWrap(t *testing.T) {
	_, err := strconv.Atoi("ananqjwe")
	err = fmt.Errorf("add local context here, caused by: %w", err)
	require.ErrorContains(t, err, "add local contasade")
}
