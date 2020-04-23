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
		{subTestName: "+zero",input: 0, expected: true},
	}

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
//357ページの「c :=c」がわからない
//テストはインプット、アウトプットが予想できる時しか使えない?
//ターミナルでの実行 $go test　<ここはなに？>
//ターミナルでの実行 $go test　<ここはなに？>　の後、coverprofile