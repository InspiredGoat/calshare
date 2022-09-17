package main

// each session has a calendar file with some dates, when updated an html site is generated based on it


import (
	"fmt"
	"os"
	"log"
	"time"
	"strconv"
	"net/http"
)

type Availability int16

const (
	Unspecified Availability = iota
	Available
	Tentative
	Unlikely
	Busy
)

type CalendarFile struct {
	age int
}


func is_leap_year(year int) bool {
	if (year % 400 == 0) {
		return true;
	} else if (year % 100 == 0) {
		return false;
	} else if (year % 4 == 0) {
		return true;
	} else {
		return false;
	}
}

func month_days(month time.Month, year int) int {
	switch (month) {
	case 1:
		return 31;
	case 2:
		if is_leap_year(year) {
			return 29;
		} else {
			return 28;
		}
	case 3:
		return 31;
	case 4:
		return 30;
	case 5:
		return 31;
	case 6:
		return 30;
	case 7:
		return 31;
	case 8:
		return 31;
	case 9:
		return 30;
	case 10:
		return 31;
	case 11:
		return 30;
	case 12:
		return 31;
	default:
		return -1;
	}
}

func calendar_to_html_file(filename string, day_start int, month_start time.Month, year_start int) {
	f, err := os.Create(filename);
	defer f.Close();
	if err != nil {
		log.Fatal(err);
	}


	write := func(line string) {
		fmt.Fprintln(f, line);
	}


	// offset date to start on a monday
	weekday_start := time.Date(year_start, month_start, day_start, 12, 0, 0, 0, time.Now().Location()).Weekday();
	if weekday_start != time.Monday {
		if (weekday_start == time.Sunday) {
			day_start -= 6;
		} else {
			day_start -= (int(weekday_start) - 1)
		}

		if (day_start < 1) {
			month_start--;
			if (month_start < 1) {
				year_start--;
				month_start = 12;
			}
			day_start = month_days(month_start, year_start) + day_start;
		}
	}

	// generate weeks
	day 	:= day_start;
	month 	:= month_start
	year 	:= year_start;

	write("<div id=\"calendar\">");

	write("<div id=\"cal-header\"><div class=\"day\"> Mon </div><div class=\"day\"> Tue </div><div class=\"day\"> Wed </div><div class=\"day\"> Thu </div><div class=\"day\"> Fri </div><div class=\"day\"> Sat </div><div class=\"day\"> Sun </div></div>");

	write("<div id=\"cal-weeks\">");
	for week := 0; week < 5; week++ {
		write("<div class=\"week\">");
		for i := 0; i < 7; i++ {
			write("<div class=\"day\">");

			// write header
			if day == 1 {
				if month == time.January {
					write(strconv.Itoa(year) + "<br> " + month.String() + " " + strconv.Itoa(day) + "st ");
				} else {
					write(month.String() + " " + strconv.Itoa(day) + "st ");
				}
			} else {
				write(strconv.Itoa(day));
			}

			// write events

			// write people available
			write("<br>5 people attending");

			// write

			write("</div>");

			day += 1;
			if (day > month_days(month, year)) {
				day = 1;
				month++;

				if (month > 12) {
					month = time.January;
					year++;
				}
			}
		}
		write("</div>");
	}
	write("</div>");
	write("</div>");
}

func listen(w http.ResponseWriter, r *http.Request) {
	println(r.FormValue("name"));
	println(r.FormValue("location"));
	println(r.FormValue("desc"));
	println(r.FormValue("testname"));

	//fmt.Fprintf(w, "hello\n");
	http.Redirect(w, r, "localhost:8000", http.StatusSeeOther);
}

func main() {
	fmt.Println("Started...");

	calendar_to_html_file("test.html", 2, 10, 2022);

	fmt.Println("Done!");

	http.HandleFunc("/listen", listen);

	http.ListenAndServe(":8888", nil);
}
