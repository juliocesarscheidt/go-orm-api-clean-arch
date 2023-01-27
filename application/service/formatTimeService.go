package service

import "time"

func FormatTimeToRFC(datetime time.Time) (result string) {
	if datetime.IsZero() {
		result = ""
	} else {
		result = datetime.Format(time.RFC3339Nano)
	}
	return
}
