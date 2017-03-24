package file

import "testing"

// We use the unique string slice in the backend provider, so we should test that it works
func Test_uniqueStringSlice(t *testing.T) {
	s := uniqueStringSlice{}

	// test adding strings
	s.add("A")
	sl1 := s.slice()
	if len(sl1) != 1 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl1[0] != "A" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test adding further
	s.add("B")
	s.add("C")
	sl2 := s.slice()
	if len(sl2) != 3 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl2[0] != "A" || sl2[1] != "B" || sl2[2] != "C" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test unique values
	s.add("A")
	sl3 := s.slice()
	if len(sl3) != 3 {
		t.Error("uniqueStringSlice improperly added a duplicate value")
	}
}
