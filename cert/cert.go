package cert

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	gopkcs12 "software.sslmate.com/src/go-pkcs12"
)

// PemCertToP12 将 pem 证书转换为 PKCS12 格式
// 传入参数为 pem 证书的内容与私钥内容
func PemCertToP12(certBuf, keyBuf []byte) (p12Cert string, err error) {
	caBlock, certInput := pem.Decode(certBuf)
	if caBlock == nil {
		err = fmt.Errorf("the pem certificate format is incorrect, value(%v)", string(certInput))
		return
	}

	crt, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("the pem certificate fails to resolve, error(%v)", err)
		return
	}

	keyBlock, keyInput := pem.Decode(keyBuf)
	if keyBlock == nil {
		err = fmt.Errorf("the pem certificate key format is incorrect. value(%v)", string(keyInput))
		return
	}

	priKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("the pem certificate key fails to resolve, error(%v)", err)
		return
	}

	pfx, err := gopkcs12.Encode(rand.Reader, priKey, crt, nil, "")
	if err != nil {
		err = fmt.Errorf("certificate conversion exception, error(%v)", err)
		return
	}

	return base64.StdEncoding.EncodeToString(pfx), err
}
