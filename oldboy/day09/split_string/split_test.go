package splitstring

import (
	"reflect"
	"testing"
)

// 切割

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		{str: "babababa", sep: "b", want: []string{"", "a", "a", "a", "a"}},
// 		{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
// 		{str: "abcdef", sep: "bc", want: []string{"a", "def"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Errorf("want:%v but got:%v", tc.want, got)
// 		}
// 	}
// }

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case 1": {str: "babababa", sep: "b", want: []string{"", "a", "a", "a", "a"}},
		"case 2": {str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case 3": {str: "abcdef", sep: "bc", want: []string{"a", "def"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want:%v but got:%v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("abc", "b")
	}
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib30(b *testing.B) {
	benchmarkFib(b, 30)
}
