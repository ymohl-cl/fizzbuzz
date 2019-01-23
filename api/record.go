package api

import (
	"database/sql"
)

func (a API) Record(data InputFB) error {
	var err error
	var id string

	id = data.Hash()
	sqlStatement := `INSERT INTO stats
				(id, limit, int1, int2, str1, str2, count)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				ON DUPLICATE KEY UPDATE count=count+1`
	if _, err = a.driver.Exec(sqlStatement, id, data.Limit, data.Int1, data.Int2, data.Str1, data.Str2, 1); err != nil {
		return err
	}
	return nil
}

func (a API) MaxRecord() ([]OutputStat, error) {
	var ret []OutputStat
	var err error
	var rows *sql.Rows

	sqlStatement := `SELECT limit, int1, int2, str1, str2, count
				FROM stats
				WHERE count=MAX(count)`
	if rows, err = a.driver.Query(sqlStatement); err != nil {
		return nil, err
	}

	for rows.Next() {
		var out OutputStat
		if err = rows.Scan(&out.Limit, &out.Int1, &out.Int2, &out.Str1, &out.Str2, &out.Count); err != nil {
			return nil, err
		}
		ret = append(ret, out)
	}
	return ret, nil
}
