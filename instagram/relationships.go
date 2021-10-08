//Rishabh Johri 19BDS0021
package instagram

import (
	"fmt"
)

type RelationshipsService struct {
	client *Client
}

type Relationship struct {
	OutgoingStatus string `json:"outgoing_status,omitempty"`

	IncomingStatus string `json:"incoming_status,omitempty"`
}

func (s *RelationshipsService) Follows(userId string) ([]User, *ResponsePagination, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v/follows", userId)
	} else {
		u = "users/self/follows"
	}

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

func (s *RelationshipsService) FollowedBy(userId string) ([]User, *ResponsePagination, error) {
	var u string
	if userId != "" {
		u = fmt.Sprintf("users/%v/followed-by", userId)
	} else {
		u = "users/self/followed-by"
	}

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

func (s *RelationshipsService) RequestedBy() ([]User, *ResponsePagination, error) {
	u := "users/self/requested-by"
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

func (s *RelationshipsService) Relationship(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "", "GET")
}

func (s *RelationshipsService) Follow(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "follow", "POST")
}

func (s *RelationshipsService) Unfollow(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "unfollow", "POST")
}

func (s *RelationshipsService) Block(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "block", "POST")
}

func (s *RelationshipsService) Unblock(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "unblock", "POST")
}

func (s *RelationshipsService) Approve(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "approve", "POST")
}

func (s *RelationshipsService) Deny(userId string) (*Relationship, error) {
	return relationshipAction(s, userId, "deny", "POST")
}

func relationshipAction(s *RelationshipsService, userId, action, method string) (*Relationship, error) {
	u := fmt.Sprintf("users/%v/relationship", userId)
	if action != "" {
		action = "action=" + action
	}
	req, err := s.client.NewRequest(method, u, action)
	if err != nil {
		return nil, err
	}

	rel := new(Relationship)
	_, err = s.client.Do(req, rel)
	return rel, err
}
