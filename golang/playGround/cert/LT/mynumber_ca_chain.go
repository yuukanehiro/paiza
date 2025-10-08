package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

// CertificateAuthority 認証局の構造体
type CertificateAuthority struct {
	Certificate *x509.Certificate
	PrivateKey  *rsa.PrivateKey
	Name        string
}

// MyNumberCASystem マイナンバーカード証明書チェーンシステム
type MyNumberCASystem struct {
	RootCA         *CertificateAuthority
	IntermediateCA *CertificateAuthority
	UserCert       *x509.Certificate
	UserPrivateKey *rsa.PrivateKey
}

// NewMyNumberCASystem 新しいマイナンバーカードCAシステムを作成
func NewMyNumberCASystem() *MyNumberCASystem {
	return &MyNumberCASystem{}
}

// CreateRootCA JPKI Root CAを作成
func (mns *MyNumberCASystem) CreateRootCA() error {
	fmt.Println("🏛️ JPKI Root CA を作成中...")

	// Root CA用のRSA鍵ペアを生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096) // Root CAは4096bit
	if err != nil {
		return fmt.Errorf("Root CA鍵ペア生成エラー: %v", err)
	}

	// Root CA証明書テンプレート
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:      []string{"JP"},
			Organization: []string{"J-LIS"},
			CommonName:   "JPKI Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(20 * 365 * 24 * time.Hour), // 20年間有効
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2, // 中間CAを1つ許可
	}

	// 自己署名証明書を作成
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("Root CA証明書作成エラー: %v", err)
	}

	// 証明書をパース
	certificate, err := x509.ParseCertificate(certDER)
	if err != nil {
		return fmt.Errorf("Root CA証明書パースエラー: %v", err)
	}

	mns.RootCA = &CertificateAuthority{
		Certificate: certificate,
		PrivateKey:  privateKey,
		Name:        "JPKI Root CA",
	}

	fmt.Println("✅ JPKI Root CA を作成しました")
	return nil
}

// CreateIntermediateCA 中間CAを作成
func (mns *MyNumberCASystem) CreateIntermediateCA() error {
	if mns.RootCA == nil {
		return fmt.Errorf("Root CAが作成されていません")
	}

	fmt.Println("🏢 中間CA（公的個人認証サービス認証局）を作成中...")

	// 中間CA用のRSA鍵ペアを生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("中間CA鍵ペア生成エラー: %v", err)
	}

	// 中間CA証明書テンプレート
	template := x509.Certificate{
		SerialNumber: big.NewInt(1001),
		Subject: pkix.Name{
			Country:            []string{"JP"},
			Organization:       []string{"J-LIS"},
			OrganizationalUnit: []string{"公的個人認証サービス"},
			CommonName:         "JPKI 公的個人認証サービス CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour), // 10年間有効
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageEmailProtection, x509.ExtKeyUsageCodeSigning},
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            0, // エンドエンティティ証明書のみ発行可能
	}

	// Root CAで署名
	certDER, err := x509.CreateCertificate(rand.Reader, &template, mns.RootCA.Certificate, &privateKey.PublicKey, mns.RootCA.PrivateKey)
	if err != nil {
		return fmt.Errorf("中間CA証明書作成エラー: %v", err)
	}

	// 証明書をパース
	certificate, err := x509.ParseCertificate(certDER)
	if err != nil {
		return fmt.Errorf("中間CA証明書パースエラー: %v", err)
	}

	mns.IntermediateCA = &CertificateAuthority{
		Certificate: certificate,
		PrivateKey:  privateKey,
		Name:        "JPKI 公的個人認証サービス CA",
	}

	fmt.Println("✅ 中間CA を作成しました")
	return nil
}

