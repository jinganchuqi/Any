package any

import (
	"log"
	"strings"
	"path/filepath"
	"os"
)

/**
	获取错误
 */
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/**
	获取根目录
 */
func getRootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	return strings.Replace(dir, "\\", "/", -1)
}

/**
  	拼装数据
 */

func newData(data interface{})  {

}