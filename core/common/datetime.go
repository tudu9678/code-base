package common

import (
	"time"
)

const (
	DateFm            = "2006-01-02"
	TimeFm            = "15:04:05"
	DateTimeFm        = "2006-01-02 15:04:05"
	DateFmInHouse     = "02/01/2006"
	DateTimeFmInHouse = "02/01/2006 15:04:05"
	DateTimeFmSBond   = "20060102"
	DateTimeFmDB      = "2006-01-02T15:04:05Z"
	DateTimeFmTSE     = "02012006"
	YearFm            = "2006"
)

var (
	Weekend                     = map[string]string{"Sunday": "sunday", "Saturday": "saturday"}
	TimeLocation *time.Location = time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)
)

func GetDate(layout string) string {
	return time.Now().Format(layout)
}

func GetTime(layout string) string {
	return time.Now().Format(layout)
}

func FormatFromUnix(location *time.Location, layout string, unix int64) string {
	if unix <= 0 {
		return ""
	}
	return time.Unix(unix, 0).In(location).Format(layout)
}

func ConvertDateToDate(location *time.Location, inputFormat, outputFormat, value string) string {
	res, err := time.ParseInLocation(inputFormat, value, location)
	if err != nil {
		return ""
	}
	return res.Format(outputFormat)
}
