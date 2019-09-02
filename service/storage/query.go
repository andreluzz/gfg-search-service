package storage

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type body struct {
	Query *query                  `json:"query,omitempty"`
	Sort  []map[string]sortFields `json:"sort,omitempty"`
	From  int                     `json:"from"`
	Size  int                     `json:"size"`
}
type must struct {
	QueryString *queryString      `json:"query_string,omitempty"`
	Match       map[string]string `json:"match,omitempty"`
}
type queryString struct {
	Query string `json:"query"`
}
type queryBool struct {
	Must []must `json:"must"`
}
type query struct {
	Bool queryBool `json:"bool"`
}
type sortFields struct {
	Order string `json:"order"`
}

func generateSearchBody(qs, filter, sort, page, limit string) ([]byte, error) {
	b := body{}
	if qs != "" || filter != "" {
		b.Query = generateQuery(qs, filter)
	}
	if sort != "" {
		b.Sort = generateSort(sort)
	}

	// parse limit string to size int
	size := 20
	if limit != "" {
		s, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("invalid limit number: %v", limit)
		}
		size = s
	}
	b.Size = size

	// parse page string to from int
	from := 0
	if page != "" {
		f, err := strconv.Atoi(page)
		if err != nil {
			return nil, fmt.Errorf("invalid page number: %v", page)
		}
		if f <= 0 {
			f = 1
		}
		from = (f - 1) * size
	}
	b.From = from

	// return struct byte array and error
	return json.MarshalIndent(b, "", "\t")
}

func generateQuery(qs, fs string) *query {
	q := query{}

	if qs != "" {
		qsMust := must{}
		qsMust.QueryString = &queryString{
			Query: qs,
		}

		q.Bool.Must = append(q.Bool.Must, qsMust)
	}

	if fs != "" {
		filters := strings.Split(fs, ",")
		for _, f := range filters {
			tokens := strings.Split(f, ":")
			filterMustMatch := must{}
			filterMustMatch.Match = make(map[string]string)
			filterMustMatch.Match[tokens[0]] = tokens[1]

			q.Bool.Must = append(q.Bool.Must, filterMustMatch)
		}
	}

	return &q
}

func generateSort(sortQueryString string) []map[string]sortFields {
	s := []map[string]sortFields{}

	for _, field := range strings.Split(sortQueryString, ",") {
		fieldMap := make(map[string]sortFields)
		tokens := strings.Split(field, ":")
		fieldMap[tokens[0]] = sortFields{
			Order: tokens[1],
		}
		s = append(s, fieldMap)
	}

	return s
}

// ParseSortTextFields used to validate structs and insert .keyword when the sort field is a text
func ParseSortTextFields(sort string, obj interface{}) string {
	parsedSort := strings.ToLower(sort)
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			fieldJSONName := t.Field(i).Tag.Get("json")
			if index := strings.Index(parsedSort, fieldJSONName); index != -1 {
				parsedSort = parsedSort[:index+len(fieldJSONName)] + ".keyword" + parsedSort[index+len(fieldJSONName):]
			}
		}
	}
	return parsedSort
}
