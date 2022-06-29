package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

type Text struct {
	Texts []string `json:"texts"`
}

type Responce [][]struct {
	Id   int
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func Replace(responces Responce, texts []string) {
	for i, value := range responces {
		for _, v := range value {
			if len(value) == 0 {
				continue
			}
			texts[i] = strings.Replace(texts[i], v.Word, v.S[0], 1)
		}
	}
}

func GetResponce(texts []string, log *zap.SugaredLogger) (Responce, error) {
	var responces Responce
	for id, val := range texts {
		req := "https://speller.yandex.net/services/spellservice.json/checkTexts"

		u, err := url.Parse(req)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		q := u.Query()

		q.Add("text", val)

		u.RawQuery = q.Encode() //

		resp, err := http.Get(u.String())
		if err != nil {
			log.Error(err)
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		responce := Responce{}
		err = json.Unmarshal(body, &responce)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if len(responce[0]) != 0 {
			responce[0][0].Id = id
		}
		responces = append(responces, responce[0])
	}
	return responces, nil
}

func Tmp(log *zap.SugaredLogger) (Text, error) {
	var text Text
	filename := "templates/text.json"
	result, err := Readfile(filename)
	if err != nil {
		return text, err
	}
	text, err = Unmarshal(result)
	if err != nil {
		return text, err
	}
	return text, err
}

func Readfile(filename string) ([]byte, error) {
	result, err := ioutil.ReadFile(filename)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Unmarshal(result []byte) (Text, error) {
	text := Text{}
	err := json.Unmarshal(result, &text)
	if err != nil {
		return text, err
	}
	return text, nil
}
