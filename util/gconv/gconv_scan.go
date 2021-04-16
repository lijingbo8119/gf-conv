// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/lijingbo8119/gf-conv.

package gconv

import (
	"fmt"
	"reflect"
)

// Scan automatically calls MapToMap, MapToMaps, Struct or Structs function according to
// the type of parameter `pointer` to implement the converting.
// It calls function MapToMap if `pointer` is type of *map to do the converting.
// It calls function MapToMaps if `pointer` is type of *[]map/*[]*map to do the converting.
// It calls function Struct if `pointer` is type of *struct/**struct to do the converting.
// It calls function Structs if `pointer` is type of *[]struct/*[]*struct to do the converting.
func Scan(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	var (
		pointerType = reflect.TypeOf(pointer)
		pointerKind = pointerType.Kind()
	)
	if pointerKind != reflect.Ptr {
		return fmt.Errorf("params should be type of pointer, but got: %v", pointerKind)
	}
	var (
		pointerElem     = pointerType.Elem()
		pointerElemKind = pointerElem.Kind()
	)
	switch pointerElemKind {
	case reflect.Map:
		return MapToMap(params, pointer, mapping...)
	case reflect.Array, reflect.Slice:
		var (
			sliceElem     = pointerElem.Elem()
			sliceElemKind = sliceElem.Kind()
		)
		for sliceElemKind == reflect.Ptr {
			sliceElem = sliceElem.Elem()
			sliceElemKind = sliceElem.Kind()
		}
		if sliceElemKind == reflect.Map {
			return MapToMaps(params, pointer, mapping...)
		}
		return Structs(params, pointer, mapping...)
	default:
		return Struct(params, pointer, mapping...)
	}
}

