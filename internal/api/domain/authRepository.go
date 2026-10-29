package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/codacube/go-microservice-banking/internal/logger"
)

type AuthRepository interface {
	IsAuthorised(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorised(token string, routeName string, vars map[string]string) bool {
	var response *http.Response
	var err error

	u := buildVerifyURL(token, routeName, vars)

	if response, err = http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	}

	m := map[string]bool{}
	if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
		logger.Error("Error while decoding response from auth server:" + err.Error())
		return false
	}

	return m["isAuthorised"]
}

/*
  This will generate a url for token verification in the below format

  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}

  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
