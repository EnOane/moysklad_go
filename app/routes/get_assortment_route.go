package routes

import (
	"encoding/json"
	"moysklad/interacts"
	"net/http"
)

func GetAssortmentRoute(w http.ResponseWriter, r *http.Request) {
	data, err := interacts.GetAssortmentInteract()
	if err != nil {
		return
	}

	encoder := json.NewEncoder(w)

	err = encoder.Encode(data)
	if err != nil {
		return
	}
}
