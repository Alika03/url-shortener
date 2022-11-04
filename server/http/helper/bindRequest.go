package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func BindRequest(r *http.Request, dto interface{}) error {
	if r.Method == http.MethodGet || r.Method == http.MethodDelete {
		return bindGet(r, dto)

	}

	return bindBody(r, dto)
}

func bindGet(r *http.Request, dto interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	temp := map[string]string{}
	for k, v := range r.Form {
		temp[k] = v[0]
	}

	dataJson, err := json.Marshal(temp)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dataJson, dto); err != nil {
		return err
	}

	return nil
}

func bindBody(r *http.Request, dto interface{}) error {
	dataJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dataJson, dto); err != nil {
		return err
	}

	return nil
}

func GetRequestParam(r *http.Request) string {
	return strings.Trim(r.URL.Path, "/")
}
