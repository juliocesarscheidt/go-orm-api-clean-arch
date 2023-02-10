package service

import "time"

func FormatTimeToRFCRFC3339(datetime time.Time) (result string) {
	if datetime.IsZero() {
		result = ""
	} else {
		result = datetime.Format(time.RFC3339Nano)
	}
	return
}
