package socialcache

import (
	"errors"
	. "github.com/gamingrobot/steamgo/internal"
	. "github.com/gamingrobot/steamgo/steamid"
	"sync"
)

// Group list is implemented as doubly-linked lists for thread-safety.
// They can be iterated over like so:
// 	for id, group := range client.Social.Groups.GetCopy() {
// 		log.Println(id, group.Name)
// 	}

// A thread-safe group list
type GroupsList struct {
	mutex sync.RWMutex
	byId  map[SteamId]*Group
}

// Returns a new groups list
func NewGroupsList() *GroupsList {
	return &GroupsList{byId: make(map[SteamId]*Group)}
}

// Adds a group to the group list
func (list *GroupsList) Add(group Group) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.byId[group.SteamId] = &group
}

// Removes a group from the group list
func (list *GroupsList) Remove(id SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	delete(list.byId, id)
}

// Adds a chat member to a given group
func (list *GroupsList) AddChatMember(id SteamId, member ChatMember) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	id = id.ChatToClan()
	group := list.byId[id]
	if group == nil { //Group doesn't exist
		list.Add(Group{SteamId: id})
	}
	if group.ChatMembers == nil { //New group chat
		group.ChatMembers = make(map[SteamId]ChatMember)
	}
	group.ChatMembers[member.SteamId] = member

}

// Removes a chat member from a given group
func (list *GroupsList) RemoveChatMember(id SteamId, member SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	id = id.ChatToClan()
	group := list.byId[id]
	if group == nil { //Group doesn't exist
		return
	}
	if group.ChatMembers == nil { //New group chat
		return
	}
	delete(group.ChatMembers, member)
}

// Returns a copy of the groups map
func (list *GroupsList) GetCopy() map[SteamId]Group {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	glist := make(map[SteamId]Group)
	for key, group := range list.byId {
		glist[key] = *group
	}
	return glist
}

// Returns a copy of the group of a given SteamId
func (list *GroupsList) ById(id SteamId) (Group, error) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	id = id.ChatToClan()
	if val, ok := list.byId[id]; ok {
		return *val, nil
	}
	return Group{}, errors.New("Group not found")
}

// Returns the number of groups
func (list *GroupsList) Count() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	return len(list.byId)
}

//Setter methods
func (list *GroupsList) SetName(id SteamId, name string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	id = id.ChatToClan()
	if val, ok := list.byId[id]; ok {
		val.Name = name
	}
}

func (list *GroupsList) SetAvatarHash(id SteamId, hash []byte) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	id = id.ChatToClan()
	if val, ok := list.byId[id]; ok {
		val.AvatarHash = hash
	}
}

func (list *GroupsList) SetRelationship(id SteamId, relationship EClanRelationship) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	id = id.ChatToClan()
	if val, ok := list.byId[id]; ok {
		val.Relationship = relationship
	}
}

// A Group
type Group struct {
	SteamId      SteamId
	Name         string
	AvatarHash   []byte
	Relationship EClanRelationship
	ChatMembers  map[SteamId]ChatMember
}

// A Chat Member
type ChatMember struct {
	SteamId     SteamId
	Permissions EChatPermission
	Rank        EClanRank
}
