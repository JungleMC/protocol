package protocol

type ConnectionState int

const (
	Handshake ConnectionState = iota
	Status
	Login
	Play
)

func (p ConnectionState) String() string {
	switch p {
	case Handshake:
		return "handshake"
	case Status:
		return "status"
	case Login:
		return "login"
	case Play:
		return "play"
	}
	return "unknown"
}
