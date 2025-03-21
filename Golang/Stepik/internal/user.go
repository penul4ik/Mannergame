package internal

type User struct {
	Group   string
	Name    string
	Acl     string
	Secrets *Secrets
}

type Secrets struct {
	Vault      string
	Kubernetes string
}
