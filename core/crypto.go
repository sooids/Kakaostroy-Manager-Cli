package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"

	"github.com/tcolgate/gostikkit/evpkdf"
)

/*
	key_size := 8
	iv_size := 4
	iter := 1
*/
type OpenSSL_KDF struct {
	key_size  int
	iv_size   int
	salt_size int
	iter      int
}

func (o *OpenSSL_KDF) New(rand io.Reader, password []byte) []byte {
	salt := make([]byte, o.salt_size)
	if _, err := io.ReadFull(rand, salt); err != nil {
		panic(err)
	}

	return evpkdf.New(md5.New, password, salt, 4*(o.key_size+o.iv_size), o.iter)
}

func (o *OpenSSL_KDF) Generate(password []byte) ([]byte, []byte) {
	s := o.New(rand.Reader, password)
	key := s[:len(s)-16]
	iv := s[len(s)-16:]
	return key, iv
}

type PKCS7 struct {
	block_size int
}

func (p *PKCS7) Encode(data []byte) []byte {
	var pad_length int
	length := len(data)

	if rest := length % p.block_size; rest != 0 {
		pad_length = p.block_size - rest
	} else {
		pad_length = p.block_size
	}

	pad := make([]byte, 1)
	pad[0] = byte(pad_length)
	return append(data, bytes.Repeat(pad, pad_length)...)
}

func (p *PKCS7) Decode(data []byte) []byte {
	pad := int(data[len(data)-1])
	return data[:len(data)-pad]
}

type AES struct {
	block cipher.Block
}

func (a *AES) Encrypt(plaintext, iv []byte) []byte {
	if (len(plaintext)+len(iv))%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	ciphertext := append(iv, plaintext...)
	mode := cipher.NewCBCEncrypter(a.block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}
