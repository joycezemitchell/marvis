package sqlm

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type SQLM interface {
	Get(ctx context.Context, tbl string, obj ...interface{}) ([]map[string]interface{}, error)
}

type sqlm struct {
	db *sql.DB
}

func NewSQLM(db *sql.DB) SQLM {
	return &sqlm{db: db}
}

func (s *sqlm) Get(ctx context.Context, tbl string, obj ...interface{}) ([]map[string]interface{}, error) {
	var sqlFilter string
	for _, x := range obj {
		sqlFilter, _ = s.QueryFilter(x)
	}

	sql := fmt.Sprintf("SELECT * FROM %v WHERE 1=1 %v ", tbl, sqlFilter)

	// fmt.Println("sql", sql)

	rows, err := s.Query(ctx, sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rows, nil
}

func (s *sqlm) QueryFilter(obj interface{}) (string, error) {
	if obj == nil {
		return "", nil
	}

	o := map[string]interface{}{}
	z, _ := json.Marshal(obj)
	json.Unmarshal(z, &o)

	var filterJoin []string

	for k, v := range o {
		switch c := v.(type) {
		case int:
			if c != 0 {
				filterJoin = append(filterJoin, fmt.Sprintf(" AND %s = %v ", k, c))
			}
		case string:
			if c != "" {
				filterJoin = append(filterJoin, fmt.Sprintf(" AND %s = '%s' ", k, c))
			}
		}
	}

	return fmt.Sprintf("%v", strings.Join(filterJoin, "")), nil
}

func (s *sqlm) Query(ctx context.Context, sql string) ([]map[string]interface{}, error) {
	stmt, err := s.db.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	columns, _ := rows.Columns()

	// for each database row / record, a map with the column names and row values is added to the allMaps slice
	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i, _ := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		resultMap := make(map[string]interface{})
		for i, val := range values {
			switch val.(type) {
			case []uint8:
				resultMap[columns[i]] = fmt.Sprintf("%s", val)
			case int64:
				resultMap[columns[i]] = val.(int64)
			}

		}
		allMaps = append(allMaps, resultMap)
	}

	return allMaps, nil
}
