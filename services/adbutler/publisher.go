package adbutler

type Publisher struct {
    Id int  `json:"id"`
    Email string `json:"email"`
    Name string `json:"name"`
}

type PublishersResponse struct {
    Data []Publisher `json:"data"`
}

func (publisher Publisher) GetZoneIds(zones []Zone) (zoneids []int) {
	for _, z := range zones {
		if publisher.Id == z.Publisher {
			zoneids = append(zoneids, z.Id)
		}
	}
	return
}

func (publisher Publisher) GetZones(zones []Zone) (publisherZones []Zone) {
	for _, z := range zones {
		if publisher.Id == z.Publisher {
			publisherZones = append(publisherZones, z)
		}
	}
	return
}