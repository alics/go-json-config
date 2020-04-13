package main

import (
	"fmt"
	"github.com/alics/go-json-config/jsonconfig"
)

type Caching struct {
	ApplicationKey string
	Host           string
}

type Service struct {
	Id   int
	Name string
}

func main() {
	cnn, err := jsonconfig.GetSection("ConnectionStrings:DbConnection")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(cnn)
	}

	var c Caching
	err = jsonconfig.Bind(&c, "Caching")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(c)
	}

	var s []Service
	err = jsonconfig.Bind(&s, "Services:List")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s)
	}
}
