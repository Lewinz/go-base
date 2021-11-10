package cert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPemCertUtil(t *testing.T) {
	_, err := PemCertToP12(nil, nil)
	assert.NotNil(t, err)
	err = nil

	_, err = PemCertToP12([]byte("This is an incorrect pem"), []byte("This is an incorrect pem key"))
	assert.NotNil(t, err)

	certBuf := []byte(`
-----BEGIN CERTIFICATE-----
MIIGBzCCBO+gAwIBAgIQB43fqgfzkdLGf7ZbZrnW8DANBgkqhkiG9w0BAQsFADBy
MQswCQYDVQQGEwJDTjElMCMGA1UEChMcVHJ1c3RBc2lhIFRlY2hub2xvZ2llcywg
SW5jLjEdMBsGA1UECxMURG9tYWluIFZhbGlkYXRlZCBTU0wxHTAbBgNVBAMTFFRy
dXN0QXNpYSBUTFMgUlNBIENBMB4XDTIxMDcyMjAwMDAwMFoXDTIyMDcyMTIzNTk1
OVowGTEXMBUGA1UEAxMOenhsLmxld2luei5vcmcwggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQCrZI5SBkY4Ac1jefupHlHSH+dhgyfPr3HP6FejKw+9+yha
dmIT4LXnNro+ROSNsKZY48bw9BjcrpfJV31GKABLKCKZWELbM8VJXXmrhsAMoBrj
xyx6LLRsgyfjD81CrQdxZ2z66aLfLl8eWtb6tsvbh73gmilkfHVit2J1UdrSZLmI
6NRD1nVetiZ60vsvZCMbiU3wahKrVjQMZQIG2tOaWI+VXOop8A45tegoiK49Zfzi
I0PARuyynPZZdKavfrhexZ1zda4CRm8hOO4B4P+3tIeqwm7mkUXQPlvx+1ZXY1+p
sPbPnqX5H6Kg+baw4AQJdI4i02AS+CVMfYVLO/O9AgMBAAGjggLwMIIC7DAfBgNV
HSMEGDAWgBR/05nzoEcOMQBWViKOt8ye3coBijAdBgNVHQ4EFgQUTyPr4yF7M2zW
PxQ1nrnSr1+70jYwGQYDVR0RBBIwEIIOenhsLmxld2luei5vcmcwDgYDVR0PAQH/
BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjA+BgNVHSAENzA1
MDMGBmeBDAECATApMCcGCCsGAQUFBwIBFhtodHRwOi8vd3d3LmRpZ2ljZXJ0LmNv
bS9DUFMwgZIGCCsGAQUFBwEBBIGFMIGCMDQGCCsGAQUFBzABhihodHRwOi8vc3Rh
dHVzZS5kaWdpdGFsY2VydHZhbGlkYXRpb24uY29tMEoGCCsGAQUFBzAChj5odHRw
Oi8vY2FjZXJ0cy5kaWdpdGFsY2VydHZhbGlkYXRpb24uY29tL1RydXN0QXNpYVRM
U1JTQUNBLmNydDAJBgNVHRMEAjAAMIIBfgYKKwYBBAHWeQIEAgSCAW4EggFqAWgA
dgBGpVXrdfqRIDC1oolp9PN9ESxBdL79SbiFq/L8cP5tRwAAAXrOuT5PAAAEAwBH
MEUCIC8TECfLn7Z2QSF50fslk6qRRUD7/M9U/s/G4lkc9YLGAiEAnFhazQEZi0iv
k8pTyYVUpTFxc7KVLWKld+FS0LtP830AdwBRo7D1/QF5nFZtuDd4jwykeswbJ8v3
nohCmg3+1IsF5QAAAXrOuT6FAAAEAwBIMEYCIQDUhLTQ8F26EVZK0WjeiAgB3thd
ZQY7xMdSt78DPO5rtQIhAJxTdyPLScpckqH1TcYrj/YI9UTKVTETqEWMaCmWeczC
AHUAQcjKsd8iRkoQxqE6CUKHXk4xixsD6+tLx2jwkGKWBvYAAAF6zrk+EgAABAMA
RjBEAiAsVp6D1slvCxoZOmJWTN7FtfWhy5eDe9HNtWk2kUN/OAIgVCK7LOpU78QB
du0dfxwhjpYVfzrIpGECTIFhx5IvQ98wDQYJKoZIhvcNAQELBQADggEBAABsCFZT
k6utJLcyl2h5+AerUcNYhJaT+u4iiqQEgMEn6axhO/2LGQg46gHZKWlWUqt5g3yG
2pjNwdV+6Ih/nOT9bckcyT12ZpHmu+MWCqdtR+5ks6tT+XWvnWm0VTZztnv0WhDN
mnu89bPphE1XLNEQTV7m5JeyvY+++HeY5vHyeKEKmUJBo5iGn2Cj/FEawCzoJJCt
3KbxU/5ccTd7hTQAr6XXXJMkQhbbT2Jdi63gFkdSwhkTqbckrrRh7N0kFYJO1Fp4
ZhfYBD+DfzgCzcCffju8/vK6Dtj6M5tvO6dRCpyLYsmPdzpuoYLpkEVa7NXZy/PQ
U0SSRQhXUliRv5k=
-----END CERTIFICATE-----`)

	keyBuf := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAq2SOUgZGOAHNY3n7qR5R0h/nYYMnz69xz+hXoysPvfsoWnZi
E+C15za6PkTkjbCmWOPG8PQY3K6XyVd9RigASygimVhC2zPFSV15q4bADKAa48cs
eiy0bIMn4w/NQq0HcWds+umi3y5fHlrW+rbL24e94JopZHx1YrdidVHa0mS5iOjU
Q9Z1XrYmetL7L2QjG4lN8GoSq1Y0DGUCBtrTmliPlVzqKfAOObXoKIiuPWX84iND
wEbsspz2WXSmr364XsWdc3WuAkZvITjuAeD/t7SHqsJu5pFF0D5b8ftWV2NfqbD2
z56l+R+ioPm2sOAECXSOItNgEvglTH2FSzvzvQIDAQABAoIBAAHC2QEymM6dsCAC
onQeAGO7GbF4u57onoVZzcAQDXZGNNIxz2IFEqwYoHMgL0QIdGZ17WGk9HYct35l
ZH9lIlROCjBOsJ3YK0G3e5b9sw6o0is/LO+9crDUa9jAsmXUvtPqvDswzZBjBYMG
QRlsPu3XeCXtF5n1exmjqO1BL9EuckOrBYMJwzO/bAN3FOreMUfXT0uX/V7sdch/
6lKUhVS8EdKvlRdgtr0JSXkHNJE4p2A0Qta1Fvy4TlmeYhVyTWIHvf9HXqNyzGcw
Y2MeecgDElkQxvozvGkpWXv5iQlr7j9/NB2JUnOaQAGVlXrAK8YQ+KcxI+Ofr8xk
QC7cNmkCgYEA8Rprl1ypGOcG7nigpRdlG14HogNBX/wVl4sx8onwolfkZr2zpADK
ewMbK6Yijo4DFVD+buN08AFJhfNqIwFSTohoUa6EUGOxLFvJIb/sYu+A+eKiz6O3
pkVmsfEJPyJKysC+xDeNzj3lTJmlFdihm8heHWH6C/wmtM7Z65ZmgHkCgYEAtfuC
94C3nPZq6f5sGbtzdVZzhcIHu77/dqnM1fXMivEDM3wMlAMEA6SVynyhAsoyrvp9
GvL1pbXfDffV2ZBwNWB4VvGxX5bgllBt9pQOkF5SZjm9t4kIl6E8WUG457yRoIAb
Lfj/E0DuxXnRGn4XPg/0fe7k5GobD1UaRjB8ZGUCgYEAntMstRURP5pQ8p78FEUJ
EbIrjQpf8n75Kk9Do+ZCYm9LwnKM+CidOdOd/m7+rLHYTh6AvUORMNloOZlT/aNN
OPaa4dP4zYweln4QTO9FJRdo+zPU1LugqyNktyt1T+WjJ1U5VcDS5V3Yw1EjcvS1
4Q1pEioMsgB07v6kh5EYDOECgYEAm1NF8Hxju8wzau8mUzxEitU0GumGcj/Oifja
BZEbeUfG5K2virGcPoO++iovv1LXubPBDjxrYHoAHUr4sw7uRxDFBeia7Sy5GnMh
uEGcwKpRCEGmZT3IIKuU99X5vYmcfnJ5QF7zT/qvEcwspsESk31IwCgkI7VQzWBk
4Z3GvmUCgYB0KmBDGPbB1QLn40DOFM58vpHR0VYW4Ct5mOJNkGDiWxBM5Ecf6gIm
iMnhxJjeZeQGY1RroUhj31idxTz9PhwigZzeiTxn5pcy5EBwzh2a/ix/JLObbcmZ
H/Me/JtAfoZC/KPdFXUAP7sKMiRn4W8yjRbwnDoJ+oILYC7BQnnDrg==
-----END RSA PRIVATE KEY-----`)
	p12Cert, err := PemCertToP12(certBuf, keyBuf)
	assert.Nil(t, err)
	assert.NotNil(t, p12Cert)

	commonName, expirationTime, err := ConvertP12Detail(p12Cert, "")
	assert.Nil(t, err)
	assert.True(t, commonName == "zxl.lewinz.org")
	assert.NotNil(t, expirationTime)

	_, _, err = ConvertP12Detail([]byte("This is an incorrect pkcs12"), "")
	assert.NotNil(t, err)
	err = nil

	_, _, err = ConvertP12Detail([]byte(""), "")
	assert.NotNil(t, err)
	err = nil

	_, _, err = ConvertP12Detail(nil, "")
	assert.NotNil(t, err)
}
