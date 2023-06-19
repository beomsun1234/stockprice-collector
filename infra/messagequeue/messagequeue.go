package messagequeue

type Messagequeue interface {
	Connect()
	SendMessage()
}
type Kafka struct {
}

func NewKafka() Messagequeue {
	return &Kafka{}
}

func (k *Kafka) Connect() {

}

func (k *Kafka) SendMessage() {

}
