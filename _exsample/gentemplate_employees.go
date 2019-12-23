// Code generated by gentemplate; DO NOT EDIT.
package _exsample




// EmployeeSlice is Employee slice.
type EmployeeSlice []Employee

// Filter returns a itself consisting of the elements of this slice that match the specified function.
func (rcv EmployeeSlice) Filter(f func(Employee) bool) EmployeeSlice {
	l := make([]Employee, 0, len(rcv))
	for _, v := range rcv {
		if f(v) {
			l = append(l, v)
		}
	}
	return l
}
	
// RegularEmployeeSlice is RegularEmployee slice.
type RegularEmployeeSlice []*RegularEmployee

// Filter returns a itself consisting of the elements of this slice that match the specified function.
func (rcv RegularEmployeeSlice) Filter(f func(*RegularEmployee) bool) RegularEmployeeSlice {
	l := make([]*RegularEmployee, 0, len(rcv))
	for _, v := range rcv {
		if f(v) {
			l = append(l, v)
		}
	}
	return l
}
	
// DepartmentIDSlice is DepartmentID slice.
type DepartmentIDSlice []DepartmentID

// Filter returns a itself consisting of the elements of this slice that match the specified function.
func (rcv DepartmentIDSlice) Filter(f func(DepartmentID) bool) DepartmentIDSlice {
	l := make([]DepartmentID, 0, len(rcv))
	for _, v := range rcv {
		if f(v) {
			l = append(l, v)
		}
	}
	return l
}
	
		