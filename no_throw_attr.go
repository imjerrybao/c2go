package main

type NoThrowAttr struct {
	Address  string
	Position string
	Children []interface{}
}

func parseNoThrowAttr(line string) *NoThrowAttr {
	groups := groupsFromRegex(
		"<(?P<position>.*)>",
		line,
	)

	return &NoThrowAttr{
		Address:  groups["address"],
		Position: groups["position"],
		Children: []interface{}{},
	}
}
