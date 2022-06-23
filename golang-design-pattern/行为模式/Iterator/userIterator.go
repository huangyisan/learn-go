package main

// user 迭代器具体实现

type userIterator struct {
	// index 记录索引位置
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

// 获取下一个user
func (u *userIterator) getNext() *user {
	if u.hasNext() {
		// 如果存在user, 那么获取user,并且 让索引加1
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
