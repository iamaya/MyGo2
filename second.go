package main

func second() {
	listinstanace := statuslist{inactive, 100}
	println(listinstanace.first, listinstanace.quantity)
}

type myStatus string
type myStatusValue int

const (
	active      myStatus = "Active"
	inactive    myStatus = "Inactive"
	unavailable myStatus = "Unavailable"
)

// const (
// 	active      myStatusValue = 1
// 	inactive    myStatusValue = 2
// 	unavailable myStatusValue = 3
// )

type statuslist struct {
	first    myStatus
	quantity int
}
