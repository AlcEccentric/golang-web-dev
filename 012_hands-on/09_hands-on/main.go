package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

type FileInfo struct {
	Date     time.Time
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("csv.html"))
}

func ymdFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

var fm = template.FuncMap{
	"ymd": ymdFormat,
}

func main() {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	data := make([]FileInfo, 0, len(lines)-1)

	f := new(FileInfo)
	for i, line := range lines {
		if i == 0 {
			continue
		}

		for j := 0; j < len(line); j++ {

			switch j {
			case 0:
				layout := "2006-01-02"
				f.Date, err = time.Parse(layout, line[0])
				if err != nil {
					log.Fatalln(err.Error())
					return
				}
			case 1:
				f.Open = line[1]
			case 2:
				f.High = line[2]
			case 3:
				f.Low = line[3]
			case 4:
				f.Close = line[4]
			case 5:
				f.Volume = line[5]
			case 6:
				f.AdjClose = line[6]
			}

		}

		data = append(data, *f)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "csv.html", data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
