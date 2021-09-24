package adbutler

import (
    "encoding/json"
    "net/http"

    "goserver/libs/utils"
)

const (
    END_POINT = "https://api.adbutler.com/v2"
)

type Api struct {
    ApiKey string
}

func (api *Api) SetApiKey(apiKey string) {
    api.ApiKey = apiKey
}

func (api Api) NewApiRequest(service, method string) (*http.Response, error) {
    req, err := http.NewRequest(method, END_POINT + service, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Authorization", "Basic " + api.ApiKey)
    return http.DefaultClient.Do(req)
}

func (api Api) NewApiGetRequest(service string) (*http.Response, error) {
    return api.NewApiRequest(service, "GET")
}

func (api Api) GetZones() (zonesResponse ZonesResponse) {
    resp, err := api.NewApiGetRequest("/zones/standard?limit=100")
    if err != nil {
        return
    }
    body := utils.GetResponseBody(resp)
    if len(body) > 0 {
        json.Unmarshal(body, &zonesResponse)
    }
    return
}

func (api Api) GetPublishers() (publishersResponse PublishersResponse) {
    resp, err := api.NewApiGetRequest("/publishers")
    if err != nil {
        return
    }
    body := utils.GetResponseBody(resp)
    if len(body) > 0 {
        json.Unmarshal(body, &publishersResponse)
    }
    return
}

func (api Api) GetZoneDailyReport() (zoneDailyResponse ZoneDailyReportResponse) {
    resp, err := api.NewApiGetRequest("/reports?preset=last-30-days&&period=day&&type=zone&&details=true")
    if err != nil {
        return
    }
    body := utils.GetResponseBody(resp)
    if len(body) > 0 {
        json.Unmarshal(body, &zoneDailyResponse)
    }
    zoneDailyResponse.CalcPublisherRevShare(1)
    return
}

func (api Api) GetZoneBlankLogs() (blankLogsResponse ZoneBlankLogsResponse) {
    resp, err := api.NewApiGetRequest("/reports/blank-logs?preset=last-30-days&&period=day&&type=zone&&details=true")
    if err != nil {
        return
    }
    body := utils.GetResponseBody(resp)
    if len(body) > 0 {
        json.Unmarshal(body, &blankLogsResponse)
    }
    return
}