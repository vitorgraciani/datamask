package mask

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type showType string

const showTag = "show"

const (
	ShowInitial     showType = "initial"
	ShowMiddle      showType = "middle"
	ShowLast        showType = "last"
	ShowStruct      showType = "struct"
	ShowAll         showType = "all"
	ShowFirstLetter showType = "firstLetter"
	ShowLastLetter  showType = "lastLetter"
	ShowEmail       showType = "email"
)

type Show struct {
}

func (m *Show) String(tagName showType, attributeValue string) string {

	switch tagName {
	case ShowLast:
		return initialData(attributeValue)
	case ShowMiddle:
		return showMiddleData(attributeValue)
	case ShowInitial:
		return lastData(attributeValue)
	case ShowAll:
		return attributeValue
	case ShowFirstLetter:
		return lastLetter(attributeValue)
	case ShowLastLetter:
		return firstLetter(attributeValue)
	case ShowEmail:
		return email(attributeValue)
	}
	return allData(attributeValue)
}

func (s *Show) ShowData(obj interface{}) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occured: ", r)
		}
	}()
	sh, err := s.createShowData(obj)
	if err != nil {
		return "", err
	}
	u, _ := json.Marshal(sh)
	return string(u), nil
}

func (s *Show) createShowData(obj interface{}) (interface{}, error) {
	if obj == nil {
		return obj, nil
	}
	typeObj := reflect.TypeOf(obj)
	var ptr, elem reflect.Value

	if typeObj.Kind() == reflect.Ptr {
		ptr = reflect.New(typeObj.Elem())
		elem = reflect.ValueOf(obj).Elem()
	} else {
		ptr = reflect.New(typeObj)
		elem = reflect.ValueOf(obj)
	}
	for i := 0; i < elem.NumField(); i++ {
		ShowApply := elem.Type().Field(i).Tag.Get(showTag)
		switch elem.Field(i).Type().Kind() {
		case reflect.Struct:
			_t, err := s.createShowData(elem.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			ptr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())

		case reflect.Ptr:
			if elem.Field(i).IsNil() {
				continue
			}
			fmt.Println(elem.Field(i).Interface())
			_t, err := s.createShowData(elem.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			ptr.Elem().Field(i).Set(reflect.ValueOf(_t))

		case reflect.Interface:
			if elem.Field(i).IsNil() {
				continue
			}
			_t, err := s.createShowData(elem.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			if reflect.TypeOf(elem.Field(i).Interface()).Kind() != reflect.Ptr {
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())
			} else {
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t))
			}

		case reflect.String:
			if ShowApply == "" {
				ptr.Elem().Field(i).SetString(allData(elem.Field(i).String()))
				continue
			} else {
				ptr.Elem().Field(i).SetString(s.String(showType(ShowApply), elem.Field(i).String()))
			}
		default:
			ptr.Elem().Field(i).Set(elem.Field(i))
		}
	}

	return ptr.Interface(), nil
}

func NewShow() *Show {
	return &Show{}
}
