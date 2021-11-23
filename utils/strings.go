/*
Copyright © 2021 cuisongliu@qq.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import "strings"

func NotIn(key string, slice []string) bool {
	for _, s := range slice {
		if key == s {
			return false
		}
	}
	return true
}

func In(key string, slice []string) bool {
	for _, s := range slice {
		if key == s {
			return true
		}
	}
	return false
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SplitLast(str, spilt string) string {
	arr := strings.Split(str, spilt)
	return arr[len(arr)-1]
}
