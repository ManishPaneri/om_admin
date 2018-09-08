package models

import (
	"encoding/csv"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

// Get all the user details
func GetDetails(from int, to int, sortField string, sortBy string, params map[string]string, baseQuery string) (userDetails []orm.Params, err error) {
	o := orm.NewOrm()
	if baseQuery == "" {
		return
	}
	havingClauseQuery := " "
	BETWEENClauseQuery := ""
	if len(params) > 0 {
		havingClauseQuery += " HAVING "
		for k, v := range params {
			arr := strings.Split(v, ",")
			if len(arr) > 1 && arr[1] != "" {
				BETWEENClauseQuery += k + " BETWEEN '" + arr[0] + "' AND  DATE_ADD('" + arr[1] + "', INTERVAL 1 DAY) AND "
			} else {
				v = strings.Replace(v, ",", "", 1)
				if k == "created_at" || k == "disbursedOn" {
					havingClauseQuery += "DATE(" + k + ") = '" + v + "' AND "
					continue
				}
				havingClauseQuery += k + " LIKE '" + v + "%' AND "
			}

		}
	}
	havingClauseQuery = havingClauseQuery + BETWEENClauseQuery
	havingClauseQuery = strings.TrimRight(havingClauseQuery, "AND ")

	// groupQuery := " GROUP BY u.id "

	sortLimitQuery := " "
	if from >= 0 && to > 0 {
		if sortBy == "" || strings.ToUpper(sortBy) == "ASC" {
			if sortField != "" {
				sortLimitQuery += "ORDER BY " + sortField + " ASC LIMIT " + cast.ToString(to) + " OFFSET " + cast.ToString(from)
			} else {
				sortLimitQuery += " LIMIT " + cast.ToString(to) + " OFFSET " + cast.ToString(from)
			}
		} else {
			if sortField != "" {
				sortLimitQuery += "ORDER BY " + sortField + " DESC LIMIT " + cast.ToString(to) + " OFFSET " + cast.ToString(from)
			} else {
				sortLimitQuery += " LIMIT " + cast.ToString(to) + " OFFSET " + cast.ToString(from)
			}
		}
	} else if from >= 0 && to == 0 {
		if sortBy == "" || strings.ToUpper(sortBy) == "ASC" {
			if sortField != "" {
				sortLimitQuery += "ORDER BY " + sortField + " ASC LIMIT " + cast.ToString(from)
			} else {
				sortLimitQuery += " LIMIT " + cast.ToString(from)
			}
		} else {
			if sortField != "" {
				sortLimitQuery += "ORDER BY " + sortField + " DESC LIMIT " + cast.ToString(from)
			} else {
				sortLimitQuery += " LIMIT " + cast.ToString(from)
			}
		}
	} else {
		fmt.Println("Error in getting the limits")
	}

	query := baseQuery + havingClauseQuery + sortLimitQuery

	_, err = o.Raw(query).Values(&userDetails)
	if err != nil {
		fmt.Println("Error in getting all the user details :", err)
	}

	return userDetails, err
}

func GetTotalNumberRecord(baseQuery string, querarams map[string]string) (maps []orm.Params, err error) {

	totalQuery := "Select COUNT(*) as Total FROM (" + baseQuery + ") a where true"
	BETWEENClauseQuery := ""
	for k, v := range querarams {
		arr := strings.Split(v, ",")
		if len(arr) > 1 && arr[1] != "" {
			BETWEENClauseQuery += " AND " + k + " BETWEEN '" + arr[0] + "' AND  DATE_ADD('" + arr[1] + "', INTERVAL 1 DAY)"
		} else {
			v = strings.Replace(v, ",", "", 1)
			totalQuery = totalQuery + ` and ` + k + ` like"` + v + `%"`

		}
	}
	totalQuery = totalQuery + BETWEENClauseQuery
	o := orm.NewOrm()
	_, err = o.Raw(totalQuery).Values(&maps)
	fmt.Println("Total Number Record: ", maps[0]["Total"])
	if err != nil {
		return nil, err
	}
	return maps, nil
}

func GetTotalNumberRecordForCSV(baseQuery string, querarams map[string]string) (maps []orm.Params, err error) {
	o := orm.NewOrm()

	recordQuery := "Select * FROM (" + baseQuery + ") a where true"
	for k, v := range querarams {
		recordQuery = recordQuery + ` and ` + k + ` like"` + v + `%"`
	}
	_, err = o.Raw(recordQuery).Values(&maps)
	if err != nil {
		return nil, err
	}
	return maps, nil
}

func CSVExport(data []orm.Params, record []string, w http.ResponseWriter) {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	w.Header().Set("Content-Te", "text/csv")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Disposition", "attachment;filename=Export.csv")
	writer.Write(record)
	for _, value := range data {
		dataArray := []string{}
		for _, v := range record {
			dataArray = append(dataArray, cast.ToString(value[v]))
		}
		writer.Write(dataArray)
	}

}

func GetTotalNumberRecordNEW(baseQuery string, querarams map[string]string, searchAll bool) (maps []orm.Params, err error) {
	totalQuery := ""
	if searchAll == true {
		totalQuery = "Select COUNT(*) as Total FROM (" + baseQuery + ") a where true"
		BETWEENClauseQuery := ""
		for k, v := range querarams {
			arr := strings.Split(v, ",")
			if len(arr) > 1 && arr[1] != "" {
				BETWEENClauseQuery += " AND " + k + " BETWEEN '" + arr[0] + "' AND  DATE_ADD('" + arr[1] + "', INTERVAL 1 DAY)"
			} else {
				v = strings.Replace(v, ",", "", 1)
				totalQuery = totalQuery + ` and ` + k + ` like"` + v + `%"`

			}
		}
		totalQuery = totalQuery + BETWEENClauseQuery

	} else {
		totalQuery = baseQuery
	}

	o := orm.NewOrm()
	_, err = o.Raw(totalQuery).Values(&maps)
	if err != nil {
		return nil, err
	}
	return maps, nil
}
