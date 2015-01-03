package security

import (
	"crypto/md5"
	"encoding/hex"
	"planiator/server/conf"
)

// EncryptPassword encrypts password
func EncryptPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password + conf.Salt1))
	encrypted1 := hex.EncodeToString(hasher.Sum(nil))
	hasher.Write([]byte(conf.Salt2 + encrypted1))
	return hex.EncodeToString(hasher.Sum(nil))
}
