package HDM

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

const (
	sessionUrl = "/api/session"
)

type HDM struct {
	http      *http.Client
	baseUrl   string
	cSRFToken string
}

type CSR struct {
	PasswordModify uint8  `json:"password_modify"`
	CSRFToken      string `json:"CSRFToken"`
}

func NewHDM(ip, user, pass string) (*HDM, error) {
	var err error
	h := new(HDM)
	jar, _ := cookiejar.New(nil)
	h.http = &http.Client{
		//Timeout: time.Second * 3, // 加超时有几率会无法获取到登陆的CSRtoken
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Jar: jar,
	}

	h.baseUrl = fmt.Sprintf("https://%s", ip)

	body := strings.NewReader(fmt.Sprintf(`username=%s&password=%s`,user,pass))
	req, err := http.NewRequest("POST", h.url(sessionUrl), body)
	if err != nil {
		return h, err
	}
	req.Header.Set("Origin", h.baseUrl)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Referer", h.baseUrl)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")

	resp, err := h.http.Do(req)
	if err != nil {
		return h, err //errors.New("Login Timeout")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return h, err
	}
	var data = &CSR{}
	if err := json.Unmarshal(respBody, data); err != nil {
		return h, err
	}
	//fmt.Println(string(respBody))
	if data.PasswordModify != 0 {
		return h, errors.New("Password Wrong")
	}

	h.cSRFToken = data.CSRFToken
	return h, nil
}


func (h HDM) url(path string) string {
	return fmt.Sprintf("%s%s", h.baseUrl, path)
}

