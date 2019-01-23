package basicGo

import "fmt"


// 목적 : 데이터 구조(array, list)에 따른 탐색을 테스트 함으로써 캐싱성능을 파악할 수 있습니다.
// go의 3가지 자료구조 : Map, Array, Slice


// 2k by 2k 매트릭 사이즈
const (
	rows = 2 * 1024
	cols = 2 * 1024
)

var matrix [rows][cols] byte


// 데이터 구조
type data struct {
	v byte		// 값
	p *data 	// 포인터
}

var list *data


// 초기화 함수 : config 파일 등 서비스 구동 전에 필요한 작업을 준비한다.
func init(){
	var last *data

	// linked list 구조 생성. 데이터는 같은 숫자.

	for row :=0 ; row < rows ; row++{
		for col :=0 ; col < cols ; col++{

			//데이터 노드 생성 후 링크 걸기
			var d data
			if list == nil {
				list = &d
			}

			if last != nil {
				last.p = &d
			}
			last = &d

			// 짝수에 대한 byte 값을 할당
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	var ctr int
	d := list
	for d != nil{
		ctr++
		d = d.p  // 다음 데이터 주소
	}

	// Count linked list
	fmt.Println("Linked list 개수 : ", ctr)
	fmt.Println("Matrix 사이즈 : ", rows*cols)
}

//LinkedListTraverse : 리스트 단위 탐색
func LinkedListTraverse() int {
	var ctr int

	data := list
	for data != nil {
		if data.v == 0xFF{
			ctr++
		}
		data = data.p
	}

	return ctr
}


//ColumnTraverse : 열 단위 탐색
func ColumnTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row :=0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

//RowTraverse : 행 단위 탐색
func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col :=0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}