package isEven

import (
	"fmt"
	"testing"
)


//サブテストとテーブル駆動テスト
func TestIsEven(t *testing.T){ //偶数の時trueである

	datas := []struct{
		subTestName string; input int; expected bool
	}{
		{subTestName: "+odd",input: 3, expected: false},
		{subTestName: "+even",input: 4, expected: true},
		{subTestName: "-odd",input: -3, expected: false},
		{subTestName: "-even",input: -4, expected: true},
		{subTestName: "zero",input: 0, expected: true},
	}

	//テストは、inputに対してoutputが正しくかえってくるか確かめる
	for _, d := range datas{

		//t.Runの使い方
		//t.Run(サブテスト名, 関数)
		t.Run(d.subTestName, func (t *testing.T){
			if actual := IsEven(d.input); actual == d.expected{
				fmt.Printf("%s を入力した時:成功です\n",d.subTestName)
			}else{
				t.Errorf("want IsEven(%d) = %v, but got %v", d.input, d.expected, actual)
			}
		})
	}

	//t.Run("A",func (t *testing.T){fmt.Println("A")}) //サブテスト例
	//t.Run("b",func (t *testing.T){fmt.Println("B")})
	//t.Run("c",func (t *testing.T){fmt.Println("c")})

}



/*
//テーブル駆動テスト

func TestIsEven(t *testing.T) { //偶数の時trueである
	datas := []struct {
		input    int;
		expected bool
	}{
		{input: 3, expected: false},
		{input: 4, expected: true},
		{input: -3, expected: false},
		{input: -4, expected: true},
		{input: 0, expected: true},
	}

	for _, d := range datas {
		if IsEven(d.input) == d.expected {
			fmt.Printf("入力データが %d の時:テスト成功です\n", d.input)
		} else {
			t.Error("error")
		}
	}
}
*/



//疑問
//+ができない


//memo
//357ページの「c :=c」がわからない →　いらない
//テストはインプット、アウトプットが予想できる時しか使えない?　→　そう
//ターミナルでの実行 $go test　<ここはなに？>　→ $go test ディレクトリ


//$go test -coverprofile
//プロファイルの作成

//$go test -coverprofile=cover.out example
//作成されたプロファイルcover.go.
//exampleはファイルパス

//$ go tool cover -html=cover.out
//$ go tool cover -func=cover.out
//作成されたプロファイル「cover.out」に対して、「go tool cover」

//$ go test -coverprofile=cover.out -covermode=count example
//-covermodeは解析のモード //count,set,automatic