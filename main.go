package main

import (
	// "crypto/hmac"
	// "crypto/sha1"
	// "encoding/hex"
	"flag"
	"fmt"
	"github.com/wuciyou/aliyun-oss/gosdk"
	"os"
	// . "github.com/wuciyou/aliyun-oss/gosdk/lang"
	"path/filepath"
	// "strings"
)

var alioss *gosdk.Alioss
var d string // 便利的目录
func main() {
	flag.StringVar(&d, "d", "../", "Usage: aliossdir -d=/path/to/dogo.json")
	flag.Parse()

	// h := hmac.New(sha1.New, []byte("22"))
	// h.Write([]byte("11"))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
	// fmt.Println("end")

	if len(gosdk.OSS_ACCESS_ID) > 0 && len(gosdk.OSS_ACCESS_KEY) > 0 {
		alioss = &gosdk.Alioss{}
		alioss.Init(gosdk.OSS_ACCESS_ID, gosdk.OSS_ACCESS_KEY)
		alioss.Set_enable_domain_style(true)
	}
	fileslist()
	// List_object()
	// Upload_file_by_file()
	// Create_object_dir()
	return
}

func Create_object_dir() {
	alioss.Create_object_dir("33mimg", "www/fwwfew/wuciyou")
}

func Upload_file_by_file() {
	alioss.Upload_file_by_file("33mimg", "data/tets/xiaoxiaorenwu吴赐有", "/Users/ayou/go/src/github.com/wuciyou/aliyun-oss/gosdk/lang/zh.inc.go")
}

func List_object() {
	options := map[string]string{
		"delimiter": "/",
		"prefix":    "",
		"max-keys":  "10",
		"marker":    "ayou",
	}
	alioss.List_object("33mimg", options)
}

func fileslist() {

	// dirstr := filepath.Dir("/Users/ayou/go/src/github.com/wuciyou")
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// object := strings.Replace(path, "/Users/ayou/", "", -1)
			fmt.Println(info.Mode().IsRegular())
			object := path
			if len(d) > 0 {
				object = d + path
			}
			fmt.Println(object)
			// fmt.Printf("文件：%s\n", path)
			// alioss.Upload_file_by_file("33mimg", object, path)
		}
		return nil
	})
}
