// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/lijingbo8119/gf-conv.

package gconv

import (
	"time"
)

// Time converts `i` to time.Time.
func Time(any interface{}, format ...string) time.Time {
	// It's already this type.
	if len(format) == 0 {
		if v, ok := any.(time.Time); ok {
			return v
		}
	}
	return time.Time{}
}

// Duration converts `i` to time.Duration.
// If `i` is string, then it uses time.ParseDuration to convert it.
// If `i` is numeric, then it converts `i` as nanoseconds.
func Duration(any interface{}) time.Duration {
	// It's already this type.
	if v, ok := any.(time.Duration); ok {
		return v
	}
	return time.Duration(Int64(any))
}
