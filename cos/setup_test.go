package cos

var client *Client

const (
	Equal  = 0
	Lt     = 1
	Gt     = 2
	Lte    = 3
	Gte    = 4
	TypeOf = 5
)

func setUp() {
	client = New(&Option{"", "", "", "", "", ""})

}

func assert(want, condition, got interface{}) {

}
