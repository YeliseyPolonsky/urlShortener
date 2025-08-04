package stat

import (
	"go-advance/pkg/event"
	"log"
)

type StatService struct {
	*event.EventBus
	*StatRepository
}

type StatServiceDep struct {
	*event.EventBus
	*StatRepository
}

func NewStatService(deps StatServiceDep) *StatService {
	return &StatService{
		deps.EventBus,
		deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.Subscribe() {
		if msg.Type == event.LinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Invalid LinkVisited data: ", msg.Data)
			}
			s.StatRepository.AddClick(id)
		}
	}
}
