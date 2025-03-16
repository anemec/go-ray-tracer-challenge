package geometry

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor_Conversion255(t *testing.T) {
	color := NewColor(1.5, 0, 0)
	sut := color.ToColorString()

	require.Equal(t, "255 0 0", sut)
}

func TestColor_Conversion128(t *testing.T) {
	color := NewColor(0, 0.5, 0)
	sut := color.ToColorString()

	require.Equal(t, "0 128 0", sut)
}

func TestColor_Conversion204(t *testing.T) {
	color := NewColor(1, 0.8, 0.6)
	sut := color.ToColorString()

	require.Equal(t, "255 204 153", sut)
}
