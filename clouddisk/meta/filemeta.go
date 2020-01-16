package meta

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// 用于存储所有的FileMeta
var fileMetas map[string]FileMeta

// 初始化fileMetas
func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMetas(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

// 通过sha1获取文件的元信息对象
func GetFileMeta(sha1 string) FileMeta {
	return fileMetas[sha1]
}