package main

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
   // Setup the mgm default config
   err := mgm.SetDefaultConfig(nil, "db", options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.e7nzi6y.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		panic(err)
	} else {
		fmt.Println("ok!")
	}
}