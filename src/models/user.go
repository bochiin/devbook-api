package models

import (
	"api/src/security"
	"api/src/utils"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdIn"`
}

func (user *User) Prepare(step utils.Step) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step utils.Step) error {

	if utils.IsBlank(user.Name) {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if utils.IsBlank(user.Nickname) {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if utils.IsBlank(user.Email) {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if step == utils.CREATED && utils.IsBlank(user.Password) {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(step utils.Step) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if step == utils.CREATED {
		hashedPassword, err := security.Hash(user.Password)

		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
