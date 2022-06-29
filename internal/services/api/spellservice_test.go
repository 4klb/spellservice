package api

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

var (
	nofile         = ": no such file or directory"
	errReadfile    = "could not read a file"
	errUnmarshal   = "unexpected end of JSON input"
	errGetResponce = "unsupported protocol scheme "
)

func Test_Readfile(t *testing.T) {
	path := "templates/text.json"
	data := map[string]struct {
		filename string
		err      error
	}{
		"correct": {
			filename: path,
		},
		"uncorrect": {
			filename: path,
			err:      errors.New("open " + path + nofile),
		},
	}

	for _, value := range data {
		_, err := Readfile(value.filename)
		if err != nil && value.err == err {
			t.Errorf(errReadfile)
			return
		}
	}
}

func Test_Unmarshal(t *testing.T) {
	data := map[string]struct {
		text []string
		err  error
	}{
		"correct": {
			text: []string{"Время лесных малышей Пришло теплое лето."},
		},
		"uncorrect": {
			text: []string{""},
			err:  errors.New(errUnmarshal),
		},
	}

	for _, value := range data {
		_, err := Unmarshal([]byte(value.text[0]))
		if err != nil && value.err == err {
			t.Errorf(errReadfile)
			return
		}
	}
}

func Test_GetResponce(t *testing.T) {
	data := map[string]struct {
		req  string
		text []string
		err  error
	}{
		"correct": {
			req:  "https://speller.yandex.net/services/spellservice.json/checkTexts",
			text: []string{"Время лесных малышей Пришло теплое лето."},
		},
		"uncorrect": {
			req:  "speller.yandex.net/services/spellservice.json/checkTexts",
			text: []string{"Время лесных малышей Пришло теплое лето."},
			err:  errors.New(errGetResponce),
		},
	}

	for _, value := range data {
		_, err := GetResponce(value.text, &zap.SugaredLogger{})
		if err != nil {
			t.Errorf(errReadfile)
			return
		}
	}
}
