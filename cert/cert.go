package cert

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	gopkcs12 "software.sslmate.com/src/go-pkcs12"
)

// PemCertToP12 将 pem 证书转换为 PKCS12 格式
// 传入参数为 pem 证书的内容与私钥内容
func PemCertToP12(certBuf, keyBuf []byte) (p12Cert []byte, err error) {
	caBlock, certInput := pem.Decode(certBuf)
	if caBlock == nil {
		err = fmt.Errorf("pem 公钥证书格式错误(%v)", string(certInput))
		return
	}

	crt, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("pem 公钥证书解析错误(%v)", err)
		return
	}

	keyBlock, keyInput := pem.Decode(keyBuf)
	if keyBlock == nil {
		err = fmt.Errorf("pem 私钥证书格式错误(%v)", string(keyInput))
		return
	}

	priKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("pem 私钥证书解析错误(%v)", err)
		return
	}

	pfx, err := gopkcs12.Encode(rand.Reader, priKey, crt, nil, "")
	if err != nil {
		err = fmt.Errorf("证书转换错误(%v)", err)
		return
	}
	return pfx, err
}

// ConvertP12Detail 解析 P12 证书内容，获取过期时间与域名
// 传入参数为 P12 证书内容与密码
func ConvertP12Detail(cert []byte, password string) (commonName string, expirationTime time.Time, err error) {
	_, crt, err := gopkcs12.Decode(cert, password)
	if err != nil {
		err = fmt.Errorf("解析 P12 证书内容错误(%v)", err)
		return
	}

	return crt.Subject.CommonName, crt.NotAfter, nil
}
