package modelos

import (
	"testing"
)

func NewChat(comment string) *createChat {

	return &createChat{
		Comment: comment,
	}
}

func Test_createChatValidation(t *testing.T) {
	r := NewChat("hola")

	err := r.Validate()

	if err != nil {
		t.Error("No paso la validacion")
		t.Fail()
	}
}
