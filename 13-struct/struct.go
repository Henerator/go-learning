package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name  string `example:"Person Name"`
	Email string
	Birth time.Time
}

func main() {
	createStruct()
	serialize()
	anonymousStruct()
	task()
}

func createStruct() {
	date := time.Now()

	p := Person{}
	p2 := Person{"Name", "Email", date}
	p3 := Person{
		Name:  "Name",
		Email: "Email",
		Birth: date,
	}

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println(p3)
}

func serialize() {
	p := Person{
		Name:  "Name",
		Email: "Email",
	}

	jsP, err := json.Marshal(p)
	if err != nil {
		fmt.Println("unable marshal to json")
	}

	fmt.Println(string(jsP))
}

func anonymousStruct() {
	response := struct {
		FieldName  string `json:"field_name"`
		FieldName2 string `json:"field_name_2"`
	}{
		FieldName:  "Name1",
		FieldName2: "Name2",
	}

	data, _ := json.Marshal(response)
	fmt.Println(string(data))
}

func task() {
	type (
		ResponseHeader struct {
			Code    int    `json:"code"`
			Message string `json:"message,omitempty"`
		}

		ResponseDataItemAttrs struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		}

		ResponseDataItem struct {
			Type       string                `json:"type"`
			Id         int                   `json:"id"`
			Attributes ResponseDataItemAttrs `json:"attributes"`
		}

		ResponseData []ResponseDataItem

		Response struct {
			Header ResponseHeader `json:"header"`
			Data   ResponseData   `json:"data,omitempty"`
		}
	)

	var ReadResponse = func(rawResp string) (Response, error) {
		response := Response{}

		if err := json.Unmarshal([]byte(rawResp), &response); err != nil {
			return Response{}, err
		}

		return response, nil
	}

	const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

	data, err := ReadResponse(rawResp)
	if err != nil {
		fmt.Println("Could not unmarshal")
		return
	}

	fmt.Println(data)
}
