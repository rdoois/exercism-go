package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var res string
	for _, v := range plain {
		b := int(byte(v))
		i := 0
		if b >= 97 && b <= 122 {
			i = b + shiftKey
			if i > 122 {
				i -= 26
			}
		} else if b >= 65 && b <= 90 {
			i = b + shiftKey
			if i > 90 {
				i -= 26
			}
		}
		if i == 0 {
			res += string(v)
			continue
		}
		res += string(byte(i))
	}

	return res
}
