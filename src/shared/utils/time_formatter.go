package utils

import "time"

func FormatTimeToRFC3339(datetime time.Time) (result string) {
	if datetime.IsZero() {
		result = ""
	} else {
		result = datetime.Format(time.RFC3339Nano)
	}
	return
}
