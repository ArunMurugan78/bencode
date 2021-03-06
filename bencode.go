// Package bencode is a golang package for bencoding and bdecoding data from and from to equivalents.
package bencode

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

func sortedKeys(mp map[string]interface{}) (keys []string) {
	keys = make([]string, 0)
	for key, _ := range mp {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func encodeInt(i int64) []byte {
	return []byte(fmt.Sprintf("i%ve", i))
}

func encodeString(str string) []byte {
	return []byte(fmt.Sprintf("%v:%v", len(str), str))
}

func encodeDictionary(dict map[string]interface{}) ([]byte, error) {
	encodedDict := []byte("d")
	for _, key := range sortedKeys(dict) {
		val := dict[key]

		encodedDict = append(encodedDict, encodeString(key)...)
		encodedVal, err := Encode(val)
		if err != nil {
			return nil, err
		}

		encodedDict = append(encodedDict, encodedVal...)
	}
	encodedDict = append(encodedDict, byte('e'))

	return encodedDict, nil
}

func encodeList(list []interface{}) ([]byte, error) {
	encodedList := []byte("l")
	for idx := range list {
		encodedVal, err := Encode(list[idx])
		if err != nil {
			return nil, err
		}

		encodedList = append(encodedList, encodedVal...)
	}

	encodedList = append(encodedList, byte('e'))

	return encodedList, nil
}

// Encode encodes the provided data into bencoded bytes if valid
func Encode(data interface{}) ([]byte, error) {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return encodeInt(v.Int()), nil
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return encodeInt(int64(v.Uint())), nil
	case reflect.String:
		return encodeString(v.String()), nil
	case reflect.Slice:
		slice := make([]interface{}, 0)

		for i := 0; i < v.Len(); i++ {
			slice = append(slice, v.Index(i).Interface())
		}

		return encodeList(slice)
	case reflect.Map:
		mp, ok := v.Interface().(map[string]interface{})
		if !ok {
			return nil, errors.New("can't encode " + v.String())
		}
		return encodeDictionary(mp)
	default:
		return nil, errors.New("can't encode " + v.String())
	}
}
