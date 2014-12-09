package mailmsg

import (
	"reflect"
	"testing"
)

func TestMailFormatter(t *testing.T) {

	imf := New()
	err := imf.AddHeader("To", "bob@example.com")
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	err = imf.AddHeader("From", "sue@example.com")
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	err = imf.AddHeader("Subject", "Hi")
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	err = imf.SetBody([]byte("Hello: Nice to hear from you."))
	if err != nil {
		t.Error("expected err == nil, got", err)
	}

	candidate := imf.Bytes()
	expected := []byte("To:bob@example.com" + crlfs + "From:sue@example.com" + crlfs + "Subject:Hi" + crlfs + "" + crlfs + "Hello: Nice to hear from you.")

	if reflect.DeepEqual(candidate, expected) == false {
		t.Error("expected candidate == expected")
		t.Error(" expected:", string(expected))
		t.Error("candidate:", string(candidate))
		t.Error(" expected:", (expected))
		t.Error("candidate:", (candidate))
	}
}
