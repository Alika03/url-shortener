package bootstrap

import (
	"log"
	"shortener/config"
	"shortener/infrastructure/db/postgres"
	"shortener/utils"
)

func InitConfig() {
	config.GetConfig()
}

func Migrate(dir string) {
	dbConn := postgres.GetDbConnection()

	files, err := utils.GetFiles(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, fileName := range files {
		queryString := utils.ReadFile(dir + fileName)

		_, err = dbConn.Exec(queryString)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
