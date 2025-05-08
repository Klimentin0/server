package main

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{map[string]int{}}
}

type InMemoryUserStore struct {
	store map[string]int
}

func (i *InMemoryUserStore) PostComment(user string) {
	i.store[user]++
}

func (i *InMemoryUserStore) GetUserPosts(user string) int {
	return i.store[user]
}
