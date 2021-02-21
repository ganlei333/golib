package xor

func XorEncodeStr(msg, key []byte) []byte {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, 0, len(msg))

	for i := 0; i < ml; i++ {

		pwd = append(pwd, ((key[i%kl]) ^ (msg[i])))

	}

	return pwd

}

func XorDecodeStr(msg, key []byte) []byte {

	ml := len(msg)

	kl := len(key)

	pwd := make([]byte, 0, len(msg))

	for i := 0; i < ml; i++ {

		pwd = append(pwd, ((key[i%kl]) ^ (msg[i])))

	}

	return pwd

}
