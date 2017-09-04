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
	client = New("1254217795", "AKIDkOq6C6qLgstiMqmGF2d3HKBpHYeZlpAH", "Rny65tVv9BQuHUVKxOehZFqJbifYN7g3", "ap-chengdu")
}

func assert(want, condition, got interface{}) {

}
