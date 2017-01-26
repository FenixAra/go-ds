package helpers

import (
	"reflect"
	"testing"
)

func AssertEqualNil(v interface{}, t *testing.T) {
	if v == nil {
		return
	} else {
		t.Errorf("Expected to be nil | Actual: %+v", v)
		t.FailNow()
	}
}

func AssertNoError(err error, t *testing.T, errMessage string) {
	if err != nil {
		t.Errorf("%s ERR: %v", errMessage, err)
		t.FailNow()
	}
}

func AssertEqual(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertNotEqual(expected interface{}, actual interface{}, t *testing.T) {
	if expected == actual {
		t.Errorf("Assert Failure: Expected Not: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertDeepEqual(expected interface{}, actual interface{}, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
	}
}

func AssertStringSliceEqual(expected []string, actual []string, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
		t.FailNow()
		return
	}

	m := make(map[string]bool)
	for _, v := range expected {
		m[v] = true
	}

	for _, v := range actual {
		if _, ok := m[v]; !ok {
			t.Errorf("Assert Failure: Expected: %+v | Actual %+v", expected, actual)
			t.FailNow()
			return
		}
	}
}
