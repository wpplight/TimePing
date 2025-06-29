package authgroup

import (
	"timeping/pkg/jumptable"
	"timeping/pkg/upool"
)



func initusr() {
	ag.Uidtable=jumptable.New(0,60000,&upool.Unusedpool)
}