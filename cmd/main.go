package main

import (
	"log"
	"net/http"

	"github.com/soncii/todolist-go/pkg/config"
	"github.com/soncii/todolist-go/pkg/controller"
	"github.com/soncii/todolist-go/pkg/router"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := config.Database()
	defer db.Close()
	err = controller.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":"+conf.Port, router.Init())
}
