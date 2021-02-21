package jsonanalysis

import "github.com/buger/jsonparser"

//JsonToInt 解析json制定key的int值
func JsonToInt(data []byte, parameter []string) (val int64, err error) {
	return jsonparser.GetInt(data, parameter...)
}

func JsonToString(data []byte, parameter []string) (val string, err error) {
	return jsonparser.GetString(data, parameter...)
}

func JsonTo(data []byte, parameter []string) (value []byte, dataType jsonparser.ValueType, offset int, err error) {
	return jsonparser.Get(data, parameter...)
}
