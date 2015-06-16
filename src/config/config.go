package config
import (
	"encoding/json"
	"io/ioutil"
	"github.com/ngaut/logging"
	"os"
	"bufio"
	"text/template"
)

type TemplateManager struct {
	Tpl string `json:"tpl"`
	ConfigDatas []ConfigData `json:"config-datas"`
	t *template.Template
}

type ConfigData struct {
	OutFilePath string `json:"out-file-path"`

	//tpl data config
	Data map[string]string `json:"data"`
}

func (t *TemplateManager) LoadConfig(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		logging.Error("Read Config Error:" + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(buf, t)
	if err != nil {
		logging.Error("Parse Config Error:" + err.Error())
		os.Exit(1)
	}

	t.t = template.New(t.Tpl)
}

func (t *TemplateManager) Gen() {
	tpl, err := t.t.ParseFiles(t.Tpl)
	if err != nil {
		logging.Error(err)
		os.Exit(1)
	}

	for _,v := range t.ConfigDatas {
		f, err := os.Create(v.OutFilePath)
		if err != nil {
			logging.Error(err)
			os.Exit(1)
		}
		defer f.Close()

		wr := bufio.NewWriter(f);
		err = tpl.ExecuteTemplate(wr, t.Tpl, v.Data)
		if err != nil {
			logging.Error(err)
			os.Exit(1)
		}
		wr.Flush()
	}
}

func (t *TemplateManager) DebugDump() {
	logging.Debugf("tpl:" + t.Tpl)
	logging.Debugf("count:%d", len(t.ConfigDatas))
	for _, v := range t.ConfigDatas {
		logging.Debugf("=======================")
		logging.Debugf("OutFilePath:" + v.OutFilePath)
		for k, v := range v.Data {
			logging.Debug(k, "->", v)
		}
	}
}