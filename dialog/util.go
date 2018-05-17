package dialog

import (
	"strings"
	"bytes"
	"strconv"
)

const eurPostfix = " Euro"

func PriceInEuroText(p float64) string {

	ps := strconv.FormatFloat(p, 'f', 2, 64)
	nad := strings.Split(ps, ".")

	if nad[1] == "00" {
		return nad[0] + eurPostfix
	}

	return ps + eurPostfix

}

func PriceInEuroTTS(p float64) string {
	ps := strconv.FormatFloat(p, 'f', 2, 64)
	nad := strings.Split(ps, ".")

	if nad[1] == "00" {
		return nad[0] + eurPostfix
	}

	if nad[1] == "99" {
		return nad[0] + eurPostfix + " 99"
	}

	return nad[0] + eurPostfix + " und " + nad[1] + " Cent"
}

func ModifyForTTS(text string) string {
	words := strings.Split(text, " ")

	longWordsToLowerCase(words, 3)
	convertNumbersWithSlash(words)

	tts := concatenate(words)

	//TODO: ampersand-bug in ssml!?
	tts = strings.Replace(tts, "&", " und ", -1)

	return concatenate(words)
}

func convertNumbersWithSlash(words []string) {
	for i, t := range words {
		if strings.Contains(t, "/") {
			pn := strings.Split(t, "/")
			if _, err := strconv.Atoi(pn[0]); err != nil {
				continue
			}
			if _, err := strconv.Atoi(pn[1]); err != nil {
				continue
			}

			words[i] = "<say-as interpret-as='cardinal'>" + pn[0] + "</say-as>" + "<say-as interpret-as='cardinal'>" + pn[1] + "</say-as>"
		}
	}
}

func longWordsToLowerCase(words []string, n int) {
	for i, t := range words {

		if len(t) < n+1 {
			continue
		}

		if t == strings.ToUpper(t) {
			words[i] = strings.ToLower(t)
		}

	}
}

func concatenate(words []string) string {
	var buffer bytes.Buffer

	for i, t := range words {
		buffer.WriteString(t)

		if i < len(words)-1 {
			buffer.WriteString(" ")
		}
	}

	return buffer.String()
}


