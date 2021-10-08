//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
)

type LikesService struct {
	client *Client
}

func (s *LikesService) MediaLikes(mediaId string) ([]User, error) {
	u := fmt.Sprintf("media/%v/likes", mediaId)
	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	users := new([]User)
	_, err = s.client.Do(req, users)
	return *users, err
}

func (s *LikesService) Like(mediaId string) error {
	return mediaLikesAction(s, mediaId, "POST")
}

func (s *LikesService) Unlike(mediaId string) error {
	return mediaLikesAction(s, mediaId, "DELETE")
}

func mediaLikesAction(s *LikesService, mediaId, method string) error {
	u := fmt.Sprintf("media/%v/likes", mediaId)
	req, err := s.client.NewRequest(method, u, "")
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
