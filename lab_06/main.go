package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type JsonWriter struct {
	Result      string
	Indentation int
	HasFields   bool

	Path []string
}

func (w *JsonWriter) indent() {
	w.Result += strings.Repeat(" ", w.Indentation*4)
}

func (w *JsonWriter) StartArray() {
	w.Result += "["
	w.Indentation += 1
	w.HasFields = false
}

func (w *JsonWriter) EndArray() {
	w.Indentation -= 1
	if w.HasFields {
		w.Result += "\n"
		w.indent()
	}
	w.Result += "]"
}

func (w *JsonWriter) StartObject() {
	w.Result += "{"
	w.Indentation += 1
	w.HasFields = false
}

func (w *JsonWriter) EndObject() {
	w.Indentation -= 1
	if w.HasFields {
		w.Result += "\n"
		w.indent()
	}
	w.Result += "}"
}

func (w *JsonWriter) StartField(name string) {
	if w.HasFields {
		w.Result += ","
	}
	w.Result += "\n"
	w.indent()
	w.Result += fmt.Sprintf(`"%v": `, strings.ReplaceAll(name, `"`, `\"`))

	w.Path = append(w.Path, name)
}

func (w *JsonWriter) EndField() {
	w.HasFields = true

	w.Path = w.Path[0 : len(w.Path)-1]
}

func (w *JsonWriter) StartElement() {
	if w.HasFields {
		w.Result += ","
	}
	w.Result += "\n"
	w.indent()

	w.Path = append(w.Path, "#")
}

func (w *JsonWriter) EndElement() {
	w.HasFields = true

	w.Path = w.Path[0 : len(w.Path)-1]
}

func (w *JsonWriter) Bool(value bool) {
	if value {
		w.Result += "true"
	} else {
		w.Result += "false"
	}
}

func (w *JsonWriter) Number(number string) {
	w.Result += number
}

func (w *JsonWriter) String(str string) {
	w.Result += fmt.Sprintf(`"%v"`, strings.ReplaceAll(str, `"`, `\"`))
}

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
	Test       map[any]any
}

func toJSON(w *JsonWriter, value any) error {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		w.Number(fmt.Sprintf("%v", value))
		return nil
	case string:
		w.String(value.(string))
		return nil
	case bool:
		w.Bool(value.(bool))
		return nil
	default:
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Pointer, reflect.Interface:
			return toJSON(w, v.Elem().Interface())
		case reflect.Array, reflect.Slice:
			w.StartArray()
			for i := range v.Len() {
				w.StartElement()
				if err := toJSON(w, v.Index(i).Interface()); err != nil {
					return err
				}
				w.EndElement()
			}
			w.EndArray()
			return nil
		case reflect.Struct:
			w.StartObject()
			for i := range v.NumField() {
				fieldValue := v.Field(i)
				fieldType := v.Type().Field(i)

				fieldName := fieldType.Name
				tagFieldName := fieldType.Tag.Get("json")
				if tagFieldName != "" {
					fieldName = tagFieldName
				}

				w.StartField(fieldName)
				if err := toJSON(w, fieldValue.Interface()); err != nil {
					return err
				}
				w.EndField()
			}
			w.EndObject()
			return nil
		case reflect.Map:
			w.StartObject()
			keys := v.MapKeys()
			for _, key := range keys {
				keyDeref := key
				if keyDeref.Kind() == reflect.Interface {
					keyDeref = key.Elem()
				}

				if keyDeref.Kind() == reflect.String {
					w.StartField(keyDeref.String())
				} else if keyDeref.CanInt() {
					w.StartField(strconv.FormatInt(keyDeref.Int(), 10))
				} else if keyDeref.CanUint() {
					w.StartField(strconv.FormatUint(keyDeref.Uint(), 10))
				} else {
					return fmt.Errorf("json key can not be of type %v at path '%v'", keyDeref.Type().Name(), strings.Join(w.Path, "."))
				}
				if err := toJSON(w, v.MapIndex(key).Interface()); err != nil {
					return err
				}
				w.EndField()
			}
			w.EndObject()
			return nil
		case reflect.Func:
			return fmt.Errorf("can not convert function to json")
		}
	}
	return fmt.Errorf("unsupported type %T at path '%v'", value, strings.Join(w.Path, "."))
}

func ToJSON(value any) (string, error) {
	w := &JsonWriter{}
	if err := toJSON(w, value); err != nil {
		return "", err
	}
	return w.Result, nil
}

func ToYAML(value any) (string, error) {
	return ToJSON(value)
}

func main() {

	server := &Server{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
		AllowedIPs: []string{
			"192.168.1.1",
			"10.0.0.1",
		},
	}

	json, err := ToJSON(server)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v\n", json)

	yaml, err := ToYAML(server)

	fmt.Printf("%v\n", yaml)
}
