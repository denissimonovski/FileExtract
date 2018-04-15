package main

import (
	"os"
	"bufio"
	"time"
	"encoding/csv"
	"fmt"
	"io"
)

func main() {
	start_date := time.Date(2017, 10, 16, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(2017, 10, 22, 0, 0, 0, 0, time.UTC)
	envlozi := make([]string, 2, 2)
	envlozi[0] = "envlog.csv"
	envlozi[1] = "envlog(1).csv"
	fajl, _ := os.Open("data_log.txt")
	defer fajl.Close()
	citac := bufio.NewScanner(fajl)

	ss1, _ := os.Create("ss1.csv")
	ss1.WriteString("Date,TempSS1\n")
	defer ss1.Close()

	ss24, _ := os.Create("ss24.csv")
	ss24.WriteString("Date,TempSS24\n")
	defer ss24.Close()

	hum, _ := os.Create("humidity.csv")
	hum.WriteString("Date,Humidity\n")
	defer hum.Close()

	pla, _ := os.Create("plafon.csv")
	pla.WriteString("Date,Plafon\n")
	defer pla.Close()

	ss2, _ := os.Create("novss2.csv")
	ss2.WriteString("Date,Temperature,Humidity\n")
	defer ss2.Close()

	oh, _ := os.Create("novoh.csv")
	oh.WriteString("Date,Temperature,Humidity\n")
	defer oh.Close()

	for citac.Scan() {
		vreme, _ := time.Parse("02-01-2006\t15:04:05", citac.Text()[:19])
		if vreme.After(start_date) && vreme.Before(end_date) {
			prv_del := vreme.Format("2006-01-02 15:04:05")
			if citac.Text()[len(citac.Text())-3:] == "SS1" {
				ss1.WriteString(prv_del + `,` + citac.Text()[32:36] + "\n")
			}
			if citac.Text()[len(citac.Text())-3:] == "4x7" {
				ss24.WriteString(prv_del + `,` + citac.Text()[32:36] + "\n")
			}
			if citac.Text()[len(citac.Text())-3:] == "ity" {
				hum.WriteString(prv_del + `,` + citac.Text()[29:33] + "\n")
			}
			if citac.Text()[len(citac.Text())-3:] == "fon" {
				pla.WriteString(prv_del + `,` + citac.Text()[29:33] + "\n")
			}
		}
	}

	for i := 0; i < 2; i++ {
		Env, _ := os.Open(envlozi[i])
		Env.Seek(30, 0)
		csvreader := csv.NewReader(bufio.NewReader(Env))
		if i == 0 {
			csvreader.Comma = ','
		} else {
			csvreader.Comma = '\t'
		}
		for {
			linija, err := csvreader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			vreme, _ := time.Parse("01/02/2006 15:04:05", linija[0]+" "+linija[1])
			if vreme.After(start_date) && vreme.Before(end_date) {
				if i == 0 {
					ss2.WriteString(vreme.Format("2006-01-02 15:04:05") +
						"," + linija[2] +
						"," + linija[3] + "\n")
				} else {
					oh.WriteString(vreme.Format("2006-01-02 15:04:05") +
						"," + linija[2] +
						"," + linija[3] + "\n")
				}
			}
		}
	}

}
