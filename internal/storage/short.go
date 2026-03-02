package storage

var symbolIndex map[byte]uint64

// Для создания коротких ссылок используется кодировка Base62 + sequence
const symbols = "0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"
const countSymbols = uint64(len(symbols))

func init() {
	symbolIndex = make(map[byte]uint64, len(symbols))
	for i := 0; i < len(symbols); i++ {
		symbolIndex[symbols[i]] = uint64(i)
	}
}

// Encode из uint64 в строку Base62
func Encode(n uint64) string {
	if n == 0 {
		return string(symbols[0])
	}

	var encoded []byte

	for n > 0 {
		position := n % countSymbols
		encoded = append(encoded, symbols[position])
		n = n / countSymbols
	}

	// Переворачиваем слайс
	for i, j := 0, len(encoded)-1; i < j; i, j = i+1, j-1 {
		encoded[i], encoded[j] = encoded[j], encoded[i]
	}

	return string(encoded)
}

// Decode из string в uinit64
func Decode(s string) (uint64, error) {
	var result uint64

	for i := 0; i < len(s); i++ {
		index, ok := symbolIndex[s[i]]
		if !ok {
			return 0, ErrInvalidShort
		}

		result = result*countSymbols + uint64(index)
	}
	return result, nil
}
