package soa

import (
	"os"
	"reflect"
)

// Manager encapsulates the state and logic.
type Manager struct {
	pageSize int
}

// Allocate creates a new slice s by using a total size of the page size.
// The length and capacity is set to "page size / size of type".
// All values are initialized with zero (0).
func (m *Manager) Allocate(s interface{}) (t interface{}) {
	size := m.pageSize / m.sizeOf(s)
	return reflect.MakeSlice(reflect.TypeOf(s), size, size).Interface()
}

// Grow appends a new page of values to a given slice s
// and save it as a new slice t.
func (m *Manager) Grow(s interface{}) (t interface{}) {
	size := m.pageSize / m.sizeOf(s)
	return reflect.AppendSlice(reflect.ValueOf(s), reflect.MakeSlice(reflect.TypeOf(s), size, size)).Interface()
}

// Pack shrinks the slice s to a total amount of values which are not zero (0).
// The order of the elements remains the same.
func (m *Manager) Pack(s interface{}) (t interface{}) {
	// create an empty slice for the result.
	out := reflect.MakeSlice(reflect.TypeOf(s), 0, 0)
	// now check each element for a value of zero(0).
	slice := reflect.ValueOf(s)
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		// if value is zero (0) then skip.
		if item.Int() == reflect.ValueOf(0).Int() {
			continue
		}
		// add the non-zero value to the result.
		tmp := reflect.MakeSlice(reflect.TypeOf(s), 1, 1)
		tmp.Index(0).Set(item)
		out = reflect.AppendSlice(out, tmp)
	}
	// finally, return the packend slice.
	return out.Interface()
}

// Unpack grows the slice s to the next multiple of page size.
func (m *Manager) Unpack(s interface{}) (t interface{}) {
	multiple := m.pageSize / m.sizeOf(s)
	size := (multiple - reflect.ValueOf(s).Len()) * (reflect.ValueOf(s).Len()/multiple + 1)
	return reflect.AppendSlice(reflect.ValueOf(s), reflect.MakeSlice(reflect.TypeOf(s), size, size)).Interface()
}

// SizeOf calculates the size of type t's values.
func (m *Manager) sizeOf(t interface{}) int {
	return int(reflect.TypeOf(t).Elem().Size())
}

// NewManager creates a new Manager and returns its pointer.
func NewManager() *Manager {
	return &Manager{
		pageSize: os.Getpagesize(),
	}
}

// DefaultManager creates a single instance of Manager at startup.
var DefaultManager = NewManager()
