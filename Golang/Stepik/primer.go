package main

import (
	"fmt"
	"net/http"
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

func SayHello(u User) {
	msg := u.GetHelloMessage()
	fmt.Println(msg)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// для запросов из vscode удобно использовать thunderclient
	// http2 в действии

	// TODO:
	// 1) определить методы GET и POST
	// 2) прикрутить HTML шаблон
	// 3) Сделать GET запрос на бекенд с query параметрами, считать параметры
	// 4) Сделать sumbit из HTML на бекенд, считать параметры
	// 5) Сделать POST запрос с json телом и обработать его внутри программы( считать в структуру и вывести в stdout)
	// 6) Отдавать корректные статус коды на запросы
	// 7) Пункт со звездочкой - прикрутить middleware, для авторизации (брать header Authorization и получить из него Bearer ***, проверять и выдавать корректный результат)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}

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
