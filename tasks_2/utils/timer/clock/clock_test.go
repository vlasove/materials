package clock

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClock_HostChecker(t *testing.T) {

	testCases := []struct {
		name    string
		isValid bool
		host    string
	}{
		{
			name:    "base host",
			isValid: true,
			host:    BaseHost,
		},
		{
			name:    "invalid host",
			isValid: false,
			host:    "test-ivalid-host",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			clock := New(test.host, os.Stderr)
			err := clock.hostChecker()
			if test.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestClock_SetCurrent(t *testing.T) {
	buffer := bytes.Buffer{}
	clock := New(BaseHost, &buffer)
	err := clock.setCurrent()
	assert.NoError(t, err)
	assert.Equal(t, len(buffer.String()), 0)
}
