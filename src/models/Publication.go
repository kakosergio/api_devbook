package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Body       string    `json:"body,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedOn  time.Time `json:"createdOn,omitempty"`
}

// Prepare valida e formata a publicação recebida
func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == ""{
		return errors.New("the title can't be blank")
	}
	if publication.Body == ""{
		return errors.New("the body can't be blank")
	}
	return nil
}

func (publication *Publication) format () {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Body = strings.TrimSpace(publication.Body)
}