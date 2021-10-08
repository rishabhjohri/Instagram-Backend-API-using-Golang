//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
)

type TagsService struct {
	client *Client
}

type Tag struct {
	MediaCount int    `json:"media_count,omitempty"`
	Name       string `json:"name,omitempty"`
}

func (s *TagsService) Get(tagName string) (*Tag, error) {
	u := fmt.Sprintf("tags/%v", tagName)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	tag := new(Tag)
	_, err = s.client.Do(req, tag)
	return tag, err
}

func (s *TagsService) RecentMedia(tagName string, opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := fmt.Sprintf("tags/%v/media/recent", tagName)
	if opt != nil {
		params := url.Values{}
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

func (s *TagsService) Search(q string) ([]Tag, *ResponsePagination, error) {
	u := "tags/search?q=" + q
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	tags := new([]Tag)

	_, err = s.client.Do(req, tags)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *tags, page, err
}
