package infra

import (
	"encoding/base64"
	"golang.org/x/net/html/charset"
	"io"
	"mime"
	"regexp"
	"strings"
)

func decodeRFC2047String(str string) (string, error) {

	re := regexp.MustCompile("\\s?=\\?([\\w\\-]+)\\?([A-Za-z])\\?([\\w+/=]+={0,2})\\?=")

	replaceFunc := func(match string, subMatches []string) string {
		s, err := decodeRFC2047stringOne(match, subMatches[1], subMatches[2], subMatches[3])
		if err != nil {
			return ""
		}
		return s
	}

	newStr := re.ReplaceAllStringFunc(str, func(match string) string {
		subMatches := re.FindStringSubmatch(match)
		return replaceFunc(match, subMatches)
	})

	return newStr, nil
}

func decodeRFC2047stringOne(_match string, _charset string, _type string, _encoded string) (string, error) {

	var sDecoded string
	var err error

	if strings.EqualFold(_type, "Q") {
		dec := new(mime.WordDecoder)
		sDecoded, err = dec.DecodeHeader(_match)
		if err != nil {
			return "", err
		}
		return sDecoded, nil
	}

	if strings.EqualFold(_type, "B") {
		bDecoded, err := base64.StdEncoding.DecodeString(_encoded)
		if err != nil {
			return "", err
		}
		sDecoded = string(bDecoded)
	}

	reader := strings.NewReader(sDecoded)
	utf8Reader, err := charset.NewReaderLabel(_charset, reader)
	if err != nil {
		return "", err
	}

	utf8Bytes, err := io.ReadAll(utf8Reader)
	if err != nil {
		return "", err
	}
	return string(utf8Bytes), nil
}
