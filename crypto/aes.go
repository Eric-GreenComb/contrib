package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// EncryptAes encrypt by aes
func EncryptAes(origData, key []byte) ([]byte, error) {
	result, err := AesEncrypt(origData, key)
	if err != nil {
		return nil, err
	}
	_enBase64 := base64.StdEncoding.EncodeToString(result)
	return []byte(_enBase64), nil
}

// DecryptAes decrypt by aes
func DecryptAes(crypted, key []byte) ([]byte, error) {
	_deBase64, err := base64.StdEncoding.DecodeString(string(crypted))
	if err != nil {
		return nil, err
	}
	origData, err := AesDecrypt(_deBase64, key)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// EncryptAesString encrypt aes 2 string
func EncryptAesString(origData, key string) (string, error) {
	result, err := AesEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return "", err
	}
	_enBase64 := base64.StdEncoding.EncodeToString(result)
	return _enBase64, nil
}

// DecryptAesString decrypt aes 2 string
func DecryptAesString(crypted, key string) (string, error) {
	_enBase64, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	origData, err := AesDecrypt(_enBase64, []byte(key))
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

// AesEncrypt encrypt aes 2 []byte
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt decrypt aes 2 []byte
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
