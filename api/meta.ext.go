// Code generated by protoc-gen-gogo. DO NOT EDIT.
// [NOT GENERATED CODE] This is hand written, but the above header is needed
//  to keep this file out of bounds for the linter so we can use _ for uniformity.

package api

var ListWatchOptions_SortOrders_vname = map[int32]string{
	0: "none",
	1: "by-name",
	2: "ny=name-reverese",
	3: "by-version",
	4: "by-version-reverse",
	5: "by-creation-time",
	6: "by-creation-time-reverse",
	7: "by-mod-time",
	8: "by-mod-time-reverse",
}

var ListWatchOptions_SortOrders_vvalue = map[string]int32{
	"none":                     0,
	"by-name":                  1,
	"ny=name-reverese":         2,
	"by-version":               3,
	"by-version-reverse":       4,
	"by-creation-time":         5,
	"by-creation-time-reverse": 6,
	"by-mod-time":              7,
	"by-mod-time-reverse":      8,
}

var ListWatchOptions_SortOrders_normal = map[string]string{
	"none":                     "none",
	"by-name":                  "by-name",
	"ny=name-reverese":         "ny=name-reverese",
	"by-version":               "by-version",
	"by-version-reverse":       "by-version-reverse",
	"by-creation-time":         "by-creation-time",
	"by-creation-time-reverse": "by-creation-time-reverse",
	"by-mod-time":              "by-mod-time",
	"by-mod-time-reverse":      "by-mod-time-reverse",
}

func (x ListWatchOptions_SortOrders) String() string {
	return ListWatchOptions_SortOrders_vname[int32(x)]
}

var EventType_vname = map[int32]string{
	0: "create-event",
	1: "update-event",
	2: "delete-event",
}

var EventType_vvalue = map[string]int32{
	"create-event": 0,
	"update-event": 1,
	"delete-event": 2,
}

var EventType_normal = map[string]string{
	"create-event": "create-event",
	"update-event": "update-event",
	"delete-event": "delete-event",
}

func (x EventType) String() string {
	return EventType_vname[int32(x)]
}
