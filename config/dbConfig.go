/*
Copyright © 2026 Matze
*/
package config

import (
	"fmt"
	"os"
	"slices"

	"github.com/mattia37773/mt/functions/env"
)

type DbConfig struct {
	Name      string
	Container string
	User      string
	Password  string
	Shell     string
	Export    string
	Import    string
	Filetype  string
}

func getMysqlDbConfig() DbConfig {
	return DbConfig{
		Name:      "mysql",
		Container: DbContainer(),
		User:      dbUser(),
		Password:  dbPassword(),
		Filetype:  ".sql",
		Shell:     "mysql  -u" + dbUser() + "  -p" + dbPassword(),
		Export:    "mysqldump -u " + dbUser() + " -p" + dbPassword() + " " + DbName() + " > /tmp/" + env.Get("PROJECT_NAME") + ".sql",
		Import:    "mysql -u" + dbUser() + " -p" + dbPassword() + " " + DbName() + " < " + "/tmp/" + env.Get("PROJECT_NAME") + ".sql",
	}
}

func getMongoDbConfig() DbConfig {
	return DbConfig{
		Name:      "mongo",
		Container: DbContainer(),
		User:      dbUser(),
		Password:  dbPassword(),
		Filetype:  ".gzip",
		Shell:     "mongosh  --username " + dbUser() + " --password " + dbPassword(),
		Export:    "mongodump --db=" + DbName() + " --username=" + dbUser() + " --password=" + dbPassword() + " --authenticationDatabase=admin --out=/tmp/" + env.Get("PROJECT_NAME") + ".gzip --gzip",
		Import:    "mongorestore --nsFrom=" + DbName() + ".* --nsTo=" + DbName() + ".* --dir=/tmp/" + env.Get("PROJECT_NAME") + ".gzip --gzip --username=" + dbUser() + " --password=" + dbPassword(),
	}
}

func GetDbConfig() DbConfig {

	var dbType string = DbType()
	allowdTypes := []string{"mysql", "mongo"}
	if !slices.Contains(allowdTypes, dbType) {
		fmt.Printf("\033[31mIError: nvalid Argument %s is not a supported DB. Choose a supported db type %s file\033[0m \n", dbType, allowdTypes)
		os.Exit(1)
	}

	if dbType == "mysql" {
		return getMysqlDbConfig()
	}

	if dbType == "mongo" {
		return getMongoDbConfig()
	}

	os.Exit(1)
	// Will never return because of exit
	// is only here to satisfy the compiler
	return DbConfig{}
}
