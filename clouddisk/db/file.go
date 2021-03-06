package db

import (
	mydb "clouddisk/db/mysql"
	"fmt"
)

func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_file(`file_sha1`, `file_name`,`file_size`, `file_addr`, `status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Println("file with hash has upload before")
		}
	}
	return true
}