// CreateUserCertificate エンドユーザー証明書を作成
func (mns *MyNumberCASystem) CreateUserCertificate(userName, myNumber string) error {
	if mns.IntermediateCA == nil {
		return fmt.Errorf("中間CAが作成されていません")
	}

	fmt.Printf("👤 ユーザー証明書を作成中: %s\n", userName)

	// ユーザー用のRSA鍵ペアを生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("ユーザー鍵ペア生成エラー: %v", err)
	}

	// ユーザー証明書テンプレート
	template := x509.Certificate{
		SerialNumber: big.NewInt(2001),
		Subject: pkix.Name{
			Country:            []string{"JP"},
			Organization:       []string{"J-LIS"},
			OrganizationalUnit: []string{"公的個人認証サービス"},
			CommonName:         userName,
			ExtraNames: []pkix.AttributeTypeAndValue{
				{
					Type:  asn1.ObjectIdentifier{1, 2, 392, 200149, 8, 5, 5, 1}, // 個人番号OID
					Value: myNumber,
				},
			},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(5 * 365 * 24 * time.Hour), // 5年間有効
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageEmailProtection, x509.ExtKeyUsageCodeSigning},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	// 中間CAで署名
	certDER, err := x509.CreateCertificate(rand.Reader, &template, mns.IntermediateCA.Certificate, &privateKey.PublicKey, mns.IntermediateCA.PrivateKey)
	if err != nil {
		return fmt.Errorf("ユーザー証明書作成エラー: %v", err)
	}

	// 証明書をパース
	certificate, err := x509.ParseCertificate(certDER)
	if err != nil {
		return fmt.Errorf("ユーザー証明書パースエラー: %v", err)
	}

	mns.UserCert = certificate
	mns.UserPrivateKey = privateKey

	fmt.Printf("✅ ユーザー証明書を作成しました: %s\n", userName)
	return nil
}

// SaveCertificateChain 証明書チェーンをファイルに保存
func (mns *MyNumberCASystem) SaveCertificateChain() error {
	// Root CA証明書を保存
	err := os.WriteFile("jpki_root_ca.der", mns.RootCA.Certificate.Raw, 0644)
	if err != nil {
		return fmt.Errorf("Root CA保存エラー: %v", err)
	}

	// 中間CA証明書を保存
	err = os.WriteFile("jpki_intermediate_ca.der", mns.IntermediateCA.Certificate.Raw, 0644)
	if err != nil {
		return fmt.Errorf("中間CA保存エラー: %v", err)
	}

	// ユーザー証明書を保存
	err = os.WriteFile("user_cert.der", mns.UserCert.Raw, 0644)
	if err != nil {
		return fmt.Errorf("ユーザー証明書保存エラー: %v", err)
	}

	fmt.Println("💾 証明書チェーンを保存しました:")
	fmt.Println("   - jpki_root_ca.der (Root CA)")
	fmt.Println("   - jpki_intermediate_ca.der (中間CA)")
	fmt.Println("   - user_cert.der (ユーザー証明書)")

	return nil
}

// PrintCertificateChain 証明書チェーン情報を表示
func (mns *MyNumberCASystem) PrintCertificateChain() {
	fmt.Println("\n📋 証明書チェーン情報:")
	fmt.Println(strings.Repeat("=", 80))

	// Root CA情報
	fmt.Println("🏛️ Root CA:")
	fmt.Printf("   発行者: %s\n", mns.RootCA.Certificate.Issuer.CommonName)
	fmt.Printf("   所有者: %s\n", mns.RootCA.Certificate.Subject.CommonName)
	fmt.Printf("   有効期間: %s ～ %s\n",
		mns.RootCA.Certificate.NotBefore.Format("2006-01-02"),
		mns.RootCA.Certificate.NotAfter.Format("2006-01-02"))
	fmt.Printf("   CA証明書: %t\n", mns.RootCA.Certificate.IsCA)
	fmt.Printf("   最大パス長: %d\n", mns.RootCA.Certificate.MaxPathLen)

	// 中間CA情報
	fmt.Println("\n🏢 中間CA:")
	fmt.Printf("   発行者: %s\n", mns.IntermediateCA.Certificate.Issuer.CommonName)
	fmt.Printf("   所有者: %s\n", mns.IntermediateCA.Certificate.Subject.CommonName)
	fmt.Printf("   有効期間: %s ～ %s\n",
		mns.IntermediateCA.Certificate.NotBefore.Format("2006-01-02"),
		mns.IntermediateCA.Certificate.NotAfter.Format("2006-01-02"))
	fmt.Printf("   CA証明書: %t\n", mns.IntermediateCA.Certificate.IsCA)
	fmt.Printf("   最大パス長: %d\n", mns.IntermediateCA.Certificate.MaxPathLen)

	// ユーザー証明書情報
	fmt.Println("\n👤 ユーザー証明書:")
	fmt.Printf("   発行者: %s\n", mns.UserCert.Issuer.CommonName)
	fmt.Printf("   所有者: %s\n", mns.UserCert.Subject.CommonName)
	fmt.Printf("   有効期間: %s ～ %s\n",
		mns.UserCert.NotBefore.Format("2006-01-02"),
		mns.UserCert.NotAfter.Format("2006-01-02"))
	fmt.Printf("   CA証明書: %t\n", mns.UserCert.IsCA)

	// マイナンバー情報
	for _, extra := range mns.UserCert.Subject.ExtraNames {
		if extra.Type.Equal(asn1.ObjectIdentifier{1, 2, 392, 200149, 8, 5, 5, 1}) {
			fmt.Printf("   マイナンバー: %s\n", extra.Value)
		}
	}
}

