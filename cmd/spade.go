package cmd

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// parseVal read template file and return map data
func parseVal() (map[string]interface{}, error) {
	vFile, err := ioutil.ReadFile(valFile)
	if err != nil {
		return nil, errors.Wrapf(err, "read value file %s error", valFile)
	}

	var val = make(map[string]interface{})
	if err := yaml.Unmarshal(vFile, val); err != nil {
		return nil, errors.Wrap(err, "unmarshal YAML error")
	}

	return val, nil
}

// renderTpl render template and print to stdout
func renderTpl(val map[string]interface{}) {
	temp, err := template.ParseFiles(tplFile)
	if err != nil {
		panic(err)
	}

	if err := temp.Execute(os.Stdout, val); err != nil {
		panic(err)
	}
}
