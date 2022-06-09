package format

import "time"

func GormDateFormat(currentTime time.Time) string {
	return currentTime.Format("2006-01-02 15:04")
}
