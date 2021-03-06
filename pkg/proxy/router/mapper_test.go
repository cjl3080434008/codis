package router

import "testing"

func TestMapKey2Slot(t *testing.T) {
	index := mapKey2Slot([]byte("xxx"))
	table := []string{"123{xxx}abc", "{xxx}aa", "x{xxx}"}
	for _, v := range table {
		if index != mapKey2Slot([]byte(v)) {
			t.Error("not match", v)
		}
	}
}
