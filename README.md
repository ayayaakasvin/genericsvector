# Vector Package

The `Vector` package provides a generic vector data structure for Go. It supports a variety of operations including appending, removing, inserting, and accessing elements. This package is designed to be flexible and easy to use, making it suitable for a wide range of applications.

## Installation

To install the `Vector` package, use the following `go get` command:

```bash
go get github.com/ayayaakasvin/genericsvector
```

## Usage
Here’s a detailed example of how to use the Vector type in your Go program:

## Basic Operations

```go
package main

import (
    "fmt"
    "github.com/yourusername/yourrepository/lecture"
)

func main() {
    // Create a new vector of integers
    var v lecture.Vector[int]
    
    // Append values to the vector
    v.Append(1)
    v.Append(2)
    v.Append(3)
    fmt.Println("After appending: ", v) // Output: After appending: [1 2 3]
    
    // Get the size of the vector
    fmt.Println("Size:", v.Size()) // Output: Size: 3
    
    // Access the first element
    front, err := v.Front()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Front:", front) // Output: Front: 1
    }
    
    // Access the last element
    back, err := v.Back()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Back:", back) // Output: Back: 3
    }
    
    // Remove an element
    err = v.Remove(2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After removal: ", v) // Output: After removal: [1 3]
    }
    
    // Insert values at a specific index
    err = v.Insert(1, 99, 100)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After insertion: ", v) // Output: After insertion: [1 99 100 3]
    }
    
    // Set a value at a specific index
    err = v.Set(2, 200)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("After setting value: ", v) // Output: After setting value: [1 99 200 3]
    }
    
    // Get an element at a specific index
    element, err := v.At(2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Element at index 2:", element) // Output: Element at index 2: 200
    }
    
    // Clear the vector
    v.Clear()
    fmt.Println("Size after clear:", v.Size()) // Output: Size after clear: 0
}
```

## API Reference
`type Vector[T comparable]`
A generic vector type that holds elements of type `T`.

### Methods
* `Append(value T)`: Adds a single element to the end of the vector.
```go
func (V *Vector[T]) Append(value T)
```

* `Size() int`: Returns the number of elements currently in the vector.
```go
func (V *Vector[T]) Size() int
```

* `Pop() (T, error)`: Removes and returns the last element of the vector. Returns an error if the vector is empty.
```go
func (V *Vector[T]) Pop() (T, error)
```

* `Capacity() int`: Returns the current capacity of the vector.
```go
func (V *Vector[T]) Capacity() int
```

* `At(index int) (T, error)`: Returns the element at the specified index. Returns an error if the index is out of bounds.
```go
func (V *Vector[T]) At(index int) (T, error)
```

* `Find(value T) (int, error)`: Finds and returns the index of the specified value. Returns an error if the value is not found.
```go
func (V *Vector[T]) Find(value T) (int, error)
```

* `Back() (T, error)`: Returns the last element of the vector. Returns an error if the vector is empty.
```go
func (V *Vector[T]) Back() (T, error)
```

* `Clear()`: Clears all elements from the vector.
```go
func (V *Vector[T]) Clear()
```

* `Empty() bool`: Checks if the vector is empty.
```go
func (V *Vector[T]) Empty() bool
```

* `Remove(value T) error`: Removes the first occurrence of the specified value. Returns an error if the vector is empty or the value is not found.
```go
func (V *Vector[T]) Remove(value T) error
```

* `Set(index int, value T) error`: Sets the value at the specified index. Returns an error if the index is out of bounds.
```go
func (V *Vector[T]) Set(index int, value T) error
```

* `Insert(index int, values ...T) error`: Inserts one or more values at the specified index. Returns an error if the index is out of bounds.
```go
func (V *Vector[T]) Insert(index int, values ...T) error
```

For any questions or feedback, please reach out to me at kozhamseitov06@gmail.com.
Feel free to customize the example usage and API details according to your specific needs and functionalities.