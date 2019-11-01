package segment

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// All errors that segment could return.
var (
	ErrPartSizeInvalid         = errors.New("part size invalid")
	ErrPartIntersected         = errors.New("part intersected")
	ErrSegmentAlreadyInitiated = errors.New("segment already initiated")
	ErrSegmentNotInitiated     = errors.New("segment not initiated")
	ErrSegmentPartsEmpty       = errors.New("segment parts are empty")
	ErrSegmentNotFulfilled     = errors.New("segment not fulfilled")
)

// Part is a part of segment.
type Part struct {
	Offset int64
	Size   int64

	Index int // The Index of this part.
}

func (p *Part) String() string {
	return fmt.Sprintf("Part {Offset: %d, Size: %d}", p.Offset, p.Size)
}

// Segment will hold the whole segment operations.
type Segment struct {
	ID    string
	Path  string
	Parts map[int64]*Part

	index int // current part Index
	l     sync.RWMutex
}

// NewSegment will init a new segment.
func NewSegment(path, id string) *Segment {
	return &Segment{
		ID:    id,
		Path:  path,
		Parts: make(map[int64]*Part),
	}
}

func (s *Segment) String() string {
	return fmt.Sprintf("Segment {ID: %s, Path: %s}", s.ID, s.Path)
}

// InsertPart will insert a part into a segment and return it's Index.
func (s *Segment) InsertPart(p *Part) (index int, err error) {
	if p.Size == 0 {
		panic(ErrPartSizeInvalid)
	}

	s.l.Lock()
	defer s.l.Unlock()

	// Update segment Index.
	p.Index = s.index
	s.index++

	s.Parts[p.Offset] = p
	return p.Index, nil
}

// ValidateParts will validate a segment's parts.
func (s *Segment) ValidateParts() (err error) {
	errorMessage := "%s validate parts failed: %w"

	s.l.RLock()
	defer s.l.RUnlock()

	// Zero parts are not allowed, cause they can't be completed.
	if len(s.Parts) == 0 {
		return fmt.Errorf(errorMessage, s, ErrSegmentPartsEmpty)
	}

	x := make([]int64, 0, len(s.Parts))
	for _, v := range s.Parts {
		x = append(x, v.Offset)
	}
	sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })

	// First part offset must be 0
	if x[0] != 0 {
		return fmt.Errorf(errorMessage, s, ErrSegmentNotFulfilled)
	}

	for idx := 1; idx < len(s.Parts); idx++ {
		last := s.Parts[x[idx-1]]
		cur := s.Parts[x[idx]]
		if last.Offset+last.Size != cur.Offset {
			return fmt.Errorf(errorMessage, s, ErrSegmentNotFulfilled)
		}
	}

	return nil
}
