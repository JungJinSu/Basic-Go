package basicGo

import "testing"

// Go 가 제공하는 test class 를 사용하여 각 함수의 탐색 시간을 측정 합니다.


var fa int

// 열 탐색 성능 측정
func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++{
		a = ColumnTraverse()
	}

	fa = a

}

// 행 탐색 측정
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++{
		a = RowTraverse()
	}

	fa = a

}

// 리스트 탐색 측정
func BenchmarkLinkedListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++{
		a = LinkedListTraverse()
	}

	fa = a
}