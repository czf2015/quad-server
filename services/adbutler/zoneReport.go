package adbutler

import (
	"time"
)

type ZoneDailyReportResponse struct {
	Data []ZoneReport `json:"data"`
}

type ZoneReport struct {
	Id int `json:"id"`
	Summary ZoneDailyReport `json:"summary"`
	Details []ZoneDailyReport `json:"details"`
}

type ZoneDailyReport struct {
	StartDate time.Time `json:"start_date,string"`
	Responses int `json:"responses"`
	Impressions int `json:"impressions"`
	Clicks int `json:"clicks"`
	Conversions int `json:"conversions"`
	Payout float64 `json:"payout"`
	Blanks int `json:"blanks"`
	Requests int `json:"requests"`
	Cpm float64 `json:"e_cpm"`
	Cpc float64 `json:"e_cpc"`
	Cpa float64 `json:"e_cpa"`
}

func (report *ZoneDailyReport) CalcPublisherRevShare(revShare float64) {
	if report.Responses > 0 {
		report.Requests = report.Responses
	} else {
		report.Requests = report.Impressions
	}
	report.Payout = report.Payout * revShare
	if report.Impressions > 0 {
		report.Cpm = report.Payout / float64(report.Impressions) * 1000
	}
	report.Cpc = 0
	report.Cpa = 0
}

func (response *ZoneDailyReportResponse) CalcPublisherRevShare(revShare float64) {
	for i, _ := range response.Data {
		response.Data[i].Summary.CalcPublisherRevShare(revShare)
		for j, _ := range response.Data[i].Details {
			response.Data[i].Details[j].CalcPublisherRevShare(revShare)
		}
	}
}

func (response *ZoneDailyReportResponse) CombineBlankLogs(blankLogsResponse ZoneBlankLogsResponse) {
	for i, zoneReport := range response.Data {
		blankLogs := blankLogsResponse.GetZoneBlankLogsById(zoneReport.Id)
		if len(blankLogs.Details) > 0 {
			for j, dailyReport := range response.Data[i].Details {
				for _, dailyBlankLog := range blankLogs.Details {
					if dailyReport.StartDate.Equal(time.Time(dailyBlankLog.StartDate)) {
						response.Data[i].Details[j].Blanks = dailyBlankLog.Blanks
						response.Data[i].Details[j].Requests = response.Data[i].Details[j].Requests + dailyBlankLog.Blanks
						response.Data[i].Summary.Blanks = response.Data[i].Summary.Blanks + dailyBlankLog.Blanks
						if response.Data[i].Summary.Requests == 0 {
							response.Data[i].Summary.Requests = response.Data[i].Summary.Responses
						}
						response.Data[i].Summary.Requests = response.Data[i].Summary.Requests + dailyBlankLog.Blanks
					}
				}
			}
		}
	}
}