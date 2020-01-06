/* +----------------------------+
 * | Author: Zoueature          |
 * +----------------------------+
 * | Email: zoueature@gmail.com |
 * +----------------------------+
 */
package encrypt

import "testing"

func TestRot13(t *testing.T) {
	originStr := "asdqopolj12379*()&(kl;mncbafvx"
	trueEnStr := "nfqdbcbyw12379*()&(xy;zaponsik"
	encryptStr := ""
	for i := 0; i < len(originStr); i ++ {
		enByte := Rot13Encrypt(originStr[i])
		encryptStr += string(enByte)
	}
	if encryptStr != trueEnStr {
		t.Error("encrypt error")
	}
	deStr := ""
	for i := 0; i < len(encryptStr); i ++ {
		deByte := Rot13Decrypt(encryptStr[i])
		deStr += string(deByte)
	}
	if deStr != originStr {
		t.Error("decrypt error")
	}
}
