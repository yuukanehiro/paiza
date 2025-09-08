package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	crlURLs := []string{
		"http://crl.globalsign.com/gs/gsorganizationvalsha2g2.crl",
		"http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl",
		"http://crl.comodoca.com/COMODORSAOrganizationValidationSecureServerCA.crl",
	}

	for i, url := range crlURLs {
		fmt.Printf("=== CRL %d: %s ===\n", i+1, url)
		
		crl, err := downloadAndParseCRL(url)
		if err != nil {
			fmt.Printf("Error processing CRL: %v\n\n", err)
			continue
		}
		
		displayCRLInfo(crl)
		fmt.Println()
	}
	
	fmt.Println("=== CRL Certificate Check Demo ===")
	demoRevokedCertCheck()
}

func downloadAndParseCRL(url string) (*pkix.CertificateList, error) {
	fmt.Printf("Downloading CRL from: %s\n", url)
	
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download CRL: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}
	
	crlData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read CRL data: %w", err)
	}
	
	fmt.Printf("Downloaded %d bytes\n", len(crlData))
	
	crl, err := x509.ParseCRL(crlData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CRL: %w", err)
	}
	
	return crl, nil
}

func displayCRLInfo(crl *pkix.CertificateList) {
	fmt.Printf("Issuer: %s\n", crl.TBSCertList.Issuer.String())
	fmt.Printf("This Update: %s\n", crl.TBSCertList.ThisUpdate.Format(time.RFC3339))
	
	if !crl.TBSCertList.NextUpdate.IsZero() {
		fmt.Printf("Next Update: %s\n", crl.TBSCertList.NextUpdate.Format(time.RFC3339))
	} else {
		fmt.Println("Next Update: Not specified")
	}
	
	fmt.Printf("Signature Algorithm: %s\n", crl.SignatureAlgorithm.Algorithm.String())
	fmt.Printf("Revoked Certificates Count: %d\n", len(crl.TBSCertList.RevokedCertificates))
	
	if len(crl.TBSCertList.RevokedCertificates) > 0 {
		fmt.Printf("First 5 revoked certificates:\n")
		for i, revokedCert := range crl.TBSCertList.RevokedCertificates {
			if i >= 5 {
				break
			}
			fmt.Printf("  Serial: %s, Revocation Date: %s\n",
				revokedCert.SerialNumber.String(),
				revokedCert.RevocationTime.Format(time.RFC3339))
			
			if len(revokedCert.Extensions) > 0 {
				for _, ext := range revokedCert.Extensions {
					if ext.Id.Equal([]int{2, 5, 29, 21}) { // CRL Reason Code
						fmt.Printf("    Reason Code: %x\n", ext.Value)
					}
				}
			}
		}
	}
}

