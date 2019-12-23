package _exsample

import "time"

//go:generate go-gentemplate -types Employee,RegularEmployee,DepartmentID -src $GOFILE -dist gentemplate_$GOFILE -debug

type (
	// Employee は従業員
	Employee interface {
		// Work は働く
		Work()
	}

	// RegularEmployee は正社員
	RegularEmployee struct {
		id int
		name string
		salary int
		joinAt time.Time
	}

	// DepartmentID は部門ID
	DepartmentID int

	// Company は会社
	Company string
)