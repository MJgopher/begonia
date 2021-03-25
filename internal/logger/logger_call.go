// Code generated by Begonia. DO NOT EDIT.
// versions:
// 	Begonia v1.0.2
// source: begonialog\begonialog.go
// begonia client file

package logger

import (
	"github.com/MashiroC/begonia/app"
	appClient "github.com/MashiroC/begonia/app/client"
	"github.com/MashiroC/begonia/app/coding"
)

var (
	CenterLogServiceService appClient.Service

	_CenterLogServiceServiceSave appClient.RemoteFunSync

	_CenterLogServiceServiceSaveInSchema = `
{
			"namespace":"begonia.func.Save",
			"type":"record",
			"name":"In",
			"fields":[
				{"name":"F1","type":{
				"type": "array",
				"items": {
				"type": "record",
				"name": "Msg",
				"fields":[{"name":"ServerName","type":"string"}
,{"name":"Level","type":"int"}
,{"name":"Fields","type":{"type":"map","values":"string"}}
,{"name":"Time","type":"long"}
,{"name":"Callers","type":{
				"type": "array",
				"items": "string"
			}}

				]
			}
			},"alias":"msg"}

			]
		}`
	_CenterLogServiceServiceSaveOutSchema = `
{
			"namespace":"begonia.func.Save",
			"type":"record",
			"name":"Out",
			"fields":[
				
			]
		}`
	_CenterLogServiceServiceSaveInCoder  coding.Coder
	_CenterLogServiceServiceSaveOutCoder coding.Coder
)

type _CenterLogServiceServiceSaveIn struct {
	F1 []Msg
}

type _CenterLogServiceServiceSaveOut struct {
}

type Msg struct {
	ServerName string            `json:&#34;server_name&#34;` // 服务名
	Level      int               `json:&#34;level&#34;`       // 日志等级
	Fields     map[string]string `json:&#34;log_msg&#34;`     // 日志信息
	Time       int64             `json:&#34;time&#34;`        // 时间
	Callers    []string          `json:&#34;callers&#34;`     // 路径
}

func init() {
	app.ServiceAppMode = app.Ast

	var err error

	_CenterLogServiceServiceSaveInCoder, err = coding.NewAvro(_CenterLogServiceServiceSaveInSchema)
	if err != nil {
		panic(err)
	}
	_CenterLogServiceServiceSaveOutCoder, err = coding.NewAvro(_CenterLogServiceServiceSaveOutSchema)
	if err != nil {
		panic(err)
	}

}

func Save(msg []Msg) (err error) {
	var in _CenterLogServiceServiceSaveIn
	in.F1 = msg

	b, err := _CenterLogServiceServiceSaveInCoder.Encode(in)
	if err != nil {
		panic(err)
	}

	begoniaResTmp, err := _CenterLogServiceServiceSave(b)
	if err != nil {
		return
	}

	var out _CenterLogServiceServiceSaveOut
	err = _CenterLogServiceServiceSaveOutCoder.DecodeIn(begoniaResTmp.([]byte), &out)
	if err != nil {
		panic(err)
	}

	return
}
