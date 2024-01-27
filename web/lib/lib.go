// lib contains helper functions.
package lib

import (
	"encoding/json"
	"time"

	"github.com/ItsNotGoodName/go-web-app-example/web"
	"github.com/dustin/go-humanize"
)

func HumanizeBytes(bytes int64) string {
	return humanize.Bytes(uint64(bytes))
}

func HumanizeTime(date time.Time) string {
	return humanize.Time(date)
}

func PrettyJSON(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

func FormatTime(meta web.Meta, date time.Time) string {
	return date.In(meta.TimeZone).Format(time.DateTime)
}
