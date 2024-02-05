package libs

import (
	"encoding/json"
	"fmt"
)

func Parser[TBody any](body *[]byte) MoyskladResponse[TBody] {
	var data MoyskladResponse[TBody]

	err := json.Unmarshal(*body, &data)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
