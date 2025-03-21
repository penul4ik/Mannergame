package v1

import "easyhttp/internal"

type User struct {
	Group   string   `json:"group"`
	Name    string   `json:"name"`
	Acl     string   `json:"access_list,omitempty"`
	Secrets *Secrets `json:"secrets,omitempty"`
}

type Secrets struct {
	Vault      string `json:"vault,omitempty"`
	Kubernetes string `json:"kubernetes,omitempty"`
}

func remapUserToInternalUser(u *User) *internal.User {
	return &internal.User{
		Name:  u.Name,
		Group: u.Group,
		Acl:   u.Acl,
	}
}

func remapInternalUserToUser(u *internal.User) *User {
	return &User{
		Name:  u.Name,
		Group: u.Group,
		Acl:   u.Acl,
	}
}
