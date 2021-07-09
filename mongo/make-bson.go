package mongo

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

/*
	功能：提取结构体中指定的bson tag字段
	输入:
		data: 任意结构体
		tag: 需要提取的tag字段，如果tag为空则全量提取
	输出：
		指定tag组成的bson.M
*/
func MakeBson(data interface{}, tag []string) bson.M {
	mapTags := make(map[string]bool)
	for _, v := range tag {
		mapTags[v] = true
	}

	output := bson.M{}

	var fn func(obj interface{})
	fn = func(obj interface{}) {
		if reflect.TypeOf(obj).Kind() == reflect.Ptr {
			if reflect.ValueOf(obj).IsNil() {
				return
			}

			fn(reflect.ValueOf(obj).Elem().Interface())
			return
		}

		if reflect.TypeOf(obj).Kind() != reflect.Struct {
			return
		}

		numField := reflect.ValueOf(obj).NumField()
		for i := 0; i < numField; i++ {
			bsonStr := reflect.TypeOf(obj).Field(i).Tag.Get("bson")
			if bsonStr == ",inline" {
				fn(reflect.ValueOf(obj).Field(i).Interface())
				continue
			}

			bsonStr = strings.Split(bsonStr, ",")[0]
			if len(mapTags) > 0 {
				if _, ok := mapTags[bsonStr]; ok {
					output[bsonStr] = reflect.ValueOf(obj).Field(i).Interface()
				}
			} else {
				output[bsonStr] = reflect.ValueOf(obj).Field(i).Interface()
			}
		}
	}

	fn(data)

	return output
}
