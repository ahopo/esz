// Installation
//	 Terminal
//		go get "github.com/ahopo/ezs"
//	 To your project
// 		import "github.com/ahopo/ezs"
package ezs

import (
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
type EZS struct {
	Key       string
	Value     interface{}
	Attribute string
	TagValue  string
	DataType  string
}

//	Usage:
//		type Person struct {
//			Id   int    `db:"integer, Auto increment primary"`
//			Name string `db:"varchar(255)"`
//			Age  int
//		}
//
//		func main() {
//			p := new(Person)
//			pdata := ezs.New(Person, "")
//			fmt.Println(pdata[0].Key)       // => Id
//			fmt.Println(pdata[0].Value)     // => an interface inside the real value
//			fmt.Println(pdata[0].Attribute) // => "Auto increment primary"
//			fmt.Println(pdata[0].TagValue)  // => "integer"
//			fmt.Println(pdata[0].DataType)  // => int
//		}
//
func New(s interface{}, tagname string) []EZS {
	return _map(s, tagname)
}

func _map(_struct interface{}, tag_name string) (ezStruct []EZS) {
	val := reflect.ValueOf(_struct).Elem()

	for i := 0; i < val.NumField(); i++ {
		_s := new(EZS)

		key := val.Type().Field(i).Name
		field, _ := val.Type().FieldByName(key)

		_s.Key = key
		_s.Value = val.Field(i).Interface()
		_s.TagValue = ""
		_s.Attribute = ""

		tagdata := strings.Split(field.Tag.Get(tag_name), ",")
		if len(tagdata) > 1 {
			_s.TagValue = tagdata[0]
		}
		if len(tagdata) > 2 {
			_s.Attribute = strings.Join(tagdata[2:], " ")
		}
		ezStruct = append(ezStruct, *_s)
	}

	return ezStruct
}
