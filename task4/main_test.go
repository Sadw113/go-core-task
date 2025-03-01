package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionOfValues(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "eagle", "fix", "gno1"}
	slice2 := []string{"banana", "date", "fig", "eagle"}

	resSlice := intersectionOfValues(slice1, slice2)
	expectedAnswer := []string{"banana", "date", "eagle"}

	require.Equal(t, expectedAnswer, resSlice)
}
