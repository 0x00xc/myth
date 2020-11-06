/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/26 17:53
 */
package task

import "sync"

type Runnable interface {
	Run()
}

type Manager struct {
	wg    *sync.WaitGroup
	limit chan bool
}

func NewManager(limit int) *Manager {
	m := new(Manager)
	if limit > 0 {
		m.limit = make(chan bool, limit)
	}
	return m
}

func (m *Manager) do(f func()) {
	m.wg.Add(1)
	go func() {
		f()
		if m.limit != nil {
			<-m.limit
		}
		m.wg.Done()
	}()
}

func (m *Manager) Do(f func()) {
	if m.limit != nil {
		m.limit <- true
	}
	m.do(f)
}

func (m *Manager) DoOrCancel(f func()) bool {
	if m.limit != nil {
		select {
		case m.limit <- true:
		default:
			return true
		}
	}
	m.do(f)
	return true
}

func (m *Manager) Run(runnable Runnable) {
	m.Do(runnable.Run)
}

func (m *Manager) RunOrCancel(runnable Runnable) bool {
	return m.DoOrCancel(runnable.Run)
}

func (m *Manager) Wait() {
	m.wg.Wait()
}
