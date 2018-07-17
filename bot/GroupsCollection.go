package bot

import "sync"

type GroupsCollection struct {
	sync.RWMutex
	groups []*Group
}

func NewGroupsCollection() *GroupsCollection {
	return &GroupsCollection{groups: make([]*Group, 0)}
}

func (gc *GroupsCollection) Groups() []*Group {
	gc.Lock()
	existingGroups := make([]*Group, len(gc.groups))
	copy(existingGroups, gc.groups)
	gc.Unlock()
	return existingGroups
}

func (gc *GroupsCollection) AddGroup(group *Group) {
	gc.Lock()
	gc.groups = append(gc.groups, group)
	gc.Unlock()
}

func (gc *GroupsCollection) GroupsCount() int {
	gc.Lock()
	defer gc.Unlock()
	return len(gc.groups)
}
