package chassis

import (
	"fmt"
)

type Signature struct {
	Namespace   string
	Topic       string
	Description string
	Version     string
	Env         string
}

func (s *Signature) TopicName() string {
	return fmt.Sprintf("%s.%s", s.Namespace, s.Topic)
}
