package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Text struct {
	Texts []string `json:"texts"`
}

type Response [][]struct {
	Id   int
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func Replace(responces Response, texts []string) {
	for i, value := range responces {
		for _, v := range value {
			if len(value) == 0 {
				continue
			}
			texts[i] = strings.Replace(texts[i], v.Word, v.S[0], 1)
		}
	}
}

func GetResponce(texts []string, log *zap.SugaredLogger, v *viper.Viper) (Response, error) {
	var responces Response
	errSaver := make(chan error)

	for id, val := range texts {
		go func(id int) { // запросы выполняются конкурентно
			req := v.GetString("api.url")

			u, err := url.Parse(req)
			if err != nil {
				log.Error(err)
				errSaver <- err
				return
			}

			q := u.Query()

			q.Add("text", val)

			u.RawQuery = q.Encode()

			resp, err := http.Get(u.String())
			if err != nil {
				log.Error(err)
				errSaver <- err
				return
			}

			if resp.StatusCode != http.StatusOK { //При вызове удаленных URL проверяются HTTP статус ответа
				log.Error(errors.New("could not to get a response"))
				errSaver <- errors.New("could not to get a response")
				return

			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err)
				errSaver <- err
				return
			}

			defer resp.Body.Close() // закрывается тело ответа после прочтения

			responce := Response{}
			err = json.Unmarshal(body, &responce)
			if err != nil {
				log.Error(err)
				errSaver <- err
				return
			}

			if len(responce[0]) != 0 {
				responce[0][0].Id = id
			}
			responces = append(responces, responce[0])
			errSaver <- nil

		}(id)

		err := <-errSaver
		if err != nil {
			return nil, err
		}
	}
	return responces, nil
}

func GetText(log *zap.SugaredLogger, v *viper.Viper) (Text, error) {
	var text Text
	filename := v.GetString("files.filename")
	result, err := Readfile(filename)
	if err != nil {
		log.Error(err)
		return text, err
	}
	text, err = Unmarshal(result)
	if err != nil {
		log.Error(err)
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
