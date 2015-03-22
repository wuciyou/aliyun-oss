package gosdk

const (
	/*%******************************************************************************************%*/
	// ANT

	/**
	 * OSS服务地址
	 */
	DEFAULT_OSS_HOST = "oss.aliyuncs.com"
	// DEFAULT_OSS_HOST = "10.230.201.90"
	/**
	 * 软件名称
	 */
	NAME = "oss-sdk-go"

	/**
	 * OSS软件Build ID
	 */
	BUILD = "201210121010245"

	/**
	 * 版本号
	 */
	VERSION = "1.1.6"

	/**
	 * 作者
	 */
	AUTHOR = "898060380@qq.com"

	/*%******************************************************************************************%*/
	//OSS 内部常量

	OSS_BUCKET                 = "bucket"
	OSS_OBJECT                 = "object"
	OSS_HEADERS                = "headers"
	OSS_METHOD                 = "method"
	OSS_QUERY                  = "query"
	OSS_BASENAME               = "basename"
	OSS_MAX_KEYS               = "max-keys"
	OSS_UPLOAD_ID              = "uploadId"
	OSS_MAX_KEYS_VALUE         = "100"
	OSS_MAX_OBJECT_GROUP_VALUE = "1000"
	OSS_FILE_SLICE_SIZE        = "8192"
	OSS_PREFIX                 = "prefix"
	OSS_DELIMITER              = "delimiter"
	OSS_MARKER                 = "marker"
	OSS_CONTENT_MD5            = "Content-Md5"
	OSS_CONTENT_TYPE           = "Content-Type"
	OSS_CONTENT_LENGTH         = "Content-Length"
	OSS_IF_MODIFIED_SINCE      = "If-Modified-Since"
	OSS_IF_UNMODIFIED_SINCE    = "If-Unmodified-Since"
	OSS_IF_MATCH               = "If-Match"
	OSS_IF_NONE_MATCH          = "If-None-Match"
	OSS_CACHE_CONTROL          = "Cache-Control"
	OSS_EXPIRES                = "Expires"
	OSS_PREAUTH                = "preauth"
	OSS_CONTENT_COING          = "Content-Coding"
	OSS_CONTENT_DISPOSTION     = "Content-Disposition"
	OSS_RANGE                  = "Range"
	OS_CONTENT_RANGE           = "Content-Range"
	OSS_CONTENT                = "content"
	OSS_BODY                   = "body"
	OSS_LENGTH                 = "length"
	OSS_HOST                   = "Host"
	OSS_DATE                   = "Date"
	OSS_AUTHORIZATION          = "Authorization"
	OSS_FILE_DOWNLOAD          = "fileDownload"
	OSS_FILE_UPLOAD            = "fileUpload"
	OSS_PART_SIZE              = "partSize"
	OSS_SEEK_TO                = "seekTo"
	OSS_SIZE                   = "size"
	OSS_QUERY_STRING           = "query_string"
	OSS_SUB_RESOURCE           = "sub_resource"
	OSS_DEFAULT_PREFIX         = "x-oss-"

	/*%******************************************************************************************%*/
	//私有URL变量

	OSS_URL_ACCESS_KEY_ID = "OSSAccessKeyId"
	OSS_URL_EXPIRES       = "Expires"
	OSS_URL_SIGNATURE     = "Signature"

	/*%******************************************************************************************%*/
	//HTTP方法

	OSS_HTTP_GET     = "GET"
	OSS_HTTP_PUT     = "PUT"
	OSS_HTTP_HEAD    = "HEAD"
	OSS_HTTP_POST    = "POST"
	OSS_HTTP_DELETE  = "DELETE"
	OSS_HTTP_OPTIONS = "OPTIONS"

	/*%******************************************************************************************%*/
	//其他常量

	//x-oss
	OSS_ACL = "x-oss-acl"

	//OBJECT GROUP
	OSS_OBJECT_GROUP = "x-oss-file-group"

	//Multi Part
	OSS_MULTI_PART = "uploads"

	//Multi Delete
	OSS_MULTI_DELETE = "delete"

	//OBJECT COPY SOURCE
	OSS_OBJECT_COPY_SOURCE = "x-oss-copy-source"

	//private,only owner
	OSS_ACL_TYPE_PRIVATE = "private"

	//public reand
	OSS_ACL_TYPE_PUBLIC_READ = "public-read"

	//public read write
	OSS_ACL_TYPE_PUBLIC_READ_WRITE = "public-read-write"

	/*%******************************************************************************************%*/

)
