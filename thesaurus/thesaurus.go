package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

type BigHugh struct {
	APIKey string
}

type Synonyms struct {
	Noun     *words `json:"noun"`
	Verb     *words `json:"verb"`
	Combined *words `json:"combined"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHugh) Synonyms(word string) (syns Synonyms, err error) {
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + word + "/json")
	if err != nil {
		err = errors.New("bighugh: failed looking for synonyms: " + err.Error())
		return
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&syns); err != nil {
		return
	}

	if syns.Noun != nil || syns.Verb != nil {
		syns.Combined = new(words)
		if syns.Noun != nil {
			syns.Combined.Syn = append(syns.Combined.Syn, syns.Noun.Syn...)
		}
		if syns.Verb != nil {
			syns.Combined.Syn = append(syns.Combined.Syn, syns.Verb.Syn...)
		}
	}

	return
}
