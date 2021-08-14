package protocol

type Direction bool

const (
	ClientToServer = false
	ServerToClient = true
)

func (d Direction) String() string {
	switch d {
	case ClientToServer:
		return "c -> s"
	case ServerToClient:
		return "c <- s"
	}
	return ""
}
