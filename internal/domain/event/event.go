package event

import (
	"strconv"
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
		Next() Header
	}
	Event interface {
		// GetEventHeader イベントヘッダーを取得
		GetEventHeader() Header

		// GetEventName イベント名を取得
		GetEventName() Name
	}
)

type header struct {
	id      ID
	time    time.Time
	version Version
}

func NewHeader(id ID, time time.Time, version Version) Header {
	return header{
		id:      id,
		time:    time,
		version: version,
	}
}

func (h header) GetID() ID {
	return h.id
}

func (h header) GetTime() time.Time {
	return h.time
}

func (h header) GetVersion() Version {
	return h.version
}

func (h header) Next() Header {
	return header{
		id:      IssueID(),
		time:    time.Now(),
		version: h.version + 1,
	}
}

func IssueID() ID {
	return ID(uuid.New().String())
}

func (v Version) Int() int {
	return int(v)
}

func (v Version) String() string {
	return strconv.Itoa(v.Int())
}

func (v ID) String() string {
	return string(v)
}

func (n Name) String() string {
	return string(n)
}
