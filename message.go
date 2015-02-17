package mailmsg

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

const (
	lf = ("\n") // U+000A
	cr = ("\r") // U+000D
)

const crlfs = "\r\n" // string version of crlf
var crlf = []byte(crlfs)

// currently do nothing but naive formatting.
type InternetMessageFormat struct {
	headers [][]byte // header is a raw header line, in the rfc compliant format of <KEY>:<VALUE>
	body    []byte
}

// Create a new, empty message
func New() *InternetMessageFormat {
	return new(InternetMessageFormat)
}

func (imf *InternetMessageFormat) AddHeader(field_name, field_body string) error {
	if imf.headers == nil {
		imf.headers = make([][]byte, 0, 3) // typically have at least 2 or 3 headers,
	}

	// TODO scan for illegal values/characters
	// TODO header value needs escaping
	// TODO support for long header fields

	// strip any trailing whitespace
	field_body = strings.TrimRightFunc(field_body, unicode.IsSpace)

	header_line := fmt.Sprintf("%s:%s", field_name, field_body)
	imf.headers = append(imf.headers, []byte(header_line))

	return nil
}

func (imf *InternetMessageFormat) SetBody(message_body []byte) error {
	// TDDO, escaping, validating, etc..
	imf.body = message_body

	return nil
}

// Format the message according to the appropriate RFC
func (imf *InternetMessageFormat) Bytes() []byte {
	output_lines := make([][]byte, 0, len(imf.headers)+1)

	for _, header := range imf.headers {
		output_lines = append(output_lines, header)
	}

	// an empty line is the lexical element that separates headers from message body
	output_lines = append(output_lines, []byte(""))

	output_lines = append(output_lines, imf.body)

	return bytes.Join(output_lines, crlf)
}