func demoRevokedCertCheck() {
	fmt.Println("This demonstrates how to check if a certificate is revoked using CRL")
	
	exampleCertPEM := `-----BEGIN CERTIFICATE-----
MIIDdTCCAl2gAwIBAgILBAAAAAABFUtaw5QwDQYJKoZIhvcNAQEFBQAwVzELMAkG
A1UEBhMCQkUxGTAXBgNVBAoTEEdsb2JhbFNpZ24gbnYtc2ExEDAOBgNVBAsTB1Jv
b3QgQ0ExGzAZBgNVBAMTEkdsb2JhbFNpZ24gUm9vdCBDQTAeFw05ODA5MDExMjAw
MDBaFw0yODAxMjgxMjAwMDBaMFcxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9i
YWxTaWduIG52LXNhMRAwDgYDVQQLEwdSb290IENBMRswGQYDVQQDExJHbG9iYWxT
aWduIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDaDuaZ
jc6j40+Kfvvxi4Mla+pIH/EqsLmVEQS98GPR4mdmzxzdzxtIK+6NiY6arymAZavp
xy0Sy6scTHAHoT0KMM0VjU/43dSMUBUc71DuxC73/OlS8pF94G3VNTCOXkNz8kHp
1Wrjsok6Vjk4bwY8iGlbKk3Fp1S4bInMm/k8yuX9ifUSPJJ4ltbcdG6TRGHRjcdG
snUOhugZitVtbNV4FpWi6cgKOOvyJBNPc1STE4U6G7weNLWLBYy5d4ux2x8gkasJ
U26Qzns3dLlwR5EiUWMWea6xrkEmCMgZK9FGqkjWZCrXgzT/LCrBbBlDSgeF59N8
9iFo7+ryUp9/k5DPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBRge2YaRQ2XyolQL30EzTSo//z9SzANBgkqhkiG9w0B
AQUFAAOCAQEA1nPnfE920I2/7LqivjTFKDK1fPxsnCwrvQmeU79rXqoRSLblCKOz
yj1hTdNGCbM+w6DjY1Ub8rrvrTnhQ7k4o+YviiY776BQVvnGCv04zcQLcFGUl5gE
38NflNUVyRRBnMRddWQVDf9VMOyGj/8N7yy5Y0b2qvzfvGn9LhJIZJrglfCm7ymP
AbEVtQwdpf5pLGkkeB6zpxxxYu7KyJesF12KwvhHhm4qxFYxldBniYUr+WymXUad
DKqC5JlR3XC321Y9YeRq4VzW9v493kHMB65jUr9TU/Qr6cf9tveCX4XSQRjbgbME
HMUfpIBvFSDJ3gyICh3WZlXi/EjJKSZp4A==
-----END CERTIFICATE-----`
	
	block, _ := pem.Decode([]byte(exampleCertPEM))
	if block == nil {
		fmt.Println("Failed to decode certificate PEM")
		return
	}
	
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse certificate: %v\n", err)
		return
	}
	
	fmt.Printf("Example Certificate:\n")
	fmt.Printf("  Subject: %s\n", cert.Subject.String())
	fmt.Printf("  Serial Number: %s\n", cert.SerialNumber.String())
	fmt.Printf("  Issuer: %s\n", cert.Issuer.String())
	
	fmt.Println("\nTo check if this certificate is revoked:")
	fmt.Println("1. Find the CRL Distribution Points in the certificate")
	fmt.Println("2. Download the CRL from those URLs")
	fmt.Println("3. Check if the certificate's serial number is in the revoked list")
	
	if len(cert.CRLDistributionPoints) > 0 {
		fmt.Printf("CRL Distribution Points found in certificate:\n")
		for i, cdp := range cert.CRLDistributionPoints {
			fmt.Printf("  %d: %s\n", i+1, cdp)
		}
	} else {
		fmt.Println("No CRL Distribution Points found in this certificate")
	}
}

// > % go run main.go
//   === CRL 1: http://crl.globalsign.com/gs/gsorganizationvalsha2g2.crl ===
//   Downloading CRL from: http://crl.globalsign.com/gs/gsorganizationvalsha2g2.crl
//   Downloaded 874 bytes
//   Issuer: CN=GlobalSign Organization Validation CA - SHA256 - G2,O=GlobalSign nv-sa,C=BE
//   This Update: 2024-02-19T10:12:22Z
//   Next Update: 2024-02-26T10:12:21Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 11
//   First 5 revoked certificates:
//     Serial: 18662672422538974321277289152, Revocation Date: 2023-05-09T07:27:02Z
//     Serial: 35514016643895662193939454532, Revocation Date: 2023-10-13T05:45:03Z
//       Reason Code: 0a0104
//     Serial: 35978575104368358173618825245, Revocation Date: 2023-06-01T13:18:03Z
//     Serial: 16574199337270349960530077189, Revocation Date: 2023-08-02T09:18:15Z
//       Reason Code: 0a0104
//     Serial: 35124528866426131862753176072, Revocation Date: 2023-06-05T07:24:02Z

