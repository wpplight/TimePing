package timeclient

import (
	"timeping/api/proto/netsend/netgrpc"
)



func (t *tengine) Init(host string) error{
	t.host=host
}
