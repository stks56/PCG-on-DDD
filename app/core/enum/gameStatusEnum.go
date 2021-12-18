package enum

type GameStatusEnum int

const (
	Started GameStatusEnum = iota
	Inited
	Playing
	Ended
)

func (status GameStatusEnum) String() string {
	switch status {
	case Started:
		return "Started"
	case Inited:
		return "Inited"
	case Playing:
		return "Playing"
	case Ended:
		return "Ended"
	default:
		return "Unknown"
	}
}
