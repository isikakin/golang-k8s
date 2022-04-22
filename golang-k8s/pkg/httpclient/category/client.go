package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	GetCategoryById(id string) (*GetCategoryByIdResponse, error)
}

type client struct {
	httpClient http.Client
	url        string
}

func (self *client) GetCategoryById(id string) (*GetCategoryByIdResponse, error) {

	var (
		responseByte     []byte
		categoryResponse *GetCategoryByIdResponse
		err              error
	)

	response, err := self.httpClient.Get(fmt.Sprintf("%s/%s", self.url, id))

	fmt.Println("url", self.url)

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error while get category err:%s", err.Error()))
	}

	if responseByte, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(responseByte, &categoryResponse); err != nil {
		return nil, errors.New(fmt.Sprintf("Category response parse error. err :%s", err))
	}

	return categoryResponse, nil
}

func NewClient(url string, duration time.Duration) Client {
	return &client{
		httpClient: http.Client{Timeout: duration},
		url:        url,
	}
}
