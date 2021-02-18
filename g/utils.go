package g

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// 密码加密
func EncryptPassword(password string) string {
	// md5 加密
	md5 := md5.New()
	md5.Write([]byte(password))
	md5Data := md5.Sum([]byte(nil))
	md5Str := hex.EncodeToString(md5Data)
	// 加盐
	md5Str += "Geekccc"
	// sha1 加密
	sha1 := sha1.New()
	sha1.Write([]byte(md5Str))
	sha1Data := sha1.Sum([]byte(nil))
	return hex.EncodeToString(sha1Data)
}
