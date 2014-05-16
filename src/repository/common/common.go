package common

import (
	. "domain"
	"time"
)

// Unique
type Unique struct {
	IdΩ string `bson:"_id"`
}

func (unique Unique) Id() string {
	return unique.IdΩ
}

// Named
type Named struct {
	NameΩ string `bson:"name"`
}

func (named *Named) Name(Ω ...string) string {
	if len(Ω) > 0 {
		named.NameΩ = Ω[0]
	}
	return named.NameΩ
}

// Active
type Activable struct {
	ActiveΩ bool `bson:"active"`
}

func (active *Activable) Active(Ω ...bool) bool {
	if len(Ω) > 0 {
		active.ActiveΩ = Ω[0]
	}
	return active.ActiveΩ
}

// Tracked
type Trackable struct {
	CreatedΩ   time.Time `bson:"created"`
	CreatedByΩ IEmployee `bson:"createdBy"`
}

func (tracked Trackable) Created() time.Time {
	return tracked.CreatedΩ
}

func (tracked Trackable) CreatedBy() IEmployee {
	return tracked.CreatedByΩ
}

//Gender
type gender struct {
	Unique
	Named
}

func AllGenders() []IGender {
	return []IGender{Male, Female}
}

func newGender(id, name string) IGender {
	g := &gender{}
	g.IdΩ = id
	g.NameΩ = name
	return g
}

var Male IGender = newGender(MaleGenderId, "Муж")
var Female IGender = newGender(FemaleGenderId, "Жен")
