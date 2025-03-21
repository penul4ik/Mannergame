package ldap

import (
	"easyhttp/internal"
	"errors"
	"fmt"
)

type Service struct {
	// group to user list
	Users map[string][]internal.User
}

func NewService() *Service {
	return &Service{
		Users: map[string][]internal.User{
			"sber": []internal.User{
				internal.User{
					Group: "sber",
					Name:  "Myakis",
					Acl:   "root",
				},
				internal.User{
					Group: "sber",
					Name:  "Lena",
					Acl:   "",
				},
			},
			"kiwi": []internal.User{
				internal.User{
					Group: "kiwi",
					Name:  "Max",
					Acl:   "",
				},
				internal.User{
					Group: "kiwi",
					Name:  "Oleg",
					Acl:   "",
				},
			},
		},
	}
}

func (svc *Service) ListLdapInfo(group string) []internal.User {
	users, ok := svc.Users[group]
	if !ok {
		fmt.Println("LDAP group not found")
	}

	return users
}

func (svc *Service) AddLdapInfo(user *internal.User) error {
	if user.Name == "" {
		return errors.New("user has no name")
	}

	if user.Group == "" {
		return errors.New("user has no group")
	}

	users := svc.Users[user.Group]

	svc.Users[user.Group] = append(users, *user)

	return nil
}
