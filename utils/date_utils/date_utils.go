package date_utils

import "time"

const (
	apiDateLayout = "01-02-2006T15:01:02Z"
)

func GetNow() {
	return time.Now().UTC()
}

func GetNowString() string {
	return now.Format(apiDateLayout) //01 = date, 02 = months

}
