package main

import (
	// "crypto/hmac"
	// "crypto/sha1"
	// "encoding/hex"
	"fmt"
	"github.com/wuciyou/aliyun-oss/gosdk"
	"os"
	// . "github.com/wuciyou/aliyun-oss/gosdk/lang"
)

var alioss *gosdk.Alioss

func main() {
	// h := hmac.New(sha1.New, []byte("22"))
	// h.Write([]byte("11"))
	// fmt.Println(hex.EncodeToString(h.Sum(nil)))
	// fmt.Println("end")
	alioss = &gosdk.Alioss{}
	if len(gosdk.OSS_ACCESS_ID) > 0 && len(gosdk.OSS_ACCESS_KEY) > 0 {
		alioss.Init(gosdk.OSS_ACCESS_ID, gosdk.OSS_ACCESS_KEY)
		alioss.Set_enable_domain_style(true)
	}
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
func files() {
	f, _ := os.Open("/Users/ayou/go/src/github.com/wuciyou/aliyun-oss/dogo.json")
	stat, _ := f.Stat()
	fb := make([]byte, stat.Size())
	f.Read(fb)
	fmt.Println(fb)
	fmt.Println(stat.Size())
}
