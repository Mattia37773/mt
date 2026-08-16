/*
Copyright © 2026 Matze
*/
package config

import (
	"fmt"
	"slices"
)

type DbConfig struct {
	Name     string
	Shell    string
	Export   string
	Import   string
	Filetype string
}

func getMysqlDbConfig() DbConfig {
	return DbConfig{
		Name:     "mysql",
		Filetype: ".sql",
		Shell:    "mysql -u" + ProjectConfig.DB.User + " -p" + ProjectConfig.DB.Password,
		Export:   "mysqldump -u" + ProjectConfig.DB.User + " -p" + ProjectConfig.DB.Password + " " + ProjectConfig.DB.Name + " > /tmp/" + ProjectConfig.ProjectName + ".sql",
		Import:   "mysql -u" + ProjectConfig.DB.User + " -p" + ProjectConfig.DB.Password + " " + ProjectConfig.DB.Name + " < " + "/tmp/" + ProjectConfig.ProjectName + ".sql",
	}
}

func getMongoDbConfig() DbConfig {
	return DbConfig{
		Name:     "mongo",
		Filetype: ".gzip",
		Shell:    "mongosh --username " + ProjectConfig.DB.User + " --password " + ProjectConfig.DB.Password,
		Export:   "mongodump --db=" + ProjectConfig.DB.Name + " --username=" + ProjectConfig.DB.User + " --password=" + ProjectConfig.DB.Password + " --authenticationDatabase=admin --out=/tmp/" + ProjectConfig.ProjectName + ".gzip --gzip",
		Import:   "mongorestore --nsFrom=" + ProjectConfig.DB.Name + ".* --nsTo=" + ProjectConfig.DB.Name + ".* --dir=/tmp/" + ProjectConfig.ProjectName + ".gzip --gzip --username=" + ProjectConfig.DB.User + " --password=" + ProjectConfig.DB.Password,
	}
}

func GetDbConfig() (DbConfig, error) {

	var dbType string = ProjectConfig.DB.Provider
	allowdTypes := []string{"mysql", "mongo"}
	if !slices.Contains(allowdTypes, dbType) {
		return DbConfig{}, fmt.Errorf("Invalid Argument %s is not a supported DB. Choose a supported db type %s file\n", dbType, allowdTypes)
	}

	if dbType == "mysql" {
		return getMysqlDbConfig(), nil
	}

	if dbType == "mongo" {
		return getMongoDbConfig(), nil
	}

	return DbConfig{}, fmt.Errorf("Invalid Argument %s is not a supported DB. Choose a supported db type %s file\n", dbType, allowdTypes)
}
