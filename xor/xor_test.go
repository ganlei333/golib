package xor

import "testing"

func TestXorBytes(t *testing.T) {
	key := []byte("1234567899")
	msg := []byte("asdfghjkl")
	data := XorEncodeBytes(msg, key)
	if string(*data) == string(msg) {
		t.Error("Xor Error: ", data)
	}
	data = XorDecodeBytes(*data, key)
	if string(*data) != string(msg) {
		t.Error("Xor Error: ", data)
	}
}

func TestXorStr(t *testing.T) {
	key := "1234567899"
	msg := "asdfghjkl"
	data := XorEncodeStr(msg, key)
	if *data == msg {
		t.Error("Xor Error: ", data)
	}
	data = XorDecodeStr(*data, key)
	if *data != msg {
		t.Error("Xor Error: ", data)
	}
}
