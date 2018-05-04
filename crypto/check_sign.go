package crypto

import ()

// BanerwaiAPIV1CheckSign banerwai api sign check
func BanerwaiAPIV1CheckSign(sign, apiKey string, args ...string) bool {
	if len(sign) == 0 {
		return false
	}
	total := apiKey
	for _, arg := range args {
		total += arg
	}
	return CompareDoubleMd5(total, sign)
}

// BanerwaiAPIV1GenSign banerwai api gen sign
func BanerwaiAPIV1GenSign(apiKey string, args ...string) string {
	total := apiKey
	for _, arg := range args {
		total += arg
	}
	return DoubleMd5(total)
}
