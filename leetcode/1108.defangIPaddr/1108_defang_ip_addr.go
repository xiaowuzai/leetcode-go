package leetcode

import "strings"

func defangIPaddr1(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func defangIPaddr2(address string) string {
	builder := strings.Builder{}

	n := len(address)
	for i := 0; i < n; i++ {
		if address[i] == '.' {
			builder.WriteString("[.]")
			continue
		}
		builder.WriteByte(address[i])
	}
	return builder.String()
}
