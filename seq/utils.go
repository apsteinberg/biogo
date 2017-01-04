package seq

func Reverse(s []byte) []byte {
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		bs = append(bs, s[i])
	}

	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return bs
}

func Complement(s []byte) []byte {
	m := make(map[byte]byte)
	m['A'] = 'T'
	m['T'] = 'A'
	m['G'] = 'C'
	m['C'] = 'G'
	bs := []byte{}
	for _, a := range s {
		b := a
		switch a {
		case 'A':
			b = 'T'
			break
		case 'a':
			b = 't'
			break
		case 'T':
			b = 'A'
			break
		case 't':
			b = 'a'
			break
		case 'C':
			b = 'G'
			break
		case 'c':
			b = 'g'
			break
		case 'G':
			b = 'C'
			break
		case 'g':
			b = 'c'
			break
		}
		bs = append(bs, b)
	}

	return bs
}
