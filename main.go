package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	yaml2Sql("./ta.yaml")
}

func yaml2Sql(yamlFile string) {
	var receiveYamlDatas map[string]string

	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Print(err)
		return
	}

	err = yaml.Unmarshal(data, &receiveYamlDatas)
	if err != nil {
		log.Print(err)
		return
	}

	sqlFile, err := os.Create("taUpdate.sql")
	if err != nil {
		log.Print(err)
		return
	}

	for key, value := range receiveYamlDatas {
		query := fmt.Sprintf(`UPDATE translations set value = value || '{"ta": "%s"}' where key = '%s';%s`, value, key, "\n")
		n, err := sqlFile.WriteString(query)
		if err != nil {
			log.Print(err)
			return
		}
		if n != len(query) {
			log.Print(err)
			return
		}
	}
	err = sqlFile.Close()
	if err != nil {
		log.Print(err)
		return
	}
}
