package main

import (
	"fmt"
)

type GithubUser struct {
	Name     string
	Surname  string
	GithubID string
}

func (u *GithubUser) GetHelloMessage() string {
	return fmt.Sprintf("I am Github user \nmy name %s, my surname %s, my id %s\n", u.Name, u.Surname, u.GithubID)
}

type YandexUser struct {
	Name     string
	Surname  string
	YandexID string
}

func (u *YandexUser) GetHelloMessage() string {
	return fmt.Sprintf("I am Yandex user \nmy name %s, my surname %s, my id %s\n", u.Name, u.Surname, u.YandexID)
}

type OPGUser struct {
	Name    string
	Surname string
	OPGID   string
}

func (u *OPGUser) GetHelloMessage() string {
	return fmt.Sprintf("I am OPG user \nmy name %s, my surname %s, my id %s\n", u.Name, u.Surname, u.OPGID)
}

type User interface {
	GetHelloMessage() string
}

func main() {

	const (
		name    = "Maks"
		surname = "Sur"
		id      = "12345"
	)

	github := &GithubUser{
		Name:     name,
		Surname:  surname,
		GithubID: id,
	}
	SayHello(github)

	yandex := &YandexUser{
		Name:     name,
		Surname:  surname,
		YandexID: id,
	}

	SayHello(yandex)

	opg := &OPGUser{
		Name:    name,
		Surname: surname,
		OPGID:   id,
	}

	SayHello(opg)
}
