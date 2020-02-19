package rsa

import (
	"crypto"
	basecrypto "crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

// GenKey 生成RSA公私钥文件
func GenKey(bits int) error {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 生成PKCS8私钥文件

	derPKCS8Stream := MarshalPKCS8PrivateKey(privateKey)
	blockPKCS8 := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derPKCS8Stream,
	}
	filePKCS8, err := os.Create("pkcs8_private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(filePKCS8, blockPKCS8)
	if err != nil {
		return err
	}

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 定义：引用IETF的安全领域的公钥基础实施（PKIX）工作组的标准实例化内容
	subject := pkix.Name{
		Organization:       []string{"www.gopper.in"},
		OrganizationalUnit: []string{"ITs"},
		CommonName:         "www.gopper.in Web",
	}

	// 设置 SSL证书的属性用途
	certificate509 := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// 生成 SSL公匙
	derBytes, _ := x509.CreateCertificate(rand.Reader, &certificate509, &certificate509, &privateKey.PublicKey, privateKey)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}

// MarshalPKCS8PrivateKey MarshalPKCS8PrivateKey
func MarshalPKCS8PrivateKey(key *rsa.PrivateKey) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = x509.MarshalPKCS1PrivateKey(key)

	k, err := asn1.Marshal(info)
	if err != nil {
		log.Panic(err.Error())
	}
	return k
}

// Encrypt 公钥加密
func Encrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// Decrypt 私钥解密
func Decrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Sign rsa sign
// hashed := md5.Sum(src)
// hashed := sha256.Sum256(src)
func Sign(hash basecrypto.Hash, hashed, privateKey []byte) ([]byte, error) {
	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
	// return rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, hashed[:])
	// return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
}

// Verify rsa verify
// hashed := md5.Sum(src)
// hashed := sha256.Sum256(src)
// signature is a valid signature of message from the public key.
func Verify(hash basecrypto.Hash, hashed, publicKey, signature []byte) error {

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.MD5, hashed[:], signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// CertVerify rsa verify
// hashed := md5.Sum(src)
// hashed := sha256.Sum256(src)
// signature is a valid signature of message from the public key.
func CertVerify(hash crypto.Hash, hashed, cert, signature []byte) error {

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).

	block, _ := pem.Decode(cert)
	var _cert *x509.Certificate
	_cert, _ = x509.ParseCertificate(block.Bytes)
	pub := _cert.PublicKey.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.MD5, hashed[:], signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}
