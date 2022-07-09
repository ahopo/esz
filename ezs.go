// Installation
//	 Terminal
//		go get "github.com/ahopo/ezs"
//	 To your project
// 		import "github.com/ahopo/ezs"
package ezs

import (
	"fmt"
	"reflect"
	"strings"
)

//Usage:
//	Key : The Key the struct
//	Value: The Value or Data inside a struct field
//	Attribute: Second value of the tag
//			e.g ID int `db:"integer, Auto increment primary"` => the "Auto increment primary"
//	TagValue
//			e.g ID int `db:"integer, Auto increment primary"` => the "integer"
// 	DataType
//			e.g ID int `db:"integer, Auto increment primary"` => the int
// 	Tag
//			e.g ID int `db:"integer, Auto increment primary"` => the db
type EZS struct {
	Key       string
	Value     interface{}
	Attribute string
	TagValue  string
	Tag       string
	DataType  string
}

//	Usage:
//		type Person struct {
//			Id   int    `db:"integer, Auto increment primary"`
//			Name string `db:"varchar(255)"`
//			Age  int
//		}
//
//	Parameters
//		_string interface{}
//			- The struct pass as arguments.
//		tagname
//			- The tag declared in the struct
//			  e.g  Id   int    `db:"integer, Auto increment primary"` => db
//
//		func main() {
//			p := new(Person)
//			pdata := ezs.New(Person, "db")
//			fmt.Println(pdata[0].Key)       // => Id
//			fmt.Println(pdata[0].Value)     // => an interface inside the real value
//			fmt.Println(pdata[0].DataType)  // => int
//			fmt.Println(pdata[0].Tag)  		// => db
//			fmt.Println(pdata[0].TagValue)  // => "integer"
//			fmt.Println(pdata[0].Attribute) // => "Auto increment primary"
//		}
//
func New(_struct interface{}) []EZS {
	return _map(_struct)
}

func _map(_struct interface{}) (ezStruct []EZS) {
	val := reflect.ValueOf(_struct).Elem()

	for i := 0; i < val.NumField(); i++ {
		_s := new(EZS)

		key := val.Type().Field(i).Name
		field, _ := val.Type().FieldByName(key)

		_s.Key = key
		_s.Value = val.Field(i).Interface()
		_s.TagValue = ""
		_s.Attribute = ""
		_s.DataType = fmt.Sprintf("%v", field.Type.Kind())
		_s.Tag = strings.Split(string(field.Tag), ":")[0]
		tagdata := strings.Split(field.Tag.Get(_s.Tag), ",")
		if len(tagdata) > 0 {
			_s.TagValue = tagdata[0]
		}
		if len(tagdata) > 1 {
			_s.Attribute = strings.Join(tagdata[1:], " ")
		}
		ezStruct = append(ezStruct, *_s)
	}

	return ezStruct
}
