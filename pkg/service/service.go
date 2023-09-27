package service

import (
	model "github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
	groupModel "github.com/go-web/pkg/model/group"
	memberModel "github.com/go-web/pkg/model/member"
	userModel "github.com/go-web/pkg/model/user"
	venueModel "github.com/go-web/pkg/model/venue"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/repository/member"
	uuid "github.com/satori/go.uuid"
)

type EventService interface {
	CreateEvent(eventRequest *eventModel.Event) (*eventModel.EventResponse, error)
	GetEventsByGroupId(eventRequest *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventImpl{
		eventRepository,
	}
}

type EventMemberService interface {
	CreateEventMember(eventRequest *eventModel.AddMemberToEventRequest) (*eventModel.AddMemberToEventResponse, error)
	GetEventMembers(request *memberModel.GetEventMembersRequest) (*memberModel.GetEventMembersResponse, error)
}

func NewEventMemberService(eventRepository member.EventMemberRepository) EventMemberService {
	return &eventMemberImpl{
		eventRepository,
	}
}

type VenueService interface {
	CreateVenue(venueRequest *venueModel.Venue) (*venueModel.Venue, error)
	GetVenues(request *model.Filter) (*venueModel.GetVenueResponse, error)
}

func NewVenueService(venueRepository repository.VenueRepository) VenueService {
	return &VenueImpl{
		venueRepository,
	}
}

type GroupService interface {
	CreateGroupWithMembers(group *groupModel.CreateGroupModel) (*groupModel.CreateGroupModel, error)
	AddMembersToGroup(groupId uuid.UUID, groupMember []*groupModel.GroupMember) ([]*groupModel.GroupMember, error)
	GetGroup(groupId string) (*groupModel.Group, error)
	GetMembersByGroupId(groupId string) ([]*groupModel.GroupMemberByIdResponse, error)
	GetGroupsByMemberId(memberId string) ([]*groupModel.Group, error)
}

func NewGroupService(groupRepository repository.GroupRepository) GroupService {
	return &GroupImpl{
		groupRepository: groupRepository,
	}
}

type UserService interface {
	CreateUser(user *userModel.UserBase) (*userModel.User, error)
	GetUserByMemberId(memberId string) (*userModel.User, error)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserImpl{
		userRepository: userRepository,
	}
}
