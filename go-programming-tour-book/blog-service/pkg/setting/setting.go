package setting

import (
	"github.com/spf13/viper"
	"sync"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

type messagePool struct {
	pool *sync.Pool
}
type Message struct {
	Count int
}

var msgPool = &messagePool{
	// 如果消息池里没有消息，则新建一个Count值为0的Message实例
	pool: &sync.Pool{
		New: func() interface{} {
			return &Message{Count: 0}
		},
	},
}

func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}
