// Code generated by "stringer -type=Type -linecomment"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Invalid-0]
	_ = x[Api-1]
	_ = x[Entrypoint-2]
}

const _Type_name = "invalidapientrypoint"

var _Type_index = [...]uint8{0, 7, 10, 20}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
