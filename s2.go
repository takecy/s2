package s2

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

// FromMap is map to struct.
func FromMap(m map[string]interface{}, dest interface{}) error {
	if m == nil {
		return errors.New("map blank")
	}
	if dest == nil {
		return errors.New("dest blank")
	}

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return err
	}

	return nil
}

// ToMap is struct to map
func ToMap(tagName string, src interface{}) (map[string]interface{}, error) {
	if src == nil {
		return nil, errors.New("src blank")
	}

	srcVal := reflect.Indirect(reflect.ValueOf(src))
	srcType := srcVal.Type()

	m := make(map[string]interface{})

	for i := 0; i < srcType.NumField(); i++ {
		typeField := srcType.Field(i)
		valField := srcVal.Field(i)

		var mKey string
		if tagName == "" {
			mKey = typeField.Name
		} else {
			mKey = typeField.Tag.Get(tagName)
			if mKey == "" {
				mKey = typeField.Name
			} else {
				// `json:"hoge,omitempty"`
				sname := strings.Split(mKey, ",")
				mKey = sname[0]

				// `json:",string"`
				if mKey == "" {
					mKey = typeField.Name
				}
			}
		}

		if valField.Kind() == reflect.Struct {
			// recursion
			mm, _ := ToMap(tagName, valField.Interface())
			m[mKey] = mm
			continue
		}

		m[mKey] = valField.Interface()
	}

	return m, nil
}
