/* +----------------------------+
 * | Author: Zoueature          |
 * +----------------------------+
 * | Email: zoueature@gmail.com |
 * +----------------------------+
 */
package encrypt

func Rot13Encrypt(b byte) byte {
	var max byte
	if b >= 'a' && b <= 'z' {
		max = 'z'
	} else if b >= 'A' && b <= 'Z' {
		max = 'Z'
	} else {
		return b
	}
	b += 13
	if b > max {
		b -= 26
	}
	return b
}

func Rot13Decrypt(b byte) byte {
	var min byte
	if b >= 'a' && b <= 'z' {
		min = 'a'
	} else if b >= 'A' && b <= 'Z' {
		min = 'A'
	} else {
		return b
	}
	b -= 13
	if b < min {
		b += 26
	}
	return b
}

