package utils

import (
	"reflect"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := UUID()
	ref := reflect.ValueOf(uuid)
	if ref.Kind() != reflect.String {
		t.Errorf("UUID is not a string")
	}
}
