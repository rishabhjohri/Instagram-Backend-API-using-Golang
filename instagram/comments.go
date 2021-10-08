//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
	"net/url"
)

type CommentsService struct {
	client *Client
}

type Comment struct {
	CreatedTime int64  `json:"created_time,string,omitempty"`
	Text        string `json:"text,omitempty"`
	From        *User  `json:"from,omitempty"`
	ID          string `json:"id,omitempty"`
}

func (s *CommentsService) MediaComments(mediaId string) ([]Comment, error) {
	u := fmt.Sprintf("media/%v/comments", mediaId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	comments := new([]Comment)
	_, err = s.client.Do(req, comments)
	return *comments, err
}

func (s *CommentsService) Add(mediaId string, text []string) error {
	u := fmt.Sprintf("media/%v/comments", mediaId)
	params := url.Values{
		"text": text,
	}

	req, err := s.client.NewRequest("POST", u, params.Encode())
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

func (s *CommentsService) Delete(mediaId, commentId string) error {
	u := fmt.Sprintf("media/%v/comments/%v", mediaId, commentId)
	req, err := s.client.NewRequest("DELETE", u, "")
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
