package app

import (
	"fmt"
	"moysklad/app/routes"
	"net/http"
)

func CreateApp() {
	http.HandleFunc("/get-assortments", routes.GetAssortmentRoute)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully start server!")
}
