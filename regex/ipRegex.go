package regex

import "regexp"

var ipv4Regex *regexp.Regexp
var ipv4MaskRegex *regexp.Regexp

func init(){
	//https://stackoverflow.com/questions/5284147/validating-ipv4-addresses-with-regexp
	ipv4Regex = regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)

	//https://stackoverflow.com/questions/5360768/regular-expression-for-subnet-masking
	ipv4MaskRegex = regexp.MustCompile(`^((255|254|252|248|240|224|192|128).){3}(255|254|252|248|240|224|192|128)$`)
}

func CheckIpv4(value string) bool{
	return ipv4Regex.Match([]byte(value))
}

func CheckIpv4Mask(value string) bool {
	return ipv4MaskRegex.Match([]byte(value))
}
