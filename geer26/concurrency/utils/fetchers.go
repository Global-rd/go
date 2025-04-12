package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type BullshitExcuse struct {
	Bs    string `json:"bs"`
	Param string `json:"param"`
	Size  string `json:"size"`
}

type UselessFact struct {
	Id        string `json:"id"`
	Language  string `json:"language"`
	Permalink string `json:"parmalink"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Text      string `json:"text"`
}

type CatFact struct {
	Data []string `json:"data"`
}

func GetCatFact() (string, error) {
	catfactUrl := "https://meowfacts.herokuapp.com/"
	resp, err := http.Get(catfactUrl)
	if err != nil {
		return "", err
	}
	var catfact CatFact
	err = json.NewDecoder(resp.Body).Decode(&catfact)
	if err != nil {
		return "", err
	}
	return catfact.Data[0], nil
}

func GetUselessFact() (string, error) {
	uselessfactUrl := "https://uselessfacts.jsph.pl/api/v2/facts/random?language=en"
	resp, err := http.Get(uselessfactUrl)
	if err != nil {
		return "", err
	}
	var uselessfact UselessFact
	err = json.NewDecoder(resp.Body).Decode(&uselessfact)
	if err != nil {
		return "", err
	}
	return uselessfact.Text, nil
}

func GetBullshitExcuse() (string, error) {
	var randexcuse string
	for _ = range 4 {
		randexcuse += strconv.Itoa(rand.Intn(10000))
	}
	bullshitExcuseUrl := fmt.Sprintf("https://bullshit.sandros.hu/excuse/getbs.php?randy=0.%s", randexcuse)
	resp, err := http.Get(bullshitExcuseUrl)
	if err != nil {
		return "", err
	}
	var bullshitexcuse BullshitExcuse
	err = json.NewDecoder(resp.Body).Decode(&bullshitexcuse)
	if err != nil {
		return "", err
	}
	return bullshitexcuse.Bs, nil
}
