package github

import (
	"context"
	"sync"

	"github.com/google/go-github/v28/github"
)

type UserMap struct {
	sync.RWMutex
	byUsername map[string]int64
	byUserID   map[int64]string
}

func newUserMap() *UserMap {
	return &UserMap{
		byUsername: make(map[string]int64),
		byUserID:   make(map[int64]string),
	}
}

func (um *UserMap) GetID(username string, client *github.Client) (int64, bool) {
	var result int64
	var ok bool

	um.RLock()
	result, ok = um.byUsername[username]
	um.RUnlock()
	if !ok {
		// Try again with the write-lock, so the user can be added if not found
		um.Lock()
		result, ok = um.byUsername[username]
		if !ok {
			user, _, err := client.Users.Get(context.Background(), username)
			if err == nil {
				um.AddUserLocked(*user.Login, *user.ID)
				result = *user.ID
				ok = true
			}
		}
		um.Unlock()
	}
	return result, ok
}

func (um *UserMap) GetUsername(id int64, client *github.Client) (string, bool) {
	var result string
	var ok bool

	um.RLock()
	result, ok = um.byUserID[id]
	um.RUnlock()
	if !ok {
		// Try again with the write-lock, so the user can be added if not found
		um.Lock()
		result, ok = um.byUserID[id]
		if !ok {
			user, _, err := client.Users.GetByID(context.Background(), id)
			if err == nil {
				um.AddUserLocked(*user.Login, *user.ID)
				result = *user.Login
				ok = true
			}
		}
		um.Unlock()
	}
	return result, ok
}

func (um *UserMap) AddUserLocked(username string, id int64) {
	um.byUsername[username] = id
	um.byUserID[id] = username
}

func (um *UserMap) AddUser(username string, id int64) {
	um.Lock()
	um.AddUserLocked(username, id)
	um.Unlock()
}
