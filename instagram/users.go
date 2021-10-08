//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

type UsersService struct {
	client *Client
}

type User struct {
	ID             string     `json:"id,omitempty"`
	Username       string     `json:"username,omitempty"`
	FullName       string     `json:"full_name,omitempty"`
	ProfilePicture string     `json:"profile_picture,omitempty"`
	Bio            string     `json:"bio,omitempty"`
	Website        string     `json:"website,omitempty"`
	Counts         *UserCount `json:"counts,omitempty"`
}

type UserCount struct {
	Media      int `json:"media,omitempty"`
	Follows    int `json:"follows,omitempty"`
	FollowedBy int `json:"followed_by,omitempty"`
}

func (s *UsersService) Get(userId string) (*User, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v", userId)
	} else {
		u = "users/self"
	}
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	user := new(User)
	_, err = s.client.Do(req, user)
	return user, err
}

func (s *UsersService) MediaFeed(opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := "users/self/feed"
	if opt != nil {
		params := url.Values{}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
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

func (s *UsersService) RecentMedia(userId string, opt *Parameters) ([]Media, *ResponsePagination, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v/media/recent", userId)
	} else {
		u = "users/self/media/recent"
	}
	if opt != nil {
		params := url.Values{}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
		}
		if opt.MaxTimestamp != 0 {
			params.Add("max_timestamp", strconv.FormatInt(opt.MaxTimestamp, 10))
		}
		if opt.MinTimestamp != 0 {
			params.Add("min_timestamp", strconv.FormatInt(opt.MinTimestamp, 10))
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

func (s *UsersService) LikedMedia(opt *Parameters) ([]Media, *ResponsePagination, error) {
	u := "users/self/media/liked"
	if opt != nil {
		params := url.Values{}
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
		}
		if opt.MaxID != "" {
			params.Add("max_like_id", opt.MaxID)
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

func (s *UsersService) Search(q string, opt *Parameters) ([]User, *ResponsePagination, error) {
	u := "users/search"
	params := url.Values{}
	params.Add("q", q)
	if opt != nil {
		if opt.Count != 0 {
			params.Add("count", strconv.FormatUint(opt.Count, 10))
		}
	}
	u += "?" + params.Encode()

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, nil, err
	}

	users := new([]User)

	_, err = s.client.Do(req, users)
	if err != nil {
		return nil, nil, err
	}

	page := new(ResponsePagination)
	if s.client.Response.Pagination != nil {
		page = s.client.Response.Pagination
	}

	return *users, page, err
}

//Rishabh Johri 19BDS0021
