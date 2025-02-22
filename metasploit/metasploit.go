package metasploit

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Metasploit struct {
	host, user, pass, token string
}

func New(host, user, pass string) (*Metasploit, error) {
	mts := &Metasploit{
		host: host,
		pass: pass,
		user: user,
	}

	err := mts.Login()
	if err != nil {
		return nil, err
	}

	log.Println("Logined")

	return mts, nil
}

func (m *Metasploit) send(req, res interface{}) error {

	buf := bytes.Buffer{}

	msgpack.NewEncoder(&buf).Encode(req)
	dest := fmt.Sprintf("%s/api", m.host)

	r, err := http.Post(dest, "binary/message-pack", &buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = msgpack.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return err
	}

	return nil

}

func (m *Metasploit) Login() error {
	logReq := &loginReq{
		Method:   "auth.login",
		Username: m.user,
		Password: m.pass,
	}

	logRes := LoginRes{}

	err := m.send(logReq, &logRes)
	if err != nil {
		return err
	}

	m.token = logRes.Token

	return nil

}

func (m *Metasploit) Logout() error {
	logoutReq := &logoutReq{
		Method:      "auth.logout",
		Token:       m.token,
		LogoutToken: m.token,
	}

	logoutRes := &LoginRes{}

	err := m.send(logoutReq, logoutRes)
	if err != nil {
		return err
	}

	m.token = ""

	return nil

}

func (m *Metasploit) SessionList() (map[uint32]SessionListRes, error) {

	req := &sessionListReq{
		Method: "session.list",
		Token:  m.token,
	}

	res := make(map[uint32]SessionListRes)

	err := m.send(req, res)
	if err != nil {
		return nil, err
	}

	for id, session := range res {
		session.ID = id
		res[id] = session
	}

	return res, nil

}
