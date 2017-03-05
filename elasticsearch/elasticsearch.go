package elasticsearch

import (
	"errors"
	"os"

	"github.com/mattbaird/elastigo/lib"
)

var (
	conn *elastigo.Conn
)

//Save - saves something into Elastic search
func Save(index, _type, id string, data interface{}) (resp elastigo.BaseResponse, err error) {
	if err = connect(); err != nil {
		return
	}

	var exists bool
	if exists, err = conn.ExistsBool(index, _type, id, nil); err != nil {
		return
	}

	if exists {
		//TODO: idk why this doesn't work as expected
		//The work around is just to delete it, then re-add it
		if resp, err = conn.Delete(index, _type, id, nil); err != nil {
			return
		}
		resp, err = conn.Index(index, _type, id, nil, data)
		//resp, err = conn.Update(index, _type, id, nil, data)
	} else {
		resp, err = conn.Index(index, _type, id, nil, data)
	}

	return
}

//Delete - Deletes something inside of elastic search
func Delete(index, _type, id string) (resp elastigo.BaseResponse, err error) {
	if err = connect(); err != nil {
		return
	}

	var exists bool
	if exists, err = conn.ExistsBool(index, _type, id, nil); err != nil {
		return
	}

	if exists {
		resp, err = conn.Delete(index, _type, id, nil)
	}

	return
}

//Search - search for something inside of elastic search
func Search(query string, fields []string) (result elastigo.SearchResult, err error) {
	if err = connect(); err != nil {
		return
	}

	var args map[string]interface{}

	qry := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":  query,
				"fields": fields,
			},
		},
		"highlight": map[string]interface{}{
			"fields": map[string]interface{}{
				"*": map[string]interface{}{},
			},
		},
	}

	result, err = conn.Search("*", "", args, qry)

	return
}

//connect - attempts to connect to a given elastic search ip
func connect() (err error) {
	if conn == nil {
		if host := os.Getenv("ELASTICSEARCH_IP"); host != "" {
			conn = &elastigo.Conn{
				Protocol: elastigo.DefaultProtocol,
				Domain:   host,
				Port:     os.Getenv("ELASTICSEARCH_PORT"),
				Username: os.Getenv("ELASTICSEARCH_USER"),
				Password: os.Getenv("ELASTICSEARCH_PASS"),
			}
		}
		if conn == nil {
			return errors.New("Unable to connect to elasticsearch host")
		}
	}
	return nil
}
