package elasticsearch

import (
	"os"
	"testing"

	"github.com/mattbaird/elastigo/lib"
)

func TestElasticSearch(t *testing.T) {
	os.Setenv("ELASTICSEARCH_IP", "127.0.0.1")
	os.Setenv("ELASTICSEARCH_PORT", "9200")
	os.Setenv("ELASTICSEARCH_USER", "")
	os.Setenv("ELASTICSEARCH_PASS", "")

	type Taco struct {
		Name   string
		Type   string
		Weight float64
	}

	type TacoGroup struct {
		Tacos []Taco
	}

	taco := Taco{Name: "El Pasa Encha", Type: "softshell"}
	taco2 := Taco{Name: "Big Beefy Burrito", Type: "burrito"}
	taco3 := Taco{Name: "El Sencha Taco", Type: "hardshell"}

	//TEST INSERT
	_, err := Save("tester", "taco", "test2", taco)
	if err != nil {
		t.Logf("Error saving: %s\n", err)
		t.Fail()
	} else {
		t.Log("Successful Insert - Single")
	}

	//TEST ARRAY INSERT
	specials := TacoGroup{Tacos: []Taco{taco2, taco3}}
	_, err = Save("tester", "specials", "test1", specials)
	if err != nil {
		t.Logf("Error saving: %s\n", err)
		t.Fail()
	} else {
		t.Log("Successful Insert - Array")
	}

	//TEST UPDATE
	taco.Name = "El Poncho Encho"
	_, err = Save("tester", "taco", "test2", taco)
	if err != nil {
		t.Logf("Error updating: %s\n", err)
		t.Fail()
	} else {
		t.Log("Successful Update")
	}

	//TEST DELETE
	_, err = Delete("tester", "specials", "")
	if err != nil {
		t.Logf("Error deleting: %s\n", err)
		t.Fail()
	} else {
		t.Log("Successful Delete")
	}

	//TEST SEARCH
	var res elastigo.SearchResult
	res, err = Search("taco", []string{})
	if err != nil {
		t.Logf("Error searching: %s\n", err)
		t.Fail()
	} else if res.Hits.Len() >= 0 {
		t.Log("Successful Search")
	}
}
