package repository

import (
	"demorestapi/internal/entity"
	"sync"
)

type DataProvider interface {
	GetUser(id string) *entity.User
	AddUser(u *entity.User) error
	UpdateUser(u *entity.User) error
}

type Repo struct {
	DataProvider
	*localCache
}

type localCache struct {
	mu    sync.RWMutex
	users map[string]*entity.User
}

func newCache() *localCache {
	c := make(map[string]*entity.User)
	return &localCache{
		users: c,
	}
}

func (l *localCache) getUser(id string) *entity.User {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.users[id]
}

func (l *localCache) setUser(u *entity.User) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.users[u.ID] = u
}

func NewRepo(p DataProvider) *Repo {
	return &Repo{
		DataProvider: p,
		localCache:   newCache(),
	}
}

func (r *Repo) GetUser(id string) (res *entity.User) {
	if res = r.localCache.getUser(id); res == nil {
		println("requested db")
		res = r.DataProvider.GetUser(id)
		r.localCache.setUser(res)
	}
	return res
}

func (r *Repo) AddUser(u *entity.User) (err error) {
	if err = r.DataProvider.AddUser(u); err == nil {
		r.localCache.setUser(u)
	}

	return
}

func (r *Repo) UpdateUser(u *entity.User) (err error) {
	if err = r.DataProvider.UpdateUser(u); err == nil {
		r.localCache.setUser(u)
	}

	return
}
