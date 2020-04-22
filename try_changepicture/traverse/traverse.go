package traverse

import (
	"os"
	"path/filepath"
	"fmt"
)

//いいいいいいいいい

func Change(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == ".go"{
		fmt.Println(path)
		//goをjsに変える処理
	}
	return nil
}
