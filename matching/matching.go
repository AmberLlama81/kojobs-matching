package matching

import "time"

type Mentee struct {
	name      string
	registime time.Time
	career    string
}

type Mentor struct {
	mentorId int64
	career   string
}

type Matcher struct {
}

func New() Matcher {
	return Matcher{}
}

func Match() {
	return
}
