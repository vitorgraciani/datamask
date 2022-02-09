package mask

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	maskTag = "mask"
)

type maskType string

const (
	MaskInitial     maskType = "initial"
	MaskMiddle      maskType = "middle"
	MaskLast        maskType = "last"
	MaskStruct      maskType = "struct"
	MaskAll         maskType = "all"
	MaskFirstLetter maskType = "firstLetter"
	MaskLastLetter  maskType = "lastLetter"
	MaskEmail       maskType = "email"
)

type Mask struct {
}

func NewMask() *Mask {
	return &Mask{}
}

func (m *Mask) createMask(obj interface{}) (interface{}, error) {
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
		maskApply := elem.Type().Field(i).Tag.Get(maskTag)
		if maskApply == "" {
			ptr.Elem().Field(i).Set(elem.Field(i))
			continue
		}
		switch elem.Field(i).Type().Kind() {
		default:
			ptr.Elem().Field(i).Set(elem.Field(i))
		case reflect.Struct:
			if maskType(maskApply) == MaskStruct {
				_t, err := m.createMask(elem.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())
			}
		case reflect.Ptr:
			if elem.Field(i).IsNil() {
				continue
			}
			if maskType(maskApply) == MaskStruct {
				fmt.Println(elem.Field(i).Interface())
				_t, err := m.createMask(elem.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t))
			}
		case reflect.Interface:
			if elem.Field(i).IsNil() {
				continue
			}
			if maskType(maskApply) != MaskStruct {
				continue
			}
			_t, err := m.createMask(elem.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			if reflect.TypeOf(elem.Field(i).Interface()).Kind() != reflect.Ptr {
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())
			} else {
				ptr.Elem().Field(i).Set(reflect.ValueOf(_t))
			}

		case reflect.String:
			ptr.Elem().Field(i).SetString(m.String(maskType(maskApply), elem.Field(i).String()))
		}
	}

	return ptr.Interface(), nil
}

func (m *Mask) String(tagName maskType, attributeValue string) string {

	switch tagName {
	case MaskInitial:
		return initialData(attributeValue)
	case MaskMiddle:
		return middleData(attributeValue)
	case MaskLast:
		return lastData(attributeValue)
	case MaskAll:
		return allData(attributeValue)
	case MaskFirstLetter:
		return firstLetter(attributeValue)
	case MaskLastLetter:
		return lastLetter(attributeValue)
	case MaskEmail:
		return email(attributeValue)
	}
	return attributeValue
}

func (m *Mask) MaskData(obj interface{}) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occured: ", r)
		}
	}()
	masked, err := m.createMask(obj)
	if err != nil {
		return "", err
	}
	u, _ := json.Marshal(masked)
	return string(u), nil
}
