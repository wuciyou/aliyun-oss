package gosdk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	. "github.com/wuciyou/aliyun-oss/gosdk/lang"
	"github.com/wuciyou/aliyun-oss/gosdk/requestcore"
	"net/url"
	"os"
	"regexp"
	// "sort"
	"io/ioutil"
	"net/http"
	"time"
)

func Sayhello(test string) {
	fmt.Println(OSS_OPTIONS_MUST_BE_ARRAY)
	fmt.Println(AUTHOR)
}

type Alioss struct {
	access_id           string
	access_key          string
	hostname            string
	body                interface{}
	request             *requestcore.BeegoHttpRequest
	vhost               string
	request_url         string
	use_ssl             bool
	set_debug_mode      bool
	enable_domain_style bool
}

type Header map[string]string

func (this *Alioss) Init(access_id, access_key string) {
	this.access_id = access_id
	this.access_key = access_key
	this.hostname = DEFAULT_OSS_HOST
}

func NewOss(access_id, access_key string) *Alioss {
	a := &Alioss{}
	a.Init(access_id, access_key)
	return a
}

/*%******************************************************************************************************%*/
//属性
/**
 * 设置debug模式
 * @param boolean $debug_mode (Optional)
 * @author 898060380@qq.com
 * @since 2014-12-20
 * @return void
 */
func (this *Alioss) Set_debug_mode(debug_mode bool) {
	this.set_debug_mode = debug_mode
}

/**
 * 设置host地址
 * @author 898060380@qq.com
 * @param string hostname host name
 * @param string port string
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) Set_host_name(hostname string, port string) {
	this.hostname = hostname + ":" + port
}

/**
 * 设置vhost地址
 * @author 898060380@qq.com
 * @param string $vhost vhost
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) Set_vhost(vhost string) {
	this.vhost = vhost
}

/**
 * 设置路径形式，如果为true,则启用三级域名，如bucket.oss.aliyuncs.com
 * @author 898060380@qq.com
 * @param boolean $enable_domain_style
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) Set_enable_domain_style(enable_domain_style bool) {
	this.enable_domain_style = enable_domain_style
}

/**
 * 设置请求主体，
 * @author 898060380@qq.com
 * @param interface{} data
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) set_body(data interface{}) {
	this.body = data
}

/**
 * 通过文件名加载文件资源上传
 * @author 898060380@qq.com
 * @param string filename
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) upload_by_filename(filepath string) os.FileInfo {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	sta, _ := f.Stat()
	fbody := make([]byte, sta.Size())
	f.Read(fbody)
	this.body = fbody
	defer f.Close()
	return sta
}

/**
 * 通过url加载文件资源上传
 * @author 898060380@qq.com
 * @param string filename
 * @since 2014-12-11
 * @return void
 */
func (this *Alioss) upload_by_url(url string) (int, error) {
	// bodys, err := requestcore.Get(url).Bytes()
	Response, err := requestcore.Get(url).Response()
	if err != nil {
		return 0, err
	}
	if Response.StatusCode != 200 {
		return 0, errors.New("请求异常")
	}
	bodys, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		return 0, err
	}
	this.body = bodys
	return len(bodys), err
}

/*%******************************************************************************************************%*/
//请求

/**
 * Authorization
 * @param array $options (Required)
 * @throws OSS_Exception
 * @author 898060380@qq.com
 * @since 2014-12-11
 */
