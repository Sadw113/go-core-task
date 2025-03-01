package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var originalSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestSliceExample(t *testing.T) {
	resSlice := sliceExample(originalSlice)
	expectedAnswer := []int{2, 4, 6, 8, 10}

	require.Equal(t, expectedAnswer, resSlice)
}

func TestAddElements(t *testing.T) {
	resSlice := addElements(originalSlice, 120)
	expectedAnswer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 120}

	require.Equal(t, expectedAnswer, resSlice)
}

func TestRemoveElement(t *testing.T) {
	resSlice := removeElement(originalSlice, 1)
	expectedAnswer := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, expectedAnswer, resSlice)
}
