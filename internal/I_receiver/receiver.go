package receiver

type Receiver interface {
	Read()
	Connect()
}

func Read(rcvr Receiver) {
	rcvr.Read()
}

func Connect(rcvr Receiver) {
	rcvr.Connect()
}
