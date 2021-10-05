package adbutler

import (
	"encoding/json"
	"strings"
	"time"
)

type SpecialDate time.Time

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	utc, _ := time.LoadLocation("America/New_York")
	newTime, err := time.Parse("2006-01-02 MST", strInput+" EDT")
	if err != nil {
		return err
	}

	*sd = SpecialDate(newTime.In(utc))
	return nil
}

func (sd SpecialDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(sd))
}
