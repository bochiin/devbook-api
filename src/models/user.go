package models

import (
	"api/src/utils"
	"errors"
	"strings"
	"time"
)

type User struct {
	Id        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdIn"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()

	return nil
}

func (usuario *User) validate() error {

	if utils.IsBlank(usuario.Name) {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}

	if utils.IsBlank(usuario.Nickname) {
		return errors.New("O nick é obrigatório e não pode estar em branco.")
	}

	if utils.IsBlank(usuario.Email) {
		return errors.New("O email é obrigatório e não pode estar em branco.")
	}

	if utils.IsBlank(usuario.Password) {
		return errors.New("A senha é obrigatória e não pode estar em branco.")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
