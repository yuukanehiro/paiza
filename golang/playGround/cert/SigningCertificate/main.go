package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	message := "Hello, World! This is a test message for signature verification."
	
	// 鍵ペア生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Failed to generate private key:", err)
	}
	
	// 公開鍵取得
	publicKey := &privateKey.PublicKey
	
	// signature ... 署名。文言をsha256でhash化し秘密鍵で署名したもの
	signature, err := signString(message, privateKey)
	if err != nil {
		log.Fatal("Failed to sign message:", err)
	}
	
	fmt.Printf("Original message: %s\n", message)
	fmt.Printf("Signature (hex): %x\n", signature)
	
	// verify ... 公開鍵で署名を検証
	valid, err := verifySignature(message, signature, publicKey)
	if err != nil {
		log.Fatal("Failed to verify signature:", err)
	}
	
	fmt.Printf("Signature valid: %t\n", valid)
	
	invalidMessage := "This is a tampered message."
	validInvalid, err := verifySignature(invalidMessage, signature, publicKey)
	if err != nil {
		log.Fatal("Failed to verify invalid signature:", err)
	}
	
	fmt.Printf("Invalid message signature valid: %t\n", validInvalid)
	
	keyPair, err := generateKeyPair()
	if err != nil {
		log.Fatal("Failed to generate key pair:", err)
	}
	
	fmt.Println("\n--- Generated Key Pair ---")
	fmt.Printf("Private Key (PEM):\n%s\n", keyPair.PrivateKeyPEM)
	fmt.Printf("Public Key (PEM):\n%s\n", keyPair.PublicKeyPEM)
}

func signString(message string, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256([]byte(message))
	
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}
	
	return signature, nil
}

func verifySignature(message string, signature []byte, publicKey *rsa.PublicKey) (bool, error) {
	hash := sha256.Sum256([]byte(message))
	
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, nil
	}
	
	return true, nil
}

type KeyPair struct {
	PrivateKeyPEM string
	PublicKeyPEM  string
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
}

func generateKeyPair() (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}
	
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}
	
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	
	return &KeyPair{
		PrivateKeyPEM: string(privateKeyPEM),
		PublicKeyPEM:  string(publicKeyPEM),
		PrivateKey:    privateKey,
		PublicKey:     &privateKey.PublicKey,
	}, nil
}

// Output:
// % go run main.go
// Original message: Hello, World! This is a test message for signature verification.
// Signature (hex): 4848a873abb006c55505b1324580cd7e82d1362b9ab6e6048e45512ffeb0b8dd5c07fc6681cc0633a7e479b3d358ac3e447139035ad52644cf5be1b73f455690d15a5fd348201ba6554aeb2a542e794082f6eafb6e1d7bae48f64b2a67dcc76c3c78c2b3cd1c41bcc941bb5466d42f102b6f1cdae54c52dda5cc93dfaed55fec1edc1dbf0199d0c1519289a3545bd92f301473ad229027ea971b0868a2eea1423e32126c400000444c4a07f37ef989faf1f38b6fd37f994a009d36ee321da3fd345ad93bb537062a9675d46e6c605860ce16f07fc0b843d7f5030877837d1ae78952405167c62141961d44b31b876bb0687ca246c22def98e645f4c7c3615a40
// Signature valid: true
// Invalid message signature valid: false