//   === CRL 2: http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl ===
//   Downloading CRL from: http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl
//   Downloaded 1077 bytes
//   Issuer: CN=DigiCert Assured ID Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
//   This Update: 2025-09-03T19:45:53Z
//   Next Update: 2025-09-24T19:45:53Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 12
//   First 5 revoked certificates:
//     Serial: 9601343444604230544260526143014183304, Revocation Date: 2024-08-13T19:18:56Z
//       Reason Code: 0a0104
//     Serial: 16235496577730144984962991517001198931, Revocation Date: 2018-10-09T15:00:17Z
//       Reason Code: 0a0105
//     Serial: 8655459947270566071807919960567640049, Revocation Date: 2020-08-12T12:33:09Z
//       Reason Code: 0a0105
//     Serial: 15246232542431631281698814536494261398, Revocation Date: 2022-10-13T09:58:07Z
//       Reason Code: 0a0105
//     Serial: 14842300446175341861835746176489778372, Revocation Date: 2022-10-13T10:00:34Z
//       Reason Code: 0a0105

//   === CRL 3: http://crl.comodoca.com/COMODORSAOrganizationValidationSecureServerCA.crl ===
//   Downloading CRL from: http://crl.comodoca.com/COMODORSAOrganizationValidationSecureServerCA.crl
//   Downloaded 3241748 bytes
//   Issuer: CN=COMODO RSA Organization Validation Secure Server CA,O=COMODO CA Limited,L=Salford,ST=Greater 
//   Manchester,C=GB
//   This Update: 2025-09-05T22:42:44Z
//   Next Update: 2025-09-12T22:42:44Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 75121
//   First 5 revoked certificates:
//     Serial: 28069386214006847952415371988138086528, Revocation Date: 2022-07-18T17:33:37Z
//     Serial: 63370405348114244513525874389227055670, Revocation Date: 2022-07-19T05:18:42Z
//     Serial: 327194611981445173749294912595784456225, Revocation Date: 2024-08-03T18:26:55Z
//     Serial: 258417543155096435663389000337562465978, Revocation Date: 2024-08-07T15:44:45Z
//       Reason Code: 0a0103
//     Serial: 231565231474474454134325696148896739907, Revocation Date: 2024-08-10T18:28:56Z

//   === CRL Certificate Check Demo ===
//   This demonstrates how to check if a certificate is revoked using CRL
//   Failed to parse certificate: x509: malformed certificate
//   kanehiroyuu@kanehiroyuunoMacBook-Pro CRL % go run main.go
//   === CRL 1: http://crl.globalsign.com/gs/gsorganizationvalsha2g2.crl ===
//   Downloading CRL from: http://crl.globalsign.com/gs/gsorganizationvalsha2g2.crl
//   Downloaded 874 bytes
//   Issuer: CN=GlobalSign Organization Validation CA - SHA256 - G2,O=GlobalSign nv-sa,C=BE
//   This Update: 2024-02-19T10:12:22Z
//   Next Update: 2024-02-26T10:12:21Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 11
//   First 5 revoked certificates:
//     Serial: 18662672422538974321277289152, Revocation Date: 2023-05-09T07:27:02Z
//     Serial: 35514016643895662193939454532, Revocation Date: 2023-10-13T05:45:03Z
//       Reason Code: 0a0104
//     Serial: 35978575104368358173618825245, Revocation Date: 2023-06-01T13:18:03Z
//     Serial: 16574199337270349960530077189, Revocation Date: 2023-08-02T09:18:15Z
//       Reason Code: 0a0104
//     Serial: 35124528866426131862753176072, Revocation Date: 2023-06-05T07:24:02Z

//   === CRL 2: http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl ===
//   Downloading CRL from: http://crl3.digicert.com/DigiCertAssuredIDRootCA.crl
//   Downloaded 1077 bytes
//   Issuer: CN=DigiCert Assured ID Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
//   This Update: 2025-09-03T19:45:53Z
//   Next Update: 2025-09-24T19:45:53Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 12
//   First 5 revoked certificates:
//     Serial: 9601343444604230544260526143014183304, Revocation Date: 2024-08-13T19:18:56Z
//       Reason Code: 0a0104
//     Serial: 16235496577730144984962991517001198931, Revocation Date: 2018-10-09T15:00:17Z
//       Reason Code: 0a0105
//     Serial: 8655459947270566071807919960567640049, Revocation Date: 2020-08-12T12:33:09Z
//       Reason Code: 0a0105
//     Serial: 15246232542431631281698814536494261398, Revocation Date: 2022-10-13T09:58:07Z
//       Reason Code: 0a0105
//     Serial: 14842300446175341861835746176489778372, Revocation Date: 2022-10-13T10:00:34Z
//       Reason Code: 0a0105

