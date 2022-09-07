package goutils

import (
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.Atoi(value); err != nil {
			return fallback
		} else {
			return result
		}
	}
	return fallback
}

func GetFloatEnv(key string, fallback float64) float64 {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseFloat(value, 32); err != nil {
			return fallback
		} else {
			return result
		}
	}
	return fallback
}

func GetBoolEnv(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseBool(value); err != nil {
			return fallback
		} else {
			return result
		}
	}
	return fallback
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func StringsContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			// 型変換可能か確認する
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			// 型変換する
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			// 等価判定をする
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}
