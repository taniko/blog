package event

import (
	"time"

	"github.com/google/uuid"
)

type (
	ID      string
	Name    string
	Version int
)

type (
	Header interface {
		GetID() ID
		GetTime() time.Time
		GetVersion() Version
		GetName() Name
		Next(name Name) Header
	}
	Event interface {
		GetHeader() Header
	}
)

type header struct {
	id      ID
	time    time.Time
	version Version
	name    Name
}

func NewHeader(id ID, time time.Time, version Version, name Name) Header {
	return header{
		id:      id,
		time:    time,
		name:    name,
		version: version,
	}
}

func (h header) GetID() ID {
	return h.id
}

func (h header) GetName() Name {
	return h.name
}

func (h header) GetTime() time.Time {
	return h.time
}

func (h header) GetVersion() Version {
	return h.version
}

func (h header) Next(name Name) Header {
	return header{
		id:      h.id,
		time:    time.Now(),
		version: h.version + 1,
		name:    name,
	}
}

func IssueID() ID {
	return ID(uuid.New().String())
}