//   === CRL 3: http://crl.comodoca.com/COMODORSAOrganizationValidationSecureServerCA.crl ===
//   Downloading CRL from: http://crl.comodoca.com/COMODORSAOrganizationValidationSecureServerCA.crl
//   Downloaded 3241748 bytes
//   Issuer: CN=COMODO RSA Organization Validation Secure Server CA,O=COMODO CA Limited,L=Salford,ST=Greater 
//   Manchester,C=GB
//   This Update: 2025-09-05T22:42:44Z
//   Next Update: 2025-09-12T22:42:44Z
//   Signature Algorithm: 1.2.840.113549.1.1.11
//   Revoked Certificates Count: 75121
//   First 5 revoked certificates:
//     Serial: 28069386214006847952415371988138086528, Revocation Date: 2022-07-18T17:33:37Z
//     Serial: 63370405348114244513525874389227055670, Revocation Date: 2022-07-19T05:18:42Z
//     Serial: 327194611981445173749294912595784456225, Revocation Date: 2024-08-03T18:26:55Z
//     Serial: 258417543155096435663389000337562465978, Revocation Date: 2024-08-07T15:44:45Z
//       Reason Code: 0a0103
//     Serial: 231565231474474454134325696148896739907, Revocation Date: 2024-08-10T18:28:56Z

//   === CRL Certificate Check Demo ===
//   This demonstrates how to check if a certificate is revoked using CRL
//   Example Certificate:
//     Subject: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
//     Serial Number: 4835703278459707669005204
//     Issuer: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE

//   To check if this certificate is revoked:
//   1. Find the CRL Distribution Points in the certificate
//   2. Download the CRL from those URLs
//   3. Check if the certificate's serial number is in the revoked list
//   No CRL Distribution Points found in this certificate

// ⏺ 🔍 CRL実行結果の詳細解説

//   📊 取得したCRLの比較

//   | 認証局        | 失効証明書数  | CRLサイズ   | 更新頻度 |
//   |------------|---------|----------|------|
//   | GlobalSign | 11件     | 874バイト   | 週1回  |
//   | DigiCert   | 12件     | 1,077バイト | 3週間  |
//   | COMODO     | 75,121件 | 3.2MB    | 週1回  |

//   🚨 失効理由コードの意味

//   0a0104 = keyCompromise (鍵の危殆化)
//   0a0105 = cessationOfOperation (運用停止)
//   0a0103 = affiliationChanged (所属変更)

//   鍵の危殆化が多い理由:
//   - 秘密鍵の漏洩
//   - サーバーへの不正侵入
//   - 内部者による悪用

//   📈 COMODOが75,000件も多い理由

//   1. 大手認証局: 世界中の多くのサイトに証明書を発行
//   2. 安価な証明書: 個人・中小企業も多く利用
//   3. セキュリティ事件: 過去に大規模な証明書漏洩事件が発生

//   🔐 証明書チェックの流れ

//   1. 証明書を受信
//      ↓
//   2. CRL Distribution Points を確認
//      ↓
//   3. CRLをダウンロード
//      ↓
//   4. 証明書のシリアル番号を検索
//      ↓
//   5. 見つかれば失効、なければ有効

//   ⚡ CRLの問題点

//   サイズ問題: COMODOは3.2MBもある
//   リアルタイム性: 週1回更新では遅い
//   可用性: CRLサーバーがダウンすると検証不可

//   🆚 OCSPとの違い

//   - CRL: 全失効証明書リストを一括ダウンロード
//   - OCSP: 1つの証明書だけリアルタイム確認

//   現在はOCSPが主流ですが、CRLも併用されています。