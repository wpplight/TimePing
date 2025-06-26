package authgroup

import "timeping/internal/authgroup/jumptable"

func initusr() {
	ag.Uidtable=jumptable.New(0,60000)
}