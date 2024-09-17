package genericsvector

import "fmt"

// Error messages for various vector operations
const (
	emptyVectorError   = "vector is empty"
	valueNotFoundError = "value %v not found"
	outOfBoundaryError = "index out of bounds"
)

// Vector represents a generic slice with additional methods.
type Vector[T comparable] []T

// Append adds a values to the end of the vector.
func (V *Vector[T]) Append(values ...T) {
	*V = append(*V, values...)
}

// Size returns the number of elements in the vector.
func (V *Vector[T]) Size() int {
	return len(*V)
}

// Pop removes and returns the last element of the vector.
// Returns an error if the vector is empty.
func (V *Vector[T]) Pop() (T, error) {
	if V.Empty() {
		var zeroValue T
		return zeroValue, fmt.Errorf(emptyVectorError)
	}

	sizeOfVector := V.Size()
	temp := (*V)[sizeOfVector-1]
	*V = (*V)[:sizeOfVector-1]
	return temp, nil
}

// Capacity returns the capacity of the vector's underlying slice.
func (V *Vector[T]) Capacity() int {
	return cap(*V)
}

// At returns the element at the specified index and an error if the index is out of bounds.
func (V *Vector[T]) At(index int) (T, error) {
	if index < 0 || index >= V.Size() {
		var zeroValue T
		return zeroValue, fmt.Errorf(outOfBoundaryError)
	}
	return (*V)[index], nil
}

// Find searches for a value in the vector and returns its index.
// Returns an error if the value is not found.
func (V *Vector[T]) Find(value T) (int, error) {
	for index, el := range *V {
		if el == value {
			return index, nil
		}
	}
	return -1, fmt.Errorf(valueNotFoundError, value)
}

// Front returns the first element of the vector and an error if the vector is empty.
func (V *Vector[T]) Front() (T, error) {
	if V.Empty() {
		var zeroValue T
		return zeroValue, fmt.Errorf(emptyVectorError)
	}
	return (*V)[0], nil
}

// Back returns the last element of the vector and an error if the vector is empty.
func (V *Vector[T]) Back() (T, error) {
	if V.Empty() {
		var zeroValue T
		return zeroValue, fmt.Errorf(emptyVectorError)
	}
	return (*V)[V.Size()-1], nil
}

// Clear removes all elements from the vector, making it empty.
func (V *Vector[T]) Clear() {
	*V = (*V)[:0]
}

// Empty checks if the vector has no elements.
func (V *Vector[T]) Empty() bool {
	return V.Size() == 0
}

// Remove deletes the first occurrence of a value from the vector.
// Returns an error if the vector is empty or the value is not found.
func (V *Vector[T]) Remove(value T) error {
	if V.Empty() {
		return fmt.Errorf(emptyVectorError)
	}

	index, err := V.Find(value)
	if err != nil {
		return err
	}

	*V = append((*V)[:index], (*V)[index+1:]...)
	return nil
}

// Set updates the element at the specified index with a new value.
// Returns an error if the index is out of bounds.
func (V *Vector[T]) Set(index int, value T) error {
	if index < 0 || index >= V.Size() {
		return fmt.Errorf(outOfBoundaryError)
	}
	(*V)[index] = value
	return nil
}

// Insert adds one or more values at the specified index in the vector.
// Returns an error if the index is out of bounds.
func (V *Vector[T]) Insert(index int, values ...T) error {
	if index < 0 || index > V.Size() {
		return fmt.Errorf(outOfBoundaryError)
	}

	*V = append((*V)[:index], append(values, (*V)[index:]...)...)
	return nil
}
