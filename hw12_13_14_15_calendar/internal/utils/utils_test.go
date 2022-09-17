package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUtils_BeginningOfDay(t *testing.T) {
	t1, _ := time.Parse(dateTimeLayout, "2020-01-01T12:12:12")
	t2, _ := time.Parse(dateTimeLayout, "2020-01-01T00:00:00")
	require.Equal(t, t2, BeginningOfDay(t1))
}

func TestUtils_EndOfDay(t *testing.T) {
	t1, _ := time.Parse(dateTimeLayout, "2020-01-01T12:12:12")
	t2, _ := time.Parse(dateTimeLayout, "2020-01-01T23:59:59")
	require.Equal(t, t2.Add(time.Second-time.Nanosecond), EndOfDay(t1))
}
