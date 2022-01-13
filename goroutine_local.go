/**
 * @author Leo Li
 */
package simplegoroutinelocal

import (
	"github.com/tal-tech/go-zero/core/collection"
)

type GoRoutineLocal struct {
	localRoutineIdMap *collection.SafeMap
}

func NewGoRoutineLocal() *GoRoutineLocal {
	return &GoRoutineLocal{
		localRoutineIdMap: collection.NewSafeMap(),
	}
}

func (m *GoRoutineLocal) Set(k, v interface{}) {
	goId := CurGoroutineID()
	kvMap := collection.NewSafeMap()
	kvMap.Set(k, v)
	m.localRoutineIdMap.Set(goId, kvMap)
}

func (m *GoRoutineLocal) Get(k interface{}) (interface{}, bool) {
	goId := CurGoroutineID()
	kvMapInterface, ok := m.localRoutineIdMap.Get(goId)
	if ok {
		kvMap := kvMapInterface.(*collection.SafeMap)
		return kvMap.Get(k)
	} else {
		return nil, false
	}
}

func (m *GoRoutineLocal) Del(k interface{}) {
	goId := CurGoroutineID()
	kvMapInterface, ok := m.localRoutineIdMap.Get(goId)
	if ok {
		kvMap := kvMapInterface.(*collection.SafeMap)
		kvMap.Del(k)
	}
}

func (m *GoRoutineLocal) DelMap() {
	goId := CurGoroutineID()
	m.localRoutineIdMap.Del(goId)
}
