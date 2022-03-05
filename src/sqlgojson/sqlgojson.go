package sqlgojson

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

func SqlGoJson(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	keyvalueslice := make([][]map[string]interface{}, 0)
	for rows.Next() {
		r := make([]map[string]interface{}, 0)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		for i, col := range columns {
			m := make(map[string]interface{})
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			m[col] = v
			r = append(r, m)
		}
		keyvalueslice = append(keyvalueslice, r)
	}

	jsonData, err := json.Marshal(keyvalueslice)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}
