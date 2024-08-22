package utils

import (
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
)

type (
	RunFunc func(*cobra.Command, []string) error
)

func MultiRun(functions ...RunFunc) RunFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, fn := range functions {
			if err := fn(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}

func GetQueryParam(r *http.Request, name, defaultValue string) string {
	values, ok := r.URL.Query()[name]
	if !ok || len(values) == 0 {
		return defaultValue
	}
	return values[0]
}

func GetQueryParamInt(r *http.Request, name string, defaultValue int) int {
	value := GetQueryParam(r, name, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}