// VerifyCertificateChain 証明書チェーンを検証
func (mns *MyNumberCASystem) VerifyCertificateChain() error {
	fmt.Println("\n🔍 証明書チェーンを検証中...")

	// Root証明書プールを作成
	rootPool := x509.NewCertPool()
	rootPool.AddCert(mns.RootCA.Certificate)

	// 中間証明書プールを作成
	intermediatePool := x509.NewCertPool()
	intermediatePool.AddCert(mns.IntermediateCA.Certificate)

	// ユーザー証明書を検証
	opts := x509.VerifyOptions{
		Roots:         rootPool,
		Intermediates: intermediatePool,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}

	chains, err := mns.UserCert.Verify(opts)
	if err != nil {
		return fmt.Errorf("❌ 証明書チェーン検証失敗: %v", err)
	}

	fmt.Println("✅ 証明書チェーン検証成功!")
	fmt.Printf("   検証されたチェーン数: %d\n", len(chains))

	for i, chain := range chains {
		fmt.Printf("   チェーン %d:\n", i+1)
		for j, cert := range chain {
			fmt.Printf("     %d. %s\n", j+1, cert.Subject.CommonName)
		}
	}

	return nil
}

// SignDocument 文書に署名（証明書チェーン付き）
func (mns *MyNumberCASystem) SignDocument(documentPath, signatureFile string) error {
	fmt.Printf("✍️ 証明書チェーン付きで文書に署名中: %s\n", documentPath)

	// 文書を読み込み
	documentData, err := os.ReadFile(documentPath)
	if err != nil {
		return fmt.Errorf("文書読み込みエラー: %v", err)
	}

	// SHA-256ハッシュを計算
	hash := sha256.Sum256(documentData)
	fmt.Printf("📊 文書のSHA-256ハッシュ: %x\n", hex.EncodeToString(hash[:]))

	// PKCS#1 v1.5署名を生成
	signature, err := rsa.SignPKCS1v15(rand.Reader, mns.UserPrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return fmt.Errorf("署名生成エラー: %v", err)
	}

	// 証明書チェーンを長さ付きで保存（TLV形式風）
	var chainData []byte

	// 署名（固定256バイト）
	chainData = append(chainData, signature...)

	// ユーザー証明書（長さ + データ）
	userCertLen := uint32(len(mns.UserCert.Raw))
	chainData = append(chainData, byte(userCertLen>>24), byte(userCertLen>>16), byte(userCertLen>>8), byte(userCertLen))
	chainData = append(chainData, mns.UserCert.Raw...)

	// 中間CA証明書（長さ + データ）
	intermediateCertLen := uint32(len(mns.IntermediateCA.Certificate.Raw))
	chainData = append(chainData, byte(intermediateCertLen>>24), byte(intermediateCertLen>>16), byte(intermediateCertLen>>8), byte(intermediateCertLen))
	chainData = append(chainData, mns.IntermediateCA.Certificate.Raw...)

	// Root CA証明書（長さ + データ）
	rootCertLen := uint32(len(mns.RootCA.Certificate.Raw))
	chainData = append(chainData, byte(rootCertLen>>24), byte(rootCertLen>>16), byte(rootCertLen>>8), byte(rootCertLen))
	chainData = append(chainData, mns.RootCA.Certificate.Raw...)

	// ファイルに保存
	err = os.WriteFile(signatureFile, chainData, 0644)
	if err != nil {
		return fmt.Errorf("署名保存エラー: %v", err)
	}

	fmt.Printf("✅ 証明書チェーン付き署名を生成: %s\n", signatureFile)
	fmt.Printf("   署名長: %d bytes\n", len(signature))
	fmt.Printf("   ユーザー証明書長: %d bytes\n", len(mns.UserCert.Raw))
	fmt.Printf("   中間CA証明書長: %d bytes\n", len(mns.IntermediateCA.Certificate.Raw))
	fmt.Printf("   Root CA証明書長: %d bytes\n", len(mns.RootCA.Certificate.Raw))

	return nil
}

