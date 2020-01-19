package meta

import mydb "clouddisk/db"

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

// UpdateFileMetaDB 新增到mysql中
func UpdateFileMetaDB(meta FileMeta) bool {
	return mydb.OnFileUploadFinished(meta.FileSha1, meta.FileName, meta.FileSize, meta.Location)
}

// 通过sha1获取文件的元信息对象
func GetFileMeta(sha1 string) FileMeta {
	return fileMetas[sha1]
}

func DeleteFileMeta(sha1 string) {
	delete(fileMetas, sha1)
}
