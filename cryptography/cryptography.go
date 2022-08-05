package cryptography

func DeeSeeChiffre(value rune, key int) rune {
	unicodeChar := int(value) + key
	if unicodeChar < 'a' {
		return rune(unicodeChar + 26)
	} else if unicodeChar > 'z' {
		return rune(unicodeChar - 26)
	}
	return rune(unicodeChar)
}
