package sparsesets

import (
	"testing"
)

func TestSet_Insert_error(t *testing.T) {
	ss := New(5)

	if gotErr := ss.Insert(5); gotErr == nil {
		t.Errorf("Insert(): want error, got nil")
	}
	if gotErr := ss.Insert(3); gotErr != nil {
		t.Errorf("Insert(): want no error, got %s", gotErr)
	}
}

func TestSet_Remove_error(t *testing.T) {
	ss := New(5)

	if gotErr := ss.Remove(5); gotErr == nil {
		t.Errorf("Remove(): want error, got nil")
	}
	if gotErr := ss.Remove(3); gotErr != nil {
		t.Errorf("Remove(): want no error, got %s", gotErr)
	}
}

func TestSet_Size(t *testing.T) {
	ss := New(5)

	ss.Insert(0)
	ss.Insert(1)
	if got := ss.Size(); got != 2 {
		t.Errorf("Size(): want 2, got %d", got)
	}

	ss.Insert(1)
	ss.Insert(2)
	if got := ss.Size(); got != 3 {
		t.Errorf("Size(): want 3, got %d", got)
	}
}

func TestSet_Size_empty(t *testing.T) {
	ss := New(0)
	if got := ss.Size(); got != 0 {
		t.Errorf("Size(): want 0, got %d", got)
	}

}
