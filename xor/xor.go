package xor

import (
	"unsafe"
)

func XorEncodeBytes(msg, key []byte) *[]byte {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, ml)

	for i := 0; i < ml; i++ {

		pwd[i] = key[i%kl] ^ msg[i]

	}

	return &pwd

}

func XorDecodeBytes(msg, key []byte) *[]byte {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, ml)

	for i := 0; i < ml; i++ {

		pwd[i] = key[i%kl] ^ msg[i]

	}

	return &pwd

}

func XorEncodeStr(msg, key string) *string {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, ml)

	for i := 0; i < ml; i++ {

		pwd[i] = key[i%kl] ^ msg[i]

	}

	return bytes2str(pwd)

}

func XorDecodeStr(msg, key string) *string {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, ml)

	for i := 0; i < ml; i++ {

		pwd[i] = key[i%kl] ^ msg[i]

	}

	return bytes2str(pwd)

}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}
