package df

import (
	"strings"
	"bytes"
)

// First parameter has to be a string
func MakeParameters(key string, value string) []byte {
	return []byte(`{"` + key + `":"` + value + `"}`)
}

//func AppendParameter(params []byte, key string, value string) []byte {
//	split := strings.Split(string(params), "}")
//
//	var sb bytes.Buffer
//
//	sb.WriteString(split[0])
//	sb.WriteString(`, "`)
//	sb.WriteString(key)
//
//	if _, err := strconv.ParseFloat(value, 64); err == nil {
//		sb.WriteString(`":`)
//		sb.WriteString(value)
//		sb.WriteString(`}`)
//	} else if strings.HasPrefix(value, "[") {
//		sb.WriteString(`":`)
//		sb.WriteString(value)
//		sb.WriteString(`}`)
//	} else {
//		sb.WriteString(`":"`)
//		sb.WriteString(value)
//		sb.WriteString(`"}`)
//	}
//
//	return []byte(sb.String())
//}

func AppendString(params []byte, key string, value string) []byte {
	split := strings.Split(string(params), "}")

	var sb bytes.Buffer

	sb.WriteString(split[0])
	sb.WriteString(`, "`)
	sb.WriteString(key)
	sb.WriteString(`":"`)
	sb.WriteString(value)
	sb.WriteString(`"}`)

	return []byte(sb.String())
}

func AppendNonString(params []byte, key string, value string) []byte {
	split := strings.Split(string(params), "}")

	var sb bytes.Buffer

	sb.WriteString(split[0])
	sb.WriteString(`, "`)
	sb.WriteString(key)
	sb.WriteString(`":`)
	sb.WriteString(value)
	sb.WriteString(`}`)

	return []byte(sb.String())
}
