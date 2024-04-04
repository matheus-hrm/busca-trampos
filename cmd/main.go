package main

import (
	"log"
	"github.com/go-sql-driver/mysql"
	"github.com/matheus-hrm/trampos/cmd/api"
	"github.com/matheus-hrm/trampos/db"
	"github.com/matheus-hrm/trampos/config"
)

func main(){
	db, err := db.NewMySQL(mysql.Config{
		User: 					config.Envs.DBUser,
		Passwd: 				config.Envs.DBPassword,
		Addr: 					config.Envs.DBAdress,
		DBName: 				config.Envs.DBName,
		Net: 	 				"tcp",
		AllowNativePasswords: 	true,
		ParseTime: 			  	true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}