package filter

import (
	"reflect"
)

type StringMap map[string][]string
type HandlerFunc func(StringMap)
type HandlerMap map[string]HandlerFunc

var FilterHandler HandlerMap
var CommitFiles []string

func init() {
	// 过滤方法集
	FilterHandler = HandlerMap{
		"Meta": MetaHandler,
	}
	GetCommitFiles()
}

func (h HandlerMap) Run(name string, config StringMap) {
	h[name](config)
}

func Execute() {
	c := GetConfigInfo()
	typeConf := reflect.TypeOf(*c)
	valueConf := reflect.ValueOf(*c)
	for i := 0; i < typeConf.NumField(); i++ {
		tf := typeConf.Field(i)
		vf := valueConf.Field(i)
		config := vf.Interface().(StringMap)
		FilterHandler.Run(tf.Name, config)
	}
}
