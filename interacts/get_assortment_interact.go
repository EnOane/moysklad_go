package interacts

import (
	"moysklad/config"
	"moysklad/core/entites"
	"moysklad/libs"
	"moysklad/repositories"
)

var r = repositories.AssortmentRepository{}

func GetAssortmentInteract() (libs.MoyskladResponse[entites.Assortment], error) {
	dataCh := make(chan libs.MoyskladResponse[entites.Assortment])

	go libs.GetAssortment(libs.MoyskladCredentials{
		Username: config.AppConfig.MoyskladConfig.Username,
		Password: config.AppConfig.MoyskladConfig.Password},
		dataCh)

	data := <-dataCh

	e := data.Rows[0]
	go r.Create(&e)

	return data, nil
}
