package adbutler

type Zone struct {
    Id int  `json:"id"`
    Name string `json:"name"`
    Publisher int `json:"publisher"`
}

type ZonesResponse struct {
    Data []Zone `json:"data"`
}

func (zone Zone) GetPublisher(publishers []Publisher) (publisher Publisher) {
	for _, p := range publishers {
		if p.Id == zone.Publisher {
			publisher = p
			return
		}
	}
	return
}