package main

import (
	mcm "member-club/member"
	"net/http"
)

func init() {
	mcm.InitMember()
}

func main() {
	mcm.MemberHttpHandlers()
	http.ListenAndServe(":8899", nil)
}