// VerifyDocumentWithChain 証明書チェーンを使用して文書の署名を検証
func VerifyDocumentWithChain(documentPath, signatureFile string) error {
	fmt.Printf("🔍 証明書チェーン付き署名を検証中...\n")
	fmt.Printf("   文書: %s\n", documentPath)
	fmt.Printf("   署名ファイル: %s\n", signatureFile)

	// 文書を読み込み
	documentData, err := os.ReadFile(documentPath)
	if err != nil {
		return fmt.Errorf("文書読み込みエラー: %v", err)
	}

	// 署名ファイルを読み込み
	signatureData, err := os.ReadFile(signatureFile)
	if err != nil {
		return fmt.Errorf("署名ファイル読み込みエラー: %v", err)
	}

	// 署名を取得（固定256バイト）
	if len(signatureData) < 256 {
		return fmt.Errorf("署名データが不正です")
	}
	signature := signatureData[:256]
	offset := 256

	// ユーザー証明書を読み込み（長さ + データ）
	if len(signatureData) < offset+4 {
		return fmt.Errorf("証明書データが不正です")
	}
	userCertLen := uint32(signatureData[offset])<<24 | uint32(signatureData[offset+1])<<16 | uint32(signatureData[offset+2])<<8 | uint32(signatureData[offset+3])
	offset += 4

	if len(signatureData) < offset+int(userCertLen) {
		return fmt.Errorf("ユーザー証明書データが不正です")
	}
	userCert, err := x509.ParseCertificate(signatureData[offset : offset+int(userCertLen)])
	if err != nil {
		return fmt.Errorf("ユーザー証明書パースエラー: %v", err)
	}
	offset += int(userCertLen)

	// 中間CA証明書を読み込み
	if len(signatureData) < offset+4 {
		return fmt.Errorf("中間CA証明書長データが不正です")
	}
	intermediateCertLen := uint32(signatureData[offset])<<24 | uint32(signatureData[offset+1])<<16 | uint32(signatureData[offset+2])<<8 | uint32(signatureData[offset+3])
	offset += 4

	if len(signatureData) < offset+int(intermediateCertLen) {
		return fmt.Errorf("中間CA証明書データが不正です")
	}
	intermediateCert, err := x509.ParseCertificate(signatureData[offset : offset+int(intermediateCertLen)])
	if err != nil {
		return fmt.Errorf("中間CA証明書パースエラー: %v", err)
	}
	offset += int(intermediateCertLen)

	// Root CA証明書を読み込み
	if len(signatureData) < offset+4 {
		return fmt.Errorf("Root CA証明書長データが不正です")
	}
	rootCertLen := uint32(signatureData[offset])<<24 | uint32(signatureData[offset+1])<<16 | uint32(signatureData[offset+2])<<8 | uint32(signatureData[offset+3])
	offset += 4

	if len(signatureData) < offset+int(rootCertLen) {
		return fmt.Errorf("Root CA証明書データが不正です")
	}
	rootCert, err := x509.ParseCertificate(signatureData[offset : offset+int(rootCertLen)])
	if err != nil {
		return fmt.Errorf("Root CA証明書パースエラー: %v", err)
	}

	fmt.Printf("📋 署名者: %s\n", userCert.Subject.CommonName)
	fmt.Printf("📋 中間CA: %s\n", intermediateCert.Subject.CommonName)
	fmt.Printf("📋 Root CA: %s\n", rootCert.Subject.CommonName)

	// 証明書チェーンを検証
	rootPool := x509.NewCertPool()
	rootPool.AddCert(rootCert)

	intermediatePool := x509.NewCertPool()
	intermediatePool.AddCert(intermediateCert)

	opts := x509.VerifyOptions{
		Roots:         rootPool,
		Intermediates: intermediatePool,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	}

	chains, err := userCert.Verify(opts)
	if err != nil {
		return fmt.Errorf("❌ 証明書チェーン検証失敗: %v", err)
	}

	fmt.Printf("✅ 証明書チェーン検証成功 (チェーン数: %d)\n", len(chains))

	// 署名を検証
	publicKey, ok := userCert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("RSA公開鍵ではありません")
	}

	hash := sha256.Sum256(documentData)
	fmt.Printf("📊 文書のSHA-256ハッシュ: %x\n", hex.EncodeToString(hash[:]))

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return fmt.Errorf("❌ 署名検証失敗: 文書が改ざんされています (%v)", err)
	}

	fmt.Println("✅ 署名検証成功: 文書は改ざんされていません")
	fmt.Printf("✅ 署名者の本人確認: %s\n", userCert.Subject.CommonName)

	// マイナンバー表示
	for _, extra := range userCert.Subject.ExtraNames {
		if extra.Type.Equal(asn1.ObjectIdentifier{1, 2, 392, 200149, 8, 5, 5, 1}) {
			fmt.Printf("✅ マイナンバー確認: %s\n", extra.Value)
		}
	}

	return nil
}

