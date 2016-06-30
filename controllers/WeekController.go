package controllers

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/nickchou/gnydemo/entity"
)

func Week(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	s01 := entity.Staff{JobNum: "06111", Name: "胡彬"}
	s02 := entity.Staff{JobNum: "06216", Name: "周黎"}
	s03 := entity.Staff{JobNum: "06938", Name: "杨卫"}
	s04 := entity.Staff{JobNum: "09905", Name: "王延"}
	s05 := entity.Staff{JobNum: "09911", Name: "季宸"}
	s06 := entity.Staff{JobNum: "10863", Name: "李铖凯"}
	s07 := entity.Staff{JobNum: "16304", Name: "吴旭芬"}
	s08 := entity.Staff{JobNum: "17009", Name: "夏刚"}
	s09 := entity.Staff{JobNum: "21347", Name: "刘若翰"}
	s10 := entity.Staff{JobNum: "26298", Name: "李旦"}
	//slice申明方法1
	ss := []entity.Staff{s01, s02, s03, s04, s05, s06, s07, s08, s09, s10}
	//slice申明方法2
	//ss := make([]entity.Staff, 0)
	//ss = append(ss, s01, s02, s03, s04, s05, s06, s07, s08, s09, s10)
	t := time.Date(2016, 5, 31, 0, 0, 0, 0, time.Local)
	//t := time(2016, 5, 31)
	i := 0
	//table style
	buffer.WriteString("<style type='text/css'>")
	buffer.WriteString("table{font-family:verdana,arial,sans-serif;font-size:11px;color:#333333;border-width:1px;border-color: #666666;border-collapse: collapse;}")
	buffer.WriteString("table th {border-width: 1px;padding: 8px;border-style: solid;border-color: #666666;	background-color: #dedede;}")
	buffer.WriteString("table td {border-width: 1px;padding: 8px;border-style: solid;border-color: #666666;background-color: #ffffff;}")
	buffer.WriteString("</style>")
	//table
	buffer.WriteString("<table>")
	buffer.WriteString("<tr><td>工号</td><td>姓名</td><td>主持周会日期</td></tr>")
	for t.Year() < 2017 {
		if i == len(ss) {
			i = 0
		}
		si := ss[i]
		//fmt.Printf("%s,%s,%s\n", si.JobNum, si.Name, t.Format("2006-01-02"))
		buffer.WriteString("<tr><td>" + si.JobNum + "</td><td>" + si.Name + "</td><td>" + t.Format("2006-01-02") + "</td></tr>")
		i++
		t = t.AddDate(0, 0, 7)
	}
	buffer.WriteString("</table>")
	//response
	io.WriteString(w, buffer.String())
	//io.WriteString(w, strconv.Itoa(len(ss)))
	//io.WriteString(w, "hi  i'm week")
}
