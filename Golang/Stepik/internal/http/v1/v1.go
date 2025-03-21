package v1

import (
	"easyhttp/internal/ldap"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterRoutes(ldapSvc *ldap.Service) {
	h := newHandler(ldapSvc)

	http.HandleFunc("POST /user", h.helloPost)

	http.HandleFunc("GET /user/{group}", h.helloGet)

	http.HandleFunc("/headers", headers)
}

var htmlTemplate = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>
<p>My first paragraph.</p>

</body>
</html>
`

func newHandler(ldapSvc *ldap.Service) *Handler {
	return &Handler{
		LDAPService: ldapSvc,
	}
}

type Handler struct {
	LDAPService *ldap.Service
}

func (h *Handler) helloGet(w http.ResponseWriter, req *http.Request) {
	// internal server error hello world
	// w.WriteHeader(http.StatusInternalServerError)
	// fmt.Fprintf(w, "hello GET\n")

	// render http template
	// w.Write([]byte(htmlTemplate))

	name := req.URL.Query().Get("name")
	group := req.PathValue("group")

	// LDAP
	// userInfo :=ldapConnector.Request(name,group)
	users := h.LDAPService.ListLdapInfo(group)
	u := User{}
	for _, v := range users {
		if v.Name == name && v.Group == group {
			u = *remapInternalUserToUser(&v)
			break
		}
	}

	if u.Name != name && u.Group != group {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handler) helloPost(w http.ResponseWriter, req *http.Request) {
	u := User{}
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.LDAPService.AddLdapInfo(remapUserToInternalUser(&u))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
