package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(2), answer)
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(0), answer)
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(1), answer)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	assert.Nil(t, err)
	require.Equal(t, int32(1), answer)
}

func TestDivisionByZero(t *testing.T) {
	arith := NewAdapter()
	_, err := arith.Division(1, 0)
	assert.NotNil(t, err)
}
