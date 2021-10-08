//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

type MediaService struct {
	client *Client
}

type Media struct {
	Type         string         `json:"type,omitempty"`
	UsersInPhoto []*UserInPhoto `json:"users_in_photo,omitempty"`
	Filter       string         `json:"filter,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	Comments     *MediaComments `json:"comments,omitempty"`
	Caption      *MediaCaption  `json:"caption,omitempty"`
	Likes        *MediaLikes    `json:"likes,omitempty"`
	Link         string         `json:"link,omitempty"`
	User         *User          `json:"user,omitempty"`
	UserHasLiked bool           `json:"user_has_liked,omitempty"`
	CreatedTime  int64          `json:"created_time,string,omitempty"`
	Images       *MediaImages   `json:"images,omitempty"`
	Videos       *MediaVideos   `json:"videos,omitempty"`
	ID           string         `json:"id,omitempty"`
	Location     *MediaLocation `json:"location,omitempty"`
}

type MediaComments struct {
	Count int        `json:"count,omitempty"`
	Data  []*Comment `json:"data,omitempty"`
}

type MediaLikes struct {
	Count int `json:"count,omitempty"`
	Data  []*User
}

type MediaCaption struct {
	CreatedTime int64  `json:"created_time,string,omitempty"`
	Text        string `json:"text,omitempty"`
	From        *User  `json:"from,omitempty"`
	ID          string `json:"id,omitempty"`
}

type UserInPhoto struct {
	User     *User                `json:"user,omitempty"`
	Position *UserInPhotoPosition `json:"position,omitempty"`
}

type UserInPhotoPosition struct {
	x float64 `json:"x,omitempty"`
	y float64 `json:"y,omitempty"`
}

type MediaImages struct {
	LowResolution      *MediaImage `json:"low_resolution,omitempty"`
	Thumbnail          *MediaImage `json:"thumbnail,omitempty"`
	StandardResolution *MediaImage `json:"standard_resolution,omitempty"`
}

type MediaImage struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type MediaVideos struct {
	LowResolution      *MediaVideo `json:"low_resolution,omitempty"`
	StandardResolution *MediaVideo `json:"standard_resolution,omitempty"`
}

type MediaVideo struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type MediaLocation struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

func (s *MediaService) Get(mediaId string) (*Media, error) {
	u := fmt.Sprintf("media/%v", mediaId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	media := new(Media)
	_, err = s.client.Do(req, media)
	return media, err
}

func (s *MediaService) Search(opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := "media/search"
	if opt != nil {
		params := url.Values{}
		if opt.Lat != 0 {
			params.Add("lat", strconv.FormatFloat(opt.Lat, 'f', 7, 64))
		}
		if opt.Lng != 0 {
			params.Add("lng", strconv.FormatFloat(opt.Lng, 'f', 7, 64))
		}
		if opt.MinTimestamp != 0 {
			params.Add("min_timestamp", strconv.FormatInt(opt.MinTimestamp, 10))
		}
		if opt.MaxTimestamp != 0 {
			params.Add("max_timestamp", strconv.FormatInt(opt.MaxTimestamp, 10))
		}
		if opt.Distance != 0 {
			params.Add("distance", strconv.FormatFloat(opt.Distance, 'f', 7, 64))
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

func (s *MediaService) Popular() ([]Media, *ResponsePagination, error) {
	u := "media/popular"
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
