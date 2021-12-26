package member

import (
	"errors"
	"time"
)

type Member struct {
	Name  string
	Email string
	Date  time.Time
}

var Members map[string]Member

func InitMember() {
	Members = make(map[string]Member)
}

func loadMemberByEmail(email string) (Member, error) {
	var member Member
	var err error
	if memberVal, ok := Members[email]; ok {
		member = memberVal
	} else {
		err = errors.New("Unknown email")
	}
	return member, err
}

func addMember(member Member) error {
	_, err := loadMemberByEmail(member.Email)
	var errAdd error
	if err != nil {
		member.Date = time.Now()
		Members[member.Email] = member
	} else {
		errAdd = errors.New("Member already exist")
	}
	return errAdd
}

func getListMembers() map[string]Member {

	return Members
}
