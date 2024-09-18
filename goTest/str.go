package goTest

import "strings"

// sep 长度不为1情况
func Split(s, sep string) (result []string) {
	// 优化；提前为result分配内存
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep) // 不存在返回-1
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
