package common

import (
	"strconv"
	"strings"
)

func ParseUriToObj(uri string) string {
	var (
		resources = strings.Split(uri, "/")
		err       error
	)
	for i, v := range resources {
		_, err = strconv.ParseUint(v, 10, 64)
		if err == nil {
			resources[i] = RBAC_URL_VAL_SYMBOL
		}
	}
	return strings.Join(resources, "/")
}
