package utils

var alphabet = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphabetLen = uint(len(alphabet))

const maxEncodedLen = 11 // int(math.Ceil(math.Log10(math.MaxUint) / math.Log10(float64(alphabetLen))))

func NumberToShortCode(num uint) string {

	buf := make([]byte, maxEncodedLen)
	pos := maxEncodedLen

	for {
		index := num % alphabetLen
		pos--
		buf[pos] = alphabet[index]
		num /= alphabetLen
		if num == 0 {
			return string(buf[pos:])
		}
	}
}
