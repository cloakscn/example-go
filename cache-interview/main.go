package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id   int
	Name string
}

type Server struct {
	hits  int
	db    map[int]*User
	cache map[int]*User
}

// 创建一个新的Server实例
func NewServer() *Server {
	// 创建一个map，用于存储用户信息
	db := make(map[int]*User)

	// 循环100次，创建100个用户
	for i := 0; i < 100; i++ {
		// 将用户信息存储到map中
		db[i+1] = &User{
			Id:   i + 1,
			Name: "user" + strconv.Itoa(i+1),
		}
	}

	// 返回一个新的Server实例，包含用户信息和缓存
	return &Server{db: db, cache: make(map[int]*User)}
}

// 尝试从缓存中获取用户信息
func (s *Server) tryCache(userid int) (*User, bool) {
	// 从缓存中获取用户信息
	user, ok := s.cache[userid]
	// 返回用户信息和是否获取成功
	return user, ok
}

// 处理获取用户信息的请求
func (s *Server) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	// 从URL中获取用户ID
	idStr := r.URL.Query().Get("id")
	// 将用户ID转换为整数
	id, _ := strconv.Atoi(idStr)

	// 尝试从缓存中获取用户信息
	user, err := s.tryCache(id)
	if err {
		// 如果缓存中存在该用户，则直接返回
		json.NewEncoder(w).Encode(user)
		return
	}

	// 如果缓存中不存在该用户，则从数据库中查找
	user, ok := s.db[id]
	if !ok {
		// 如果数据库中也不存在该用户，则返回用户未找到
		json.NewEncoder(w).Encode("User not found")
		return
	}

	// 将用户信息存入缓存
	s.cache[id] = user

	// 命中次数加一
	s.hits++

	// 返回用户信息
	json.NewEncoder(w).Encode(user)
}

func main() {

}