// --- Generated Key Pair ---
// Private Key (PEM):
// -----BEGIN RSA PRIVATE KEY-----
// MIIEogIBAAKCAQEAqurZI28ysRgcEvIF/MW359ZbJyP8o4v0qKeRL1GL5FA0Rmd+
// HLTXLRMhom7JElTSvmnYd8N3IVIHpRZJ7nAGax0rvA0yqADoEOYKGVkA4puuJie6
// g1rw8JFcua2/+5GL9uah6BWhcxnpxOOsriQPKTfrc4OQMPxMPjasZ6iwn9xgowG6
// MFVsbArYj+Vg4Hl+ghEGaOZTT7g4R+jeFZPs/SBNDyuvCD1aPxMhOhXkcrZ9u6Qx
// wI+eFYpKDn0F4zT+1Tdx81gO2QJulpj7EOiM4ZutBkhf41RNLvRRqldZH2CNx4Pn
// sgr9zC3QSfVM7B8HdJBRkfqZ6ej+bxiY0XuHQQIDAQABAoIBAFVIW9KHgLbG/fMi
// GoS3L32aAt+z4DLG+exRuGrw+KmSr5LVvyKZxahzHFxNhT0FYCas79tKiB3zLWdd
// iMRV144I/zKVL0qPMTyFXFKri3qHDc1yE2nJreGL1sW7gckIJdur1uvUJzwkmPWq
// WRaHXZ1GjtCTvBYFPhHo+/u5O3B4GGORCTPkaZfnYHdi11oCgwDCuUng3IrLrJcN
// SlgWHfaoh7YQhB2jKNPoA5/cEQvNiEOseIENd26wIMt1Fn1VakFbQGxeOGPnKy+H
// dYm6z6GeIjusrfTdLri3XXT2Q+B3gtOuSnxSVKmJgGxmzFtpX3fDrNp+kefuDNtf
// gszCusECgYEA1QOuRkVT2DcAO5HhgD7DWAvzD9vMGmoyZoOpoDyTn08T6uhdl7P4
// mNpUWa0v8RAKeGUJh7QHYLerCyTHwhuJOeiNTvbuQ3RXW5YkwpMSHlESzASQuSzy
// Ctdg/l44xeIVlQ46luJ5rO8k5ofc1fiOfk+khbg6YGx0bxcuymyBImkCgYEAzWhv
// 30FViS6jUi9VE6i4F3O3QLC0jtcx4w/ZDZD5rhz5dYG/PTHIJz3sxpCihkfSz4ID
// 6slfJOK/rWxh5j/e4H+OArRGv4g2j9Zm62UayZ9XGT+KL6b56o3TjJFxauy6BkP9
// ruOWKW6oz4iP5/RLdTo4qNkXCCCeDX0YiIHacxkCgYAGmfGZreZWSgYQ0omJfuhw
// 3xXuROKDYw0izivAkoAErN+yJPdZjzNCEWoRyWM19khMyNzNvs+HuMltM/uY0V+k
// UsemTdK+dmmKphGHkiU6mUqa1f1iL7lkip4v492gjEwc7W71ZBlfOZ4MIuO+IdDQ
// q0ku97GfcVz+YFqL3Wb3GQKBgDpoBt95iWPfggsdcFiNsgLhIWtgk4bGQ+YnnHVS
// EtziWkCCaylwXVRjNE9l0wDRrWzZsFlzdv1bCFu8pL2+zZYSTwQP1MojZjhSI6Ot
// dQd9qluAiL9yAEAkodGnjZN0ypR7gsfW6NzRHWkdkKXI/ifQLUo9qGHkPGzjdDVk
// jwQJAoGACDiaUh3dpSX243mMHjeSay7+Y48xxmJSgRP/Dk99yMQ8X2l2qZGck22c
// 2tFdiMr/V8YHFwuXE1HUDbiZVjxvhBJxuOxN/21ZzRbABKarT+Q6tygrL3FnXHFz
// Wl3n9ZgO0ZfLI+JKvE8Zh26J23PbEa37UeHA7nySJG70SJC19Gw=
// -----END RSA PRIVATE KEY-----

// Public Key (PEM):
// -----BEGIN PUBLIC KEY-----
// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqurZI28ysRgcEvIF/MW3
// 59ZbJyP8o4v0qKeRL1GL5FA0Rmd+HLTXLRMhom7JElTSvmnYd8N3IVIHpRZJ7nAG
// ax0rvA0yqADoEOYKGVkA4puuJie6g1rw8JFcua2/+5GL9uah6BWhcxnpxOOsriQP
// KTfrc4OQMPxMPjasZ6iwn9xgowG6MFVsbArYj+Vg4Hl+ghEGaOZTT7g4R+jeFZPs
// /SBNDyuvCD1aPxMhOhXkcrZ9u6QxwI+eFYpKDn0F4zT+1Tdx81gO2QJulpj7EOiM
// 4ZutBkhf41RNLvRRqldZH2CNx4Pnsgr9zC3QSfVM7B8HdJBRkfqZ6ej+bxiY0XuH
// QQIDAQAB
// -----END PUBLIC KEY-----