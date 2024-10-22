package utils

import "time"

func StringUtcNow() string {
	return time.Now().In(time.UTC).String()
}

func GetJakartaTimezone() (*time.Location, error) {
	return time.LoadLocation("Asia/Jakarta")
}

func StringDateToStringDate(inputTimezone *time.Location, inputFormat string, input string, outputTimezone *time.Location, outputFormat string) (string, error) {
	inputTime, err := time.ParseInLocation(inputFormat, input, inputTimezone)
	if err != nil {
		return "", err
	}

	localizedPickupTime := inputTime.In(outputTimezone)
	formattedPickupTime := localizedPickupTime.Format(outputFormat)
	return formattedPickupTime, nil
}

func TruncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
