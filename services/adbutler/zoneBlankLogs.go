package adbutler


type ZoneBlankLogsResponse struct {
	Data []ZoneBlankLogs `json:"data"`
}

type ZoneBlankLogs struct {
	Type string `json:"type"`
	Id int `json:"id"`
	Summary ZoneDailyBlankLog `json:"summary"`
	Details []ZoneDailyBlankLog `json:"details"`
}

type ZoneDailyBlankLog struct {
	StartDate SpecialDate `json:"start_date,string"`
	Blanks int `json:"blanks"`
}


func (response ZoneBlankLogsResponse) GetZoneBlankLogsById(zoneId int) (logs ZoneBlankLogs) {
	for _, blankLogs := range response.Data {
		if blankLogs.Id == zoneId {
			logs = blankLogs
			break
		}
	}
	return
}

