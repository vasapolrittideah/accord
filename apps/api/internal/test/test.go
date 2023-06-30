package test

import (
	"encoding/json"
	"github.com/vasapolrittideah/accord/apps/api/internal/response"
	"io"
	"net/http"
)

func GetDataFromResponse[T interface{}](resp *http.Response) (*T, error) {
	body, _ := io.ReadAll(resp.Body)
	r := response.Response{}
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return ConvertMapToStruct[T](r.Data)
}

func ConvertMapToStruct[T interface{}](input interface{}) (output *T, err error) {
	jsonString, _ := json.Marshal(input)
	if err := json.Unmarshal(jsonString, &output); err != nil {
		return nil, err
	}

	return
}
