package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/sinmetal/gcpmetadata"
)

const Service = "dsrunner"

type EnvConfig struct {
	DatastoreProject string
	Goroutine        int `default:"3"`
}

func main() {
	var env EnvConfig
	if err := envconfig.Process("srunner", &env); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("ENV_CONFIG %+v\n", env)

	project, err := gcpmetadata.GetProjectID()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client, err := CreateClient(ctx, project)
	if err != nil {
		panic(err)
	}
	ds := NewDatastoreStore(client)

	endCh := make(chan error, 10)

	goQueryKeysOnly(ds, env.Goroutine, endCh)

	err = <-endCh
	fmt.Printf("BOMB %+v", err)
}
