package service

import (
	"fmt"
	"go-api/model"
	"go-api/util"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	// finished用来判断下面gorouting结束与否.
	finished := make(chan bool, 1)

	// 并发请求提升效率
	for _, u := range users {
		wg.Add(1)
		// 需要传入u的方式,否则上面for循环出来的u在真正gorouting执行中会变动
		go func(u *model.UserModel) {
			defer wg.Done()
			// 用shortId区分不同的携程,在日志里面就方便查看了
			shortId, err := util.GenShortId()
			if err != nil {
				// 在这里处理err
				errChan <- err
				return
			}
			userList.Lock.Lock()
			//以下操作会产生竞争,所以先加锁
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
			userList.Lock.Unlock()
		}(u)
	}

	// 等待结束
	go func() {
		wg.Wait()
		// 关闭finished表示结束了,这样下面的select就不会阻塞,而是继续执行下去.
		close(finished)
	}()
	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}
	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}
	return infos, count, nil
}
