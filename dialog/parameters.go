package dialog

import (
	"strings"
	"bytes"
	"strconv"
)

func MakeParameters(key string, value string) []byte{
	return []byte(`{"` + key + `":"` + value + `"}`)
}

func AppendParameter(params []byte, key string, value string) []byte {
	split := strings.Split(string(params), "}")

	var sb bytes.Buffer

	sb.WriteString(split[0])
	sb.WriteString(`, "`)
	sb.WriteString(key)

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		sb.WriteString(`":`)
		sb.WriteString(value)
		sb.WriteString(`}`)
	} else {
		sb.WriteString(`":"`)
		sb.WriteString(value)
		sb.WriteString(`"}`)
	}

	return []byte(sb.String())
}
