// Code generated by "stringer -type Period"; DO NOT EDIT.

package period

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Weekly-1]
	_ = x[Monthly-2]
	_ = x[Yearly-3]
	_ = x[Daily-4]
}

const _Period_name = "UnknownWeeklyMonthlyYearlyDaily"

var _Period_index = [...]uint8{0, 7, 13, 20, 26, 31}

func (i Period) String() string {
	if i < 0 || i >= Period(len(_Period_index)-1) {
		return "Period(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Period_name[_Period_index[i]:_Period_index[i+1]]
}
