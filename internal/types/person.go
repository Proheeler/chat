package types

type Person struct {
	ClientID string
	CrmID    string
	Name     string
	Surname  string
}

type PersonList struct {
	Total int
	Data  []Person
}
