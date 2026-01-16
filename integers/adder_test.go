package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	// Test item:
	// given (preconditions):
	expected := 4

	// when (action):
	got := Add(2, 2)

	// then (expected result):
	if expected != got {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}

// テストモジュール内にサンプルコードを特定のフォーマットで書くと、テストが実行される。
// このサンプルコードは godoc にも反映される。
func ExampleAdd() {
	// "Output: 6" というコメントを削除すると、サンプル関数は実行されないことに注意。
	// 関数はコンパイルされるが、実行ない。
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
