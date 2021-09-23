package splitstring

import (
	"reflect"
	"testing"
)

// 切割

func Test1Split(t *testing.T) {
	got := Split("babababa", "b")
	want := []string{"", "a", "a", "a", "a"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got:%v", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v but got:%v", want, got)
	}
}
func Test3Split(t *testing.T) {
	got := Split("abcdef", "bc")
	want := []string{"a", "def"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want:%v but got:%v", want, got)
	}
}
