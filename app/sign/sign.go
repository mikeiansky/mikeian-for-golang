package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// -----BEGIN RSA PRIVATE KEY-----
// 这种格式直接用 ParsePKCS1PrivateKey
func LoadPKCS1PrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	// 读取文件内容
	keyData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read private key file failed: %w", err)
	}

	// 解析 PEM 块
	block, _ := pem.Decode(keyData)
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

func main() {

	fmt.Println("sign start ... ")

	fmt.Println("sign complete ... ")

}
