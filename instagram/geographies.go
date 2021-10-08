//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

type GeographiesService struct {
	client *Client
}

func (s *GeographiesService) RecentMedia(geoId string, opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := fmt.Sprintf("geographies/%v/media/recent", geoId)
	if opt != nil {
		params := url.Values{}
		if opt.MinID != "" {
			params.Add("min_id", opt.MinID)
		}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
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
