package apierrors

import "regexp"

type Unauthorized struct {
	s string
}

func NewUnauthorized(s string) *Unauthorized {
	return &Unauthorized{s: s}
}
func (e *Unauthorized) Error() string {
	return e.s
}

type DuplicateEntry struct {
	s     string
	entry string
	key   string
}

func NewDuplicateEntry(e error) *DuplicateEntry {
	re := regexp.MustCompile(`Duplicate entry '(\w*)' for key '(\w*)'`)
	match := re.FindStringSubmatch(e.Error())
	var entry string
	var key string
	if len(match) == 3 {
		entry = match[1]
		key = match[2]
	}
	return &DuplicateEntry{
		s:     e.Error(),
		entry: entry,
		key:   key,
	}
}
func (e *DuplicateEntry) Error() string {
	return e.s
}
func (e *DuplicateEntry) Entry() string {
	return e.entry
}
func (e *DuplicateEntry) Key() string {
	return e.key
}

type NoRows struct {
	s string
}

func NewNoRows(e error) *NoRows {
	return &NoRows{
		s: e.Error(),
	}
}

func (e NoRows) Error() string {
	return e.s
}
