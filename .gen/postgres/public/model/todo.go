//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Todo struct {
	ID        int32 `sql:"primary_key"`
	Title     string
	Completed *bool
	CreatedAt *time.Time
}