func (this *Alioss) auth(options map[string]string, h Header) {
	// msg := "---LOG START---------------------------------------------------------------------------\n"
	//验证Bucket,list_bucket时不需要验证
	// fmt.Println(this.validate_object("334323"))
	// 定义scheme
	scheme := "http://"
	if this.use_ssl {
		scheme = "https://"
	}
	var hostname string
	if bucket, ok := options[OSS_BUCKET]; ok {
		if len(bucket) > 0 {
			hostname = this.hostname + "/" + bucket
		} else {
			hostname = this.hostname
		}
	}
	if this.enable_domain_style {
		if len(this.vhost) > 0 {
			hostname = this.vhost
		} else if bucket, ok := options[OSS_BUCKET]; ok {
			hostname = bucket + "." + this.hostname
		} else {
			hostname = this.vhost
		}
	}

	// object url编码
	var signable_resource string
	// var query_string string
	var string_to_sign string
	header := make(map[string]string)
	header[OSS_CONTENT_MD5] = ""
	// 合并header
	for k, v := range h {
		header[k] = v
	}

	if options[OSS_METHOD] == OSS_HTTP_PUT || options[OSS_METHOD] == OSS_HTTP_POST {
		// this.request.Body(this.body)
		header[OSS_CONTENT_TYPE] = "application/octet-stream"
	}

	if content_type, ok := options[OSS_CONTENT_TYPE]; ok {
		header[OSS_CONTENT_TYPE] = content_type
	} else {
		header[OSS_CONTENT_TYPE] = "application/x-www-form-urlencoded"
	}
	times := time.Now().Add(-(3600 * 8 * 1000000000))
	datestr := times.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	header[OSS_DATE] = datestr
	if object, ok := options[OSS_OBJECT]; ok && object != "/" {
		signable_resources := "/" + url.QueryEscape(object)
		fmt.Println(signable_resources)
		signable_resource = "/" + object
	}
	// if query_string, ok := options[OSS_QUERY_STRING]; ok && len(query_string) > 0 {
	// 	url.Parse(rawurl)
	// }
	// fmt.Println(signable_resource)
	string_to_sign = options[OSS_METHOD] + "\n" + header[OSS_CONTENT_MD5] + "\n" + header[OSS_CONTENT_TYPE] + "\n" + header[OSS_DATE] + "\n"
	string_to_sign += "/" + options[OSS_BUCKET]
	// fmt.Println(string_to_sign)
	if this.enable_domain_style && options[OSS_BUCKET] != "" && options[OSS_OBJECT] == "/" {
		string_to_sign += "/"
	}
	string_to_sign += signable_resource
	this.request_url = scheme + hostname + signable_resource
	this.request = requestcore.NewBeegoRequest(this.request_url, options[OSS_METHOD])
	if options[OSS_METHOD] == OSS_HTTP_PUT || options[OSS_METHOD] == OSS_HTTP_POST {
		this.request.Body(this.body)
	}
	signature := this.sige(string_to_sign, this.access_key)
	this.request.Headers(header)
	this.request.Header("Authorization", "OSS "+this.access_id+":"+signature)
}

func (this *Alioss) sige(access, key string) string {
	if len(key) <= 0 {
		key = this.access_key
	}
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(access))
	return base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

/**
 * 记录日志
 * @param string $msg (Required)
 * @throws OSS_Exception
 * @author 898060380@qq.com
 * @since 2014-12-20
 * @return void
 */
func (this *Alioss) log(msg string) {
	fmt.Println(msg)
}

/**
 * 校验BUCKET/OBJECT/OBJECT GROUP是否为空
 * @param  string $name (Required)
 * @param  string $errMsg (Required)
 * @throws OSS_Exception
 * @author 898060380@qq.com
 * @since 2014-12-20
 * @return void
 */
func (this *Alioss) is_empty(name string, errMsg string) {
	if len(name) <= 0 {
		panic(errMsg)
	}
}

// 获取object list
func (this *Alioss) List_object(bucket string, options map[string]string) {

	header := make(Header)

	options[OSS_BUCKET] = bucket
	options[OSS_METHOD] = OSS_HTTP_GET
	options[OSS_OBJECT] = "/"

	// 对 Object 名字进行分组的字符
	if delimitr, ok := options[OSS_DELIMITER]; ok {
		header[OSS_DELIMITER] = delimitr
	} else {
		header[OSS_DELIMITER] = "/"
	}

	// 限定返回的object key必须以prefix作为前缀
	if prefix, ok := options[OSS_PREFIX]; ok {
		header[OSS_PREFIX] = prefix
	} else {
		header[OSS_PREFIX] = ""
	}

	// 限定此次返回 object 的最大数,如果不设定,默认为 100
	if max_key, ok := options[OSS_MAX_KEYS]; ok {
		header[OSS_MAX_KEYS] = max_key
	} else {
		header[OSS_MAX_KEYS] = OSS_MAX_KEYS_VALUE
	}
	// 设定结果从 marker 之后按字母排序的第一个开始返回。
	if marker, ok := options[OSS_MARKER]; ok {
		header[OSS_MARKER] = marker
	} else {
		header[OSS_MARKER] = ""
	}
	this.auth(options, header)
	fmt.Println(this.request.String())
	return
	request := requestcore.NewBeegoRequest("http://33mimg.oss.aliyuncs.com", "GET")
	request.Header("Content-Type", "application/x-www-form-urlencoded")
	times := time.Now().Add(-(3600 * 8 * 1000000000))
	datestr := times.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	request.Header("Date", datestr)
	// request.Header("Date", "Thu, 18 Dec 2014 08:40:26 GMT")
	request.Header("delimiter", "/")
	request.Header("Host", "33mimg.oss.aliyuncs.com")
	request.Header("max-keys", "")

	// string_to_sig := "GET application/x-www-form-urlencoded " + datestr + " /33mimg/"
	// string_to_sig = "GET\napplication/x-www-form-urlencoded\nThu, 18 Dec 2014 08:31:46 GMT\n33mimg/"
	// sige := this.sige(string_to_sig, "")
	// fmt.Println(string_to_sig)
	// fmt.Println(sige)
	// return
	sige := this.sige("GET\n\napplication/x-www-form-urlencoded\n"+datestr+"\n/33mimg/", "")
	// 60UiS90R2o8Uk7XKUg9THnQJz1g=

	request.Header("Authorization", "OSS Ofn5oRuOVFgRfTpH:"+sige)
	fmt.Println(request.GetHeaderParam())
	fmt.Println(request.String())
}