func main() {
	fmt.Println("🏛️ マイナンバーカード証明書チェーンシステム")
	fmt.Println("===========================================")

	// CAシステムを初期化
	caSystem := NewMyNumberCASystem()

	// 1. Root CAを作成
	err := caSystem.CreateRootCA()
	if err != nil {
		log.Fatalf("Root CA作成エラー: %v", err)
	}

	// 2. 中間CAを作成
	err = caSystem.CreateIntermediateCA()
	if err != nil {
		log.Fatalf("中間CA作成エラー: %v", err)
	}

	// 3. ユーザー証明書を作成
	err = caSystem.CreateUserCertificate("田中花子", "9876543210123")
	if err != nil {
		log.Fatalf("ユーザー証明書作成エラー: %v", err)
	}

	// 4. 証明書チェーン情報を表示
	caSystem.PrintCertificateChain()

	// 5. 証明書チェーンを検証
	err = caSystem.VerifyCertificateChain()
	if err != nil {
		log.Fatalf("証明書チェーン検証エラー: %v", err)
	}

	// 6. 証明書チェーンをファイルに保存
	err = caSystem.SaveCertificateChain()
	if err != nil {
		log.Fatalf("証明書チェーン保存エラー: %v", err)
	}

	// 7. サンプル文書を作成
	documentContent := `電子納税申告書

申告者: 田中花子（マイナンバー: 9876543210123）
申告年度: 2024年（令和6年）
申告日: 2024年09月28日
申告内容: 所得税確定申告

【収入情報】
- 給与所得: 5,000,000円
- 配当所得: 50,000円

【所得控除】
- 基礎控除: 480,000円
- 社会保険料控除: 700,000円

本申告書はマイナンバーカードによる電子署名で保護されています。
改ざんされた場合、署名検証が失敗し、申告が無効となります。

【重要】この申告書は公的個人認証サービスによる証明書チェーンで認証されています。`

	documentFile := "tax_return.txt"
	err = os.WriteFile(documentFile, []byte(documentContent), 0644)
	if err != nil {
		log.Fatalf("文書作成エラー: %v", err)
	}
	fmt.Printf("\n📄 電子納税申告書を作成: %s\n", documentFile)

	// 8. 証明書チェーン付きで文書に署名
	signatureFile := "tax_return.p7c" // 証明書チェーン付きPKCS#7
	err = caSystem.SignDocument(documentFile, signatureFile)
	if err != nil {
		log.Fatalf("署名エラー: %v", err)
	}

	fmt.Println("\n" + strings.Repeat("=", 80))

	// 9. 署名を検証（正常ケース）
	fmt.Println("【正常ケース】証明書チェーン付き署名検証")
	err = VerifyDocumentWithChain(documentFile, signatureFile)
	if err != nil {
		fmt.Printf("検証エラー: %v\n", err)
	}

	fmt.Println("\n" + strings.Repeat("=", 80))

	// 10. 文書を改ざんして検証（異常ケース）
	fmt.Println("【異常ケース】改ざんされた文書で署名検証")
	tamperedContent := documentContent + "\n\n[不正な追記: 給与所得を10,000,000円に変更]"
	tamperedFile := "tampered_tax_return.txt"
	err = os.WriteFile(tamperedFile, []byte(tamperedContent), 0644)
	if err != nil {
		log.Fatalf("改ざん文書作成エラー: %v", err)
	}

	err = VerifyDocumentWithChain(tamperedFile, signatureFile)
	if err != nil {
		fmt.Printf("検証結果: %v\n", err)
	}

	fmt.Println("\n🎯 証明書チェーンシステム実行完了!")
	fmt.Println("\n生成されたファイル:")
	fmt.Println("  - jpki_root_ca.der (JPKI Root CA)")
	fmt.Println("  - jpki_intermediate_ca.der (中間CA)")
	fmt.Println("  - user_cert.der (ユーザー証明書)")
	fmt.Println("  - tax_return.txt (電子納税申告書)")
	fmt.Println("  - tax_return.p7c (証明書チェーン付き署名)")
	fmt.Println("  - tampered_tax_return.txt (改ざんされた申告書)")

	fmt.Println("\n📋 証明書チェーンの特徴:")
	fmt.Println("  - 3層構造: Root CA → 中間CA → ユーザー証明書")
	fmt.Println("  - 信頼の連鎖による検証")
	fmt.Println("  - Root CAは長期間有効（20年）")
	fmt.Println("  - 中間CAは中期間有効（10年）")
	fmt.Println("  - ユーザー証明書は短期間有効（5年）")
	fmt.Println("  - 各レベルでの証明書失効管理（CRL/OCSP）")
}