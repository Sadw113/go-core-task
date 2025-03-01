package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionOfValues(t *testing.T) {
	slice1 := []int{1, 3, 58, 678, 0}
	slice2 := []int{64, 2, 3, 43}

	resSlice, _ := intersectionOfValues(slice1, slice2)
	expectedAnswer := []int{3}

	require.Equal(t, expectedAnswer, resSlice)

	slice1 = []int{1, 2, 3, 4, 5, 6}
	slice2 = []int{7, 8}

	_, exist := intersectionOfValues(slice1, slice2)

	require.Equal(t, false, exist)
}
