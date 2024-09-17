package genericsvector

import (
	"fmt"
	"testing"
)

// TestAppend tests the Append method of the Vector type.
func TestAppend(t *testing.T) {
	var v Vector[int]
	v.Append(1)
	v.Append(2)
	v.Append(3)

	if v.Size() != 3 {
		t.Errorf("expected size 3, got %d", v.Size())
	}
	if v[0] != 1 || v[1] != 2 || v[2] != 3 {
		t.Errorf("unexpected vector contents: %v", v)
	}
}

// TestPop tests the Pop method of the Vector type.
func TestPop(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	value, err := v.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != 3 {
		t.Errorf("expected value 3, got %d", value)
	}
	if v.Size() != 2 {
		t.Errorf("expected size 2, got %d", v.Size())
	}
}

// TestPopEmpty tests the Pop method when the vector is empty.
func TestPopEmpty(t *testing.T) {
	var v Vector[int]

	_, err := v.Pop()
	if err == nil || err.Error() != emptyVectorError {
		t.Errorf("expected error %v, got %v", emptyVectorError, err)
	}
}

// TestAt tests the At method of the Vector type.
func TestAt(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	value, err := v.At(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != 2 {
		t.Errorf("expected value 2, got %d", value)
	}
}

// TestAtOutOfBounds tests the At method with an out-of-bounds index.
func TestAtOutOfBounds(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	_, err := v.At(3)
	if err == nil || err.Error() != outOfBoundaryError {
		t.Errorf("expected error %v, got %v", outOfBoundaryError, err)
	}
}

// TestFind tests the Find method of the Vector type.
func TestFind(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	index, err := v.Find(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if index != 1 {
		t.Errorf("expected index 1, got %d", index)
	}
}

// TestFindNotFound tests the Find method when the value is not found.
func TestFindNotFound(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	_, err := v.Find(4)
	if err == nil || err.Error() != fmt.Sprintf(valueNotFoundError, 4) {
		t.Errorf("expected error %v, got %v", fmt.Sprintf(valueNotFoundError, 4), err)
	}
}

// TestFront tests the Front method of the Vector type.
func TestFront(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	value, err := v.Front()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != 1 {
		t.Errorf("expected value 1, got %d", value)
	}
}

// TestFrontEmpty tests the Front method when the vector is empty.
func TestFrontEmpty(t *testing.T) {
	var v Vector[int]

	_, err := v.Front()
	if err == nil || err.Error() != emptyVectorError {
		t.Errorf("expected error %v, got %v", emptyVectorError, err)
	}
}

// TestBack tests the Back method of the Vector type.
func TestBack(t *testing.T) {
	v := Vector[int]{1, 2, 3}

	value, err := v.Back()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != 3 {
		t.Errorf("expected value 3, got %d", value)
	}
}

// TestBackEmpty tests the Back method when the vector is empty.
func TestBackEmpty(t *testing.T) {
	var v Vector[int]

	_, err := v.Back()
	if err == nil || err.Error() != emptyVectorError {
		t.Errorf("expected error %v, got %v", emptyVectorError, err)
	}
}

// TestClear tests the Clear method of the Vector type.
func TestClear(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	v.Clear()
	if v.Size() != 0 {
		t.Errorf("expected size 0, got %d", v.Size())
	}
}

// TestEmpty tests the Empty method of the Vector type.
func TestEmpty(t *testing.T) {
	v := Vector[int]{}

	if !v.Empty() {
		t.Errorf("expected vector to be empty")
	}

	v.Append(1)
	if v.Empty() {
		t.Errorf("expected vector to be non-empty")
	}
}

// TestRemove tests the Remove method of the Vector type.
func TestRemove(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Remove(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v.Size() != 2 || v[0] != 1 || v[1] != 3 {
		t.Errorf("unexpected vector contents: %v", v)
	}
}

// TestRemoveNotFound tests the Remove method when the value is not found.
func TestRemoveNotFound(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Remove(4)
	if err == nil || err.Error() != fmt.Sprintf(valueNotFoundError, 4) {
		t.Errorf("expected error %v, got %v", fmt.Sprintf(valueNotFoundError, 4), err)
	}
}

// TestSet tests the Set method of the Vector type.
func TestSet(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Set(1, 99)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v[1] != 99 {
		t.Errorf("expected value 99 at index 1, got %d", v[1])
	}
}

// TestSetOutOfBounds tests the Set method with an out-of-bounds index.
func TestSetOutOfBounds(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Set(3, 99)
	if err == nil || err.Error() != outOfBoundaryError {
		t.Errorf("expected error %v, got %v", outOfBoundaryError, err)
	}
}

// TestInsert tests the Insert method of the Vector type.
func TestInsert(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Insert(1, 99, 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if v.Size() != 5 || v[1] != 99 || v[2] != 100 {
		t.Errorf("unexpected vector contents: %v", v)
	}
}

// TestInsertOutOfBounds tests the Insert method with an out-of-bounds index.
func TestInsertOutOfBounds(t *testing.T) {
	v := Vector[int]{1, 2, 3}
	err := v.Insert(4, 99)
	if err == nil || err.Error() != outOfBoundaryError {
		t.Errorf("expected error %v, got %v", outOfBoundaryError, err)
	}
}
