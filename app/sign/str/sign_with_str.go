package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// -----BEGIN RSA PRIVATE KEY-----
// 这种格式直接用 ParsePKCS1PrivateKey
func LoadPKCS1PrivateKeyFromFile() (*rsa.PrivateKey, error) {
	pkContent := `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC+BzOaMuB0lGtj8rWS/VChLMWB9pAjmmI5MLyUNibHK3VZbvrv
u+C8FmNJP0CzrBNX8zcn3lTuqAyDksxBtlOvJtBDK7GFkWmUMtWYdebvWPZzVHT0
CzcmrZHEPdYEiCTpwW/ewqYytRWewpPZzm18+jfHDyk5R7WRKlfwUZeXkwIDAQAB
AoGAXCFltVcBV8Q1pMmhmthR5TKtt2rxSzGoeY3VgA4ZEutRA8E4zE5MkSRRksul
9PZykmuKzDuQ9fxnxMrBWQKq4yPb4cA+XihDHZw3sZekB+xN9/poIDqEwi7oHVeT
RETk+8CYqj9dYv/oD+XmZAE9VGA4/GOBNozzS0zi3Pq08eECQQDvRPEzSw5QJmh3
3xBQRpqvnkAO6CXgUnmY63soodDAFKtH8uGs5FitQOZOClyhkQWLA2Qmg55nbQ39
K6OUFSkpAkEAy1DQahDqF8qTfhMWRq3BzC72eKVjz65iDrefNq75fx9OI/wrBnO6
oWcuMve5aV+3cSjYDVt7JY7IXiAeaqMGWwJBAM0MpDh1BhsiDz1LaMui9kWytOsR
gAQyKgsnIzC7HA7Ap9jNCSIFvwkbKUOQFbpQchOkIFvxR/sStn5Uu6bYS6ECQCTM
rcxk31oLCoMpRUDptkgUv/u6Q8SVVZ5AUgr9xJ7jtasBu2/hhogsOspy2BZggr1y
nRhd1H5Jx65xtc+rTj0CQHsj3r4puhEEmbsXmm5HkLPAGSFsMv1NiRTyRQM70qal
uNr7Y9+CKoNuHaAd3KCw/orMIAudQ16JIOtBhuoI+fQ=
-----END RSA PRIVATE KEY-----
	`

	// 解析 PEM 块
	block, _ := pem.Decode([]byte(pkContent))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing private key")
	}

	// 解析 PKCS1 私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse PKCS1 private key failed: %w", err)
	}

	return privateKey, nil
}

// LoadRSAPublicKeyFromFile 从文件加载 RSA 公钥
func LoadRSAPublicKeyFromFile() (*rsa.PublicKey, error) {
	pubContent := `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC+BzOaMuB0lGtj8rWS/VChLMWB
9pAjmmI5MLyUNibHK3VZbvrvu+C8FmNJP0CzrBNX8zcn3lTuqAyDksxBtlOvJtBD
K7GFkWmUMtWYdebvWPZzVHT0CzcmrZHEPdYEiCTpwW/ewqYytRWewpPZzm18+jfH
Dyk5R7WRKlfwUZeXkwIDAQAB
-----END PUBLIC KEY-----
`

	// 解析 PEM 块
	block, _ := pem.Decode([]byte(pubContent))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing public key")
	}

	var publicKey *rsa.PublicKey

	// 尝试解析 PKCS#8 格式的公钥 (-----BEGIN PUBLIC KEY-----)
	if block.Type == "PUBLIC KEY" {
		fmt.Println("block type : PUBLIC KEY")
		pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("parse PKIX public key failed: %w", err)
		}
		publicKey = pubInterface.(*rsa.PublicKey)
	} else if block.Type == "RSA PUBLIC KEY" {
		fmt.Println("RSA PUBLIC KEY")
		// 尝试解析 PKCS#1 格式的公钥 (-----BEGIN RSA PUBLIC KEY-----)
		publicKey, _ = x509.ParsePKCS1PublicKey(block.Bytes)
		//if err != nil {
		//	return nil, fmt.Errorf("parse PKCS1 public key failed: %w", err)
		//}
	} else {
		return nil, fmt.Errorf("unsupported public key type: %s", block.Type)
	}

	return publicKey, nil
}

// SignData 对数据进行 RSA-SHA256 签名
func SignData(data string, privateKey *rsa.PrivateKey) (string, error) {
	// 计算数据的 SHA256 哈希
	hashed := sha256.Sum256([]byte(data))

	// 使用私钥进行 RSA-SHA256 签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("RSA sign failed: %w", err)
	}

	// 将签名转换为 Base64 字符串
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySignature 验证签名
func VerifySignature(data, signatureBase64 string, pub *rsa.PublicKey) bool {
	// 解码 Base64 签名
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Printf("Failed to decode signature: %v\n", err)
		return false
	}

	// 计算数据的 SHA256 哈希
	hashed := sha256.Sum256([]byte(data))

	// 使用公钥验证签名
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Printf("Signature verification failed: %v\n", err)
		return false
	}

	return true
}

func main() {

	fmt.Println("sign with str start ... ")
	pk, err := LoadPKCS1PrivateKeyFromFile()
	if err != nil {
		fmt.Printf("load private key failed: %v\n", err)
	}
	fmt.Println("pk ...", pk)

	data := "hello world"

	sd, err := SignData(data, pk)
	if err != nil {
		fmt.Printf("sign data failed: %v\n", err)
	}
	fmt.Println("sign ...", sd)

	pub, err := LoadRSAPublicKeyFromFile()
	if err != nil {
		fmt.Printf("load public key failed: %v\n", err)
	}
	fmt.Println("pub ...", pub)

	ret := VerifySignature(data, sd, pub)
	fmt.Println("verify signature", ret)

	fmt.Println("sign with str complete ... ")

}
