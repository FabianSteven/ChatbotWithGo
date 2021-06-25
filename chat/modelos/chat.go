package modelos

import (
	"errors"
	"time"
)

const maxLengthInComments = 400

type Chat struct {
	Id      int64
	Comment string
	Date    time.Time
}

type createChat struct {
	Comment string `json:"comment"`
}

func (cmd *createChat) validate() error {

	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("El comentario debe ser menor a 400 caracters")
	}
	return nil
}
