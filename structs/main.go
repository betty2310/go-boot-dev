package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	characterLimit := 0
	if membershipType == TypeStandard {
		characterLimit = 100
	} else {
		characterLimit = 1000
	}
	return User{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: characterLimit}}
}
