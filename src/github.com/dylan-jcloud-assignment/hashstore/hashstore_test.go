package hashstore

import "testing"

func TestSimpleKVHashStore_GetHash(t *testing.T) {
	s := SimpleKVHashStore {}
	s.m = make(map[uint64]string)
	s.m[1] = "hashedpassword1"

	assertEquals("hashedpassword1", s.GetHash(1), t)
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Assert equals fail: got %v, expected %v", actual, expected)
	}
}
