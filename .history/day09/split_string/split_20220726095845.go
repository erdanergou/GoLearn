package split_string

import "strings"

// Split 切割字符串
// abc b -> ac
func Split(str string, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)
	
}
