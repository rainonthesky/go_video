package session

import (
	"sync"
	_ "../defs"
	"../dbops"
	)

var sessionMap *sync.Map
func init(){
	sessionMap=&sync.Map{}
}
func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
func LoadSessionsFromDB()  {
	res,err:=dbops.
	
}
