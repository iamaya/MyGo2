package main

func second() {
	listinstanace := statuslist{active, 100}

	println(listinstanace.first, listinstanace.quantity)
}

type myStatus string

const (
	active      myStatus = "Active"
	inactive    myStatus = "Inactive"
	unavailable myStatus = "Unavailable"
)

type statuslist struct {
	first    myStatus
	quantity int
}
