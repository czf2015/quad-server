package apiv1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"

	"goserver/libs/conf"
	"goserver/services/adbutler"
	"goserver/libs/jwt"
	"goserver/models"
)

var apiKey = conf.GetSectionKey("app", "ADBUTLER_API_KEY").String()

func GetLast30DaysReportApi(c *gin.Context) {
    token := c.Query("token")
    claims, _ := jwt.ParseToken(token)
    user := models.GetUserById(claims.Id)
    userDomains := funk.Map(user.GetApprovedDomains(), func (approved models.ApprovedDomain) string {
        return approved.Domain
    }).([]string)

	api := adbutler.Api{ApiKey: apiKey}

    publishers := api.GetPublishers()
    var myPubs []adbutler.Publisher
    for _, p := range publishers.Data {
        if funk.Contains(userDomains, p.Name) {
            myPubs = append(myPubs, p)
        }
    }
	zones := api.GetZones()
    var publisherZones []adbutler.Zone
    var zoneids []int
    for _, publisher := range myPubs {
        publisherZones = append(publisherZones, publisher.GetZones(zones.Data)...)
        zoneids = append(zoneids, publisher.GetZoneIds(zones.Data)...)
    }
    reports := api.GetZoneDailyReport()
    // blankLogs := api.GetZoneBlankLogs()
    // reports.CombineBlankLogs(blankLogs)
    dailyReportsByZone := []adbutler.ZoneReport{}
    for _, zoneReport := range reports.Data {
        exist := false
        for _, id := range zoneids {
            if zoneReport.Id == id {
                exist = true
                break
            }
        }
        if exist {
            dailyReportsByZone = append(dailyReportsByZone, zoneReport)
        }
	}
    // blankLogsByZone := []adbutler.ZoneBlankLogs{}
    // for _, blankLogs := range blankLogs.Data {
    //     exist := false
    //     for _, id := range zoneids {
    //         if blankLogs.Id == id {
    //             exist = true
    //             break
    //         }
    //     }
    //     if exist {
    //         blankLogsByZone = append(blankLogsByZone, blankLogs)
    //     }
	// }
	c.JSON(http.StatusOK, gin.H{
		"publishers": myPubs,
        "zones": publisherZones,
        // "blankLogsByZone": blankLogsByZone,
		"dailyReportsByZone": dailyReportsByZone,
	})
}