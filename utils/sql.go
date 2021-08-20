package utils

import "bytes"

type SqlStatement struct {
	Sql string
}

func SqlFormatWhere(where map[string]interface{}) string {
	i := len(where)
	if i == 0 {
		return ""
	}
	keyList := make([]string, 0)
	for i := range where {
		keyList = append(keyList, i)
	}
	var sb bytes.Buffer
	sb.WriteString(" where ")
	for i, e := range keyList {
		if i == 0 {
			sb.WriteString(e)
			sb.WriteString("=:")
			sb.WriteString(e)
		} else {
			sb.WriteString(" and ")
			sb.WriteString(e)
			sb.WriteString("=:")
			sb.WriteString(e)
		}
	}
	s := sb.String()
	return s
}