func (this *Alioss) Upload_by_url(bucket, object, url string) (*http.Response, error) {
	options := make(map[string]string)
	h := make(Header)
	options[OSS_METHOD] = OSS_HTTP_PUT
	options[OSS_BUCKET] = bucket
	if len(object) > 0 {
		options[OSS_OBJECT] = object
	} else {
		options[OSS_OBJECT] = url
	}
	content, err := this.upload_by_url(url)
	if err != nil {
		return nil, err
	}
	h[OSS_CONTENT_LENGTH] = fmt.Sprintf("%s", content)
	h[OSS_CACHE_CONTROL] = "3600"
	this.auth(options, h)
	return this.request.Response()
}

/**
 * 上传文件，适合比较大的文件
 * @param string $bucket (Required)
 * @param string $object (Required)
 * @param string $file (Required)
 * @param array $options (Optional)
 * @author 898060380@qq.com
 * @since 2014-12-11
 * @return ResponseCore
 */
func (this *Alioss) Upload_file_by_file(bucket, object, file_path string) *requestcore.BeegoHttpRequest {
	options := make(map[string]string)
	h := make(Header)
	fileinfo := this.upload_by_filename(file_path)
	options[OSS_METHOD] = OSS_HTTP_PUT
	options[OSS_BUCKET] = bucket
	if len(object) > 0 {
		options[OSS_OBJECT] = object
	} else {
		options[OSS_OBJECT] = fileinfo.Name()
	}
	h[OSS_CONTENT_LENGTH] = fmt.Sprintf("%d", fileinfo.Size())
	this.auth(options, h)
	return this.request
	// return
	// request := requestcore.NewBeegoRequest("http://33mimg.oss.aliyuncs.com/testupdates.jpg", "PUT")
	// request.Header("Content-Type", "application/octet-stream")
	// times := time.Now().Add(-(3600 * 8 * 1000000000))
	// datestr := times.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	// request.Header("Date", datestr)
	// request.Header("Host", "33mimg.oss.aliyuncs.com")
	// sige := this.sige("PUT\n\napplication/octet-stream\n"+datestr+"\n/33mimg/testupdates.jpg", "")
	// request.Header("Authorization", "OSS Ofn5oRuOVFgRfTpH:"+sige)
	// f, _ := os.Open("/Users/ayou/www/new.msc50.cn/notes/guanggun11/guanggun.png")
	// fsta, _ := f.Stat()
	// request.Header("content-length", fmt.Sprintf("%d", fsta.Size()))
	// filedate := make([]byte, 3942909)
	// f.Read(filedate)
	// request.Body(filedate)
	// response, _ := request.Response()

	// fmt.Println(response.Header)
}

/**
 * 创建目录(目录和文件的区别在于，目录最后增加'/')
 * @param string $bucket
 * @param string $object
 * @param array $options
 * @author 898060380@qq.com
 * @since 2014-12-11
 * @return ResponseCore
 */
func (this *Alioss) Create_object_dir(bucket string, object string) {
	fmt.Println(bucket)
	fmt.Println(object)
	options := make(map[string]string)
	options[OSS_BUCKET] = bucket
	options[OSS_METHOD] = OSS_HTTP_PUT
	options[OSS_OBJECT] = object + "/"
	this.auth(options, nil)
	response, _ := this.request.Response()
	fmt.Println(response.Header)
	fmt.Println(this.request.String())
}

/**
 * 检验$options
 * @param array $options (Optional)
 * @throws OSS_Exception
 * @author 898060380@qq.com
 * @since 2014-12-20
 * @return boolean
 */
func (this *Alioss) validate_options(options map[string]string) error {
	return nil
}

/*%******************************************************************************************************%*/
//工具方法相关

/**
* map排序
 */
type MapSorter []Item
type Item struct {
	Key string
	Val string
}

func NewMapSorter(m map[string]string) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}
func (ms MapSorter) Len() int {
	return len(ms)
}
func (ms MapSorter) Less(i, j int) bool {
	// return ms[i].Val < ms[j].Val // 按值排序
	return ms[i].Key < ms[j].Key // 按键排序
}
func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

/**
* 检验object名称是否合法
* object命名规范:
* 1. 规则长度必须在1-1023字节之间
* 2. 使用UTF-8编码
* @param string $object (Required)
* @author 898060380@qq.com
* @since 2014-12-20
* @return boolean
 */
func (this *Alioss) validate_object(object string) bool {
	regexp := regexp.MustCompile(`/^\d{1,13}$/`)
	return regexp.MatchString(object)
}

/**
* 处理异常的
 */
func (this *Alioss) exception(err string) {
	e := errors.New(err)
	panic(e)
}
