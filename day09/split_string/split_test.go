package split_string

import (
	// "fmt"
	"reflect"
	"testing"
)

/*

单元测试
go test命令是一个按照一定约定和组织的测试代码的驱动程序。
在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
*/

// func TestSplit(t *testing.T) {
// 	ret := Split("babcbef", "b")
// 	want := []string{"", "a", "c", "ef"}
// 	// fmt.Printf("%#v\n", ret)
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败
// 		t.Errorf("want:%v but got :%v\n", want, ret)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	ret := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	// fmt.Printf("%#v\n", ret)
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败
// 		t.Errorf("want:%v but got :%v\n", want, ret)
// 	}
// }

// func Test3Split(t *testing.T) {
// 	ret := Split("abcef", "bc")
// 	want := []string{"a", "ef"}
// 	// fmt.Printf("%#v\n", ret)
// 	if !reflect.DeepEqual(ret, want) {
// 		// 测试用例失败
// 		t.Errorf("want:%v but got :%v\n", want, ret)
// 	}
// }

/*
func TestGroupSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcef", "bc", []string{"a", "ef"}},
		testCase{"小鳄鱼爱洗澡", "爱", []string{"小鳄鱼", "洗澡"}},
	}
	for _, testItem := range testGroup {
		got := Split(testItem.str, testItem.sep)
		if !reflect.DeepEqual(got, testItem.want) {
			t.Errorf("want:%v but got :%v\n", testItem.want, got)
		}
	}
}
*/

//子测试
func TestGroupSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case 1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case 2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case 3": testCase{"abcef", "bc", []string{"a", "ef"}},
		"case 4": testCase{"小鳄鱼爱洗澡", "爱", []string{"小鳄鱼", "洗澡"}},
	}

	for name, testItem := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(testItem.str, testItem.sep)
			if !reflect.DeepEqual(got, testItem.want) {
				t.Errorf("want:%#v but got :%#v\n", testItem.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// benchmarkFib 性能比较测试

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib(b *testing.B) {
	benchmarkFib(b, 5)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 20)
}
