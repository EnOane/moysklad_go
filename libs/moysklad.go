package libs

import (
	"fmt"
	"io"
	"moysklad/core/entites"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

type MoyskladCredentials struct {
	Username string
	Password string
}

func request(url url.URL, mc MoyskladCredentials) ([]byte, error) {
	r, err := http.NewRequest("GET",
		url.Path, nil)
	if err != nil {
		fmt.Println(err)
	}

	r.SetBasicAuth(mc.Username, mc.Password)

	body, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(body.Body)

	return bytes, err
}

func GetAssortment(mc MoyskladCredentials, dataCh chan MoyskladResponse[entites.Assortment]) {
	body, err := request(url.URL{Path: "https://api.moysklad.ru/api/remap/1.2/entity/assortment"}, mc)
	if err != nil {
		fmt.Println(err)
	}

	data := Parser[entites.Assortment](&body)
	if err != nil {
	}

	defer close(dataCh)

	dataCh <- data
}
