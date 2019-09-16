package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/xml"
	"hash"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// X is a convenient alias for a map[string]interface{}.
type X map[string]interface{}

// CDATA XML CDATA section which is defined as blocks of text that are not parsed by the parser, but are otherwise recognized as markup.
type CDATA string

// MarshalXML encodes the receiver as zero or more XML elements.
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// NonceStr 随机字符串
func NonceStr() string {
	now := time.Now()

	return strconv.FormatInt(now.Unix()+int64(now.Nanosecond()), 10)
}

// MD5 calculate the md5 hash of a string.
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 calculate the sha1 hash of a string.
func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// Hash Generate a hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func Hash(t, s string) string {
	var h hash.Hash

	switch strings.ToUpper(t) {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	case "SHA224":
		h = sha256.New224()
	case "SHA256":
		h = sha256.New()
	case "SHA384":
		h = sha512.New384()
	case "SHA512":
		h = sha512.New()
	default:
		return s
	}

	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// HMAC Generate a keyed hash value, expects: MD5, SHA1, SHA224, SHA256, SHA384, SHA512.
func HMAC(t, s, key string) string {
	var mac hash.Hash

	switch strings.ToUpper(t) {
	case "MD5":
		mac = hmac.New(md5.New, []byte(key))
	case "SHA1":
		mac = hmac.New(sha1.New, []byte(key))
	case "SHA224":
		mac = hmac.New(sha256.New224, []byte(key))
	case "SHA256":
		mac = hmac.New(sha256.New, []byte(key))
	case "SHA384":
		mac = hmac.New(sha512.New384, []byte(key))
	case "SHA512":
		mac = hmac.New(sha512.New, []byte(key))
	default:
		return s
	}

	mac.Write([]byte(s))

	return hex.EncodeToString(mac.Sum(nil))
}

// EncodeUint32ToBytes 把整数 uint32 格式化成 4 字节的网络字节序
func EncodeUint32ToBytes(i uint32) []byte {
	b := make([]byte, 4)

	b[0] = byte(i >> 24)
	b[1] = byte(i >> 16)
	b[2] = byte(i >> 8)
	b[3] = byte(i)

	return b
}

// DecodeBytesToUint32 从 4 字节的网络字节序里解析出整数 uint32
func DecodeBytesToUint32(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}

	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	salt := make([]string, 0, n)

	pattern := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	l := len(pattern)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		p := rand.Intn(l)
		salt = append(salt, string(pattern[p]))
	}

	return strings.Join(salt, "")
}
