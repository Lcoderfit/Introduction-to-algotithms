package utils

import (
	"fmt"
	"github.com/futurez/litego/logger"
	"net/http"
	"sync"
	"time"
)

func SetAndGetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Println("Read cookie error")
		return
	}
	if cookie == nil {
		fmt.Println("cookie is nil")
		return
	}

	sessionId := cookie.Value
	if GvSessionMgr.getSession(sessionId) != nil {
		fmt.Println("sessionId: ", sessionId)
		return
	}

	sessionId = GvSessionMgr.addsession(1, 1, "lcoder")
	expireTime := time.Now().AddDate(0, 0, 1)
	cookie = &http.Cookie{
		Name: "sessionId",
		Value: sessionId,
		Expires: expireTime,
	}

	http.SetCookie(w, cookie)
}

// ??????????????????????
type sessionData struct {
	sessionId    string
	userId       int64
	liveroomId   int64
	nickname     string
	tickTime     time.Time //标记时间
	allLastId    int64     //？？？？
	fcousLastId  int64     //？？？？
	bfcousOnline bool
}

//更新xx???时间
//是否改为Unix时间
func (sd *sessionData) UpdateTime() {
	sd.tickTime = time.Now()
}

type sessionManager struct {
	mux        sync.RWMutex
	sessionMap map[string]*sessionData
}

//全局session管理器
var GvSessionMgr = &sessionManager{sessionMap: make(map[string]*sessionData)}

const (
	//心脏包过期时间
	HeartBeatTimeOut = time.Hour
	//重新启动go程（goroutine）
	PreStartSessionGC = time.Minute
)

//对过期的session进行GC清理
func (sm *sessionManager) GC() {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	delIdList := make([]string, 0, len(sm.sessionMap))
	now := time.Now()
	for id, session := range sm.sessionMap {
		//计算now与session.tickTime的差值, 如果session超过心脏包的过期时间，则放入delIdList中
		if now.Sub(session.tickTime) > HeartBeatTimeOut {
			delIdList = append(delIdList, id)
		}
	}

	//删除过期的session
	for _, id := range delIdList {
		//Overduc: 过期的
		logger.Debug("GC: Overduc sessionid=", id)
		delete(sm.sessionMap, id)
	}
	//AfterFunc另起一个go程等待时间段d过去，然后调用f。它返回一个Timer，
	//可以通过调用其Stop方法来取消等待和对f的调用。
	//没过一分钟重新GC一次
	time.AfterFunc(PreStartSessionGC, func() {
		GvSessionMgr.GC()
	})
}

//通过userId, liveroomId, nickname创建session
//返回32位的sessionId
func (sm *sessionManager) addsession(userId, liveroomId int64, nickname string) string {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sd := &sessionData{
		sessionId:    GenUUID(),
		userId:       userId,
		liveroomId:   liveroomId,
		nickname:     nickname,
		tickTime:     time.Now(),
		bfcousOnline: true,
	}
	sm.sessionMap[sd.sessionId] = sd
	logger.Info("addSession: sessionId=", sd.sessionId, ", userId=", sd.userId,
		", liveroomId=", liveroomId)

	return sd.sessionId
}

//删除session
func (sm *sessionManager) deleteSession(sessionId string) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	logger.Info("deleteSession: sessionId=", sessionId)
	delete(sm.sessionMap, sessionId)
}

//读取session, 通过sessionId获取sessionData, 不存在则为nil
func (sm *sessionManager) getSession(sessionId string) *sessionData {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	return sm.sessionMap[sessionId]
}
