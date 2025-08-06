// Package utils provides utility functions
package utils

import (
	"errors"
	"math"
)

// Operation represents a mathematical operation
type Operation string

// Operation constants
const (
	OperationAdd      Operation = "add"
	OperationSubtract Operation = "subtract"
	OperationMultiply Operation = "multiply"
	OperationDivide   Operation = "divide"
)

// ErrDivisionByZero is returned when attempting to divide by zero
var ErrDivisionByZero = errors.New("division by zero")

// Calculate performs a mathematical operation on two numbers
func Calculate(a, b float64, op Operation) (float64, error) {
	switch op {
	case OperationAdd:
		return a + b, nil
	case OperationSubtract:
		return a - b, nil
	case OperationMultiply:
		return a * b, nil
	case OperationDivide:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

// MathUtils provides advanced mathematical operations
type MathUtils struct {
	precision int
}

// NewMathUtils creates a new MathUtils instance
func NewMathUtils(precision int) *MathUtils {
	return &MathUtils{
		precision: precision,
	}
}

// Round rounds a number to the specified precision
func (m *MathUtils) Round(value float64) float64 {
	multiplier := math.Pow(10, float64(m.precision))
	return math.Round(value*multiplier) / multiplier
}

// Average calculates the average of a slice of numbers
func (m *MathUtils) Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	
	return m.Round(sum / float64(len(numbers)))
}

// Vector2D represents a 2D vector
type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Add adds two vectors
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Magnitude returns the magnitude of the vector
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a normalized version of the vector
func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0}
	}
	return Vector2D{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

// Vector3D represents a 3D vector
type Vector3D[T ~float32 | ~float64] struct {
	X T `json:"x"`
	Y T `json:"y"`
	Z T `json:"z"`
}

// Add adds two 3D vectors
func (v Vector3D[T]) Add(other Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Cross computes the cross product of two vectors
func (v Vector3D[T]) Cross(other Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

// Matrix2x2 represents a 2x2 matrix
type Matrix2x2 [2][2]float64

// Multiply multiplies two 2x2 matrices
func (m Matrix2x2) Multiply(other Matrix2x2) Matrix2x2 {
	return Matrix2x2{
		{
			m[0][0]*other[0][0] + m[0][1]*other[1][0],
			m[0][0]*other[0][1] + m[0][1]*other[1][1],
		},
		{
			m[1][0]*other[0][0] + m[1][1]*other[1][0],
			m[1][0]*other[0][1] + m[1][1]*other[1][1],
		},
	}
}