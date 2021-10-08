//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

type LocationsService struct {
	client *Client
}

// Location represents information about a location.
type Location struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

func (s *LocationsService) Get(locationId string) (*Location, error) {
	u := fmt.Sprintf("locations/%v", locationId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	location := new(Location)
	_, err = s.client.Do(req, location)
	return location, err
}

func (s *LocationsService) RecentMedia(locationId string, opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := fmt.Sprintf("locations/%v/media/recent", locationId)
	if opt != nil {
		params := url.Values{}
		if opt.MinTimestamp != 0 {
			params.Add("min_timestamp", strconv.FormatInt(opt.MinTimestamp, 10))
		}
		if opt.MaxTimestamp != 0 {
			params.Add("max_timestamp", strconv.FormatInt(opt.MaxTimestamp, 10))
		}
		if opt.MinID != "" {
			params.Add("min_id", opt.MinID)
		}
		if opt.MaxID != "" {
			params.Add("max_id", opt.MaxID)
		}
		u += "?" + params.Encode()
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	media := new([]Media)

	_, err = s.client.Do(req, media)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *media, page, err
}

func (s *LocationsService) Search(lat, lng float64, opt *Parameters) ([]Location, error) {
	u := "locations/search"
	params := url.Values{}
	params.Add("lat", strconv.FormatFloat(lat, 'f', 7, 64))
	params.Add("lng", strconv.FormatFloat(lng, 'f', 7, 64))
	if opt != nil {
		if opt.Distance != 0 {
			distance := opt.Distance
			if distance > 5000 {
				distance = 5000
			}
			params.Add("distance", strconv.FormatFloat(distance, 'f', 7, 64))
		}
	}
	u += "?" + params.Encode()
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	locations := new([]Location)
	_, err = s.client.Do(req, locations)
	return *locations, err
}
