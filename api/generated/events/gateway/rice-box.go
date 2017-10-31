package eventsGwService

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "events.swagger.json",
		FileModTime: time.Unix(1509409439, 0),
		Content:     string("{\n  \"swagger\": \"2.0\",\n  \"info\": {\n    \"title\": \"Service name\",\n    \"version\": \"version not set\"\n  },\n  \"schemes\": [\n    \"http\",\n    \"https\"\n  ],\n  \"consumes\": [\n    \"application/json\"\n  ],\n  \"produces\": [\n    \"application/json\"\n  ],\n  \"paths\": {\n    \"/{O.Tenant}/eventPolicy\": {\n      \"post\": {\n        \"operationId\": \"AutoAddEventPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"EventPolicyV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/eventPolicy/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetEventPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.CreationTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"O.ModTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"Spec.Levels\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          },\n          {\n            \"name\": \"Spec.RetentionPolicy\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Spec.ExportPolicies\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"EventPolicyV1\"\n        ]\n      },\n      \"delete\": {\n        \"operationId\": \"AutoDeleteEventPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"EventPolicyV1\"\n        ]\n      },\n      \"put\": {\n        \"operationId\": \"AutoUpdateEventPolicy\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"body\",\n            \"in\": \"body\",\n            \"required\": true,\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventPolicy\"\n            }\n          }\n        ],\n        \"tags\": [\n          \"EventPolicyV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/events\": {\n      \"get\": {\n        \"operationId\": \"AutoListEvent\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEventList\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"description\": \"Name of the object, unique within a Namespace for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.CreationTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"O.ModTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"LabelSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"FieldSelector\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"PrefixWatch\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"boolean\",\n            \"format\": \"boolean\"\n          }\n        ],\n        \"tags\": [\n          \"EventV1\"\n        ]\n      }\n    },\n    \"/{O.Tenant}/events/{O.Name}\": {\n      \"get\": {\n        \"operationId\": \"AutoGetEvent\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"\",\n            \"schema\": {\n              \"$ref\": \"#/definitions/eventsEvent\"\n            }\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"O.Tenant\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Name\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"T.APIVersion\",\n            \"description\": \"APIVersion defines the version of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.Namespace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.ResourceVersion\",\n            \"description\": \"Resource version in the object store. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.UUID\",\n            \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"O.CreationTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"O.ModTime.time\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\",\n            \"format\": \"date-time\"\n          },\n          {\n            \"name\": \"Status.Severity\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.Description\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.ObjectRef.Kind\",\n            \"description\": \"Kind represents the type of the API object.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.ObjectRef.NameSpace\",\n            \"description\": \"Namespace of the object, for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.ObjectRef.Name\",\n            \"description\": \"Name of the object, unique within a Namespace for scoped objects.\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.Source.Component\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.Source.Node\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.CreatedTime\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          },\n          {\n            \"name\": \"Status.EventPolicy\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"type\": \"string\"\n          }\n        ],\n        \"tags\": [\n          \"EventV1\"\n        ]\n      }\n    }\n  },\n  \"definitions\": {\n    \"apiListMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version of object store at the time of list generation.\"\n        }\n      },\n      \"description\": \"ListMeta contains the metadata for list of objects.\"\n    },\n    \"apiListWatchOptions\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"LabelSelector\": {\n          \"type\": \"string\"\n        },\n        \"FieldSelector\": {\n          \"type\": \"string\"\n        },\n        \"PrefixWatch\": {\n          \"type\": \"boolean\",\n          \"format\": \"boolean\"\n        }\n      }\n    },\n    \"apiObjectMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Name\": {\n          \"type\": \"string\",\n          \"description\": \"Name of the object, unique within a Namespace for scoped objects.\"\n        },\n        \"Tenant\": {\n          \"type\": \"string\",\n          \"description\": \"Tenant is global namespace isolation for various objects. This can be automatically\\nfilled in many cases based on the tenant a user, who created the object, belongs go.\"\n        },\n        \"Namespace\": {\n          \"type\": \"string\",\n          \"description\": \"Namespace of the object, for scoped objects.\"\n        },\n        \"ResourceVersion\": {\n          \"type\": \"string\",\n          \"description\": \"Resource version in the object store. This can only be set by the server.\"\n        },\n        \"UUID\": {\n          \"type\": \"string\",\n          \"description\": \"UUID is the unique identifier for the object. This can only be set by the server.\"\n        },\n        \"Labels\": {\n          \"type\": \"object\",\n          \"additionalProperties\": {\n            \"type\": \"string\"\n          },\n          \"description\": \"Labels are arbitrary (key,value) pairs associated with any object.\"\n        },\n        \"CreationTime\": {\n          \"$ref\": \"#/definitions/apiTimestamp\",\n          \"title\": \"CreationTime is the creation time of Object\"\n        },\n        \"ModTime\": {\n          \"$ref\": \"#/definitions/apiTimestamp\",\n          \"title\": \"ModTime is the Last Modification time of Object\"\n        }\n      },\n      \"description\": \"ObjectMeta contains metadata that all objects stored in kvstore must have.\"\n    },\n    \"apiObjectRef\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Kind\": {\n          \"type\": \"string\",\n          \"description\": \"Kind represents the type of the API object.\"\n        },\n        \"NameSpace\": {\n          \"type\": \"string\",\n          \"description\": \"Namespace of the object, for scoped objects.\"\n        },\n        \"Name\": {\n          \"type\": \"string\",\n          \"description\": \"Name of the object, unique within a Namespace for scoped objects.\"\n        }\n      },\n      \"description\": \"ObjectRef contains identifying information about an object.\"\n    },\n    \"apiTimestamp\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"time\": {\n          \"type\": \"string\",\n          \"format\": \"date-time\"\n        }\n      }\n    },\n    \"apiTypeMeta\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Kind\": {\n          \"type\": \"string\",\n          \"description\": \"Kind represents the type of the API object.\"\n        },\n        \"APIVersion\": {\n          \"type\": \"string\",\n          \"description\": \"APIVersion defines the version of the API object.\"\n        }\n      },\n      \"description\": \"TypeMeta contains the metadata about kind and version for all API objects.\"\n    },\n    \"eventsAutoMsgEventPolicyWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/eventsEventPolicy\",\n          \"title\": \"ObjectMeta.Name will be an UUID for an Event object.\\nObjectMeta.Labels will be used to tag an events\\nwith Reason, UserInfo etc that will provide more\\ncontext for an event.\\neg: Reason: NodeJoined, NetworkDeleted, NicRejected\\n    User: user-foo\\nTBD: Should there be predefined list of labels for\\n     Reason or keep it free form ?\"\n        }\n      },\n      \"description\": \"-------------------------- Event Object -----------------------------\\nEvent is a system notification of a fault, condition or configuration\\nthat should be user visible. These objects are created internally by\\nEvent client and persisted in EventDB.\"\n    },\n    \"eventsAutoMsgEventWatchHelper\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Type\": {\n          \"type\": \"string\"\n        },\n        \"Object\": {\n          \"$ref\": \"#/definitions/eventsEvent\"\n        }\n      },\n      \"title\": \"EventSpec is empty for Event Object\"\n    },\n    \"eventsEvent\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"description\": \"Component from which the event is generated.\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\",\n          \"description\": \"Node name on which the event is generated.\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/eventsEventSpec\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/eventsEventStatus\"\n        }\n      },\n      \"title\": \"EventSource has info about the component and\\nhost/node that generated the event\"\n    },\n    \"eventsEventList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"title\": \"Severity represents the criticality level of an Event\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\",\n          \"title\": \"Description represents the human readable description of an Event\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/eventsEvent\"\n          },\n          \"title\": \"ObjectRef is the reference to the object associated with an event\"\n        }\n      },\n      \"title\": \"EventStatus is status of the Event object\"\n    },\n    \"eventsEventPolicy\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\"\n        },\n        \"O\": {\n          \"$ref\": \"#/definitions/apiObjectMeta\"\n        },\n        \"Spec\": {\n          \"$ref\": \"#/definitions/eventsEventPolicySpec\",\n          \"description\": \"Spec contains the configuration of the event policy.\"\n        },\n        \"Status\": {\n          \"$ref\": \"#/definitions/eventsEventPolicyStatus\",\n          \"description\": \"Status contains the current state of the event policy.\"\n        }\n      },\n      \"description\": \"-------------------------- Event Policy -----------------------------\\nEvent Policy represents the policy definition for Events.\\nEvent Client module will be consumer of this policy.\"\n    },\n    \"eventsEventPolicyList\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"T\": {\n          \"$ref\": \"#/definitions/apiTypeMeta\",\n          \"title\": \"Levels is a match list of levels permitted for event generation\"\n        },\n        \"ListMeta\": {\n          \"$ref\": \"#/definitions/apiListMeta\",\n          \"title\": \"RetentionPolicy specifies for how long the data is kept\"\n        },\n        \"Items\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"$ref\": \"#/definitions/eventsEventPolicy\"\n          },\n          \"title\": \"ExportPolicies defines the location, frequency and format of data to an external collector\"\n        }\n      },\n      \"description\": \"EventPolicySpec is the specification of an Event Policy,\\nIt consists of the Object Selector, Level selector,\\nRetention and Export policies.\\n\\nTBD: Decide if we need event specific collection policy\"\n    },\n    \"eventsEventPolicySpec\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Levels\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        },\n        \"RetentionPolicy\": {\n          \"type\": \"string\"\n        },\n        \"ExportPolicies\": {\n          \"type\": \"array\",\n          \"items\": {\n            \"type\": \"string\"\n          }\n        }\n      },\n      \"title\": \"EventPolicyStatus\"\n    },\n    \"eventsEventPolicyStatus\": {\n      \"type\": \"object\"\n    },\n    \"eventsEventSource\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Component\": {\n          \"type\": \"string\"\n        },\n        \"Node\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"eventsEventSpec\": {\n      \"type\": \"object\"\n    },\n    \"eventsEventStatus\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"Severity\": {\n          \"type\": \"string\"\n        },\n        \"Description\": {\n          \"type\": \"string\"\n        },\n        \"ObjectRef\": {\n          \"$ref\": \"#/definitions/apiObjectRef\"\n        },\n        \"Source\": {\n          \"$ref\": \"#/definitions/eventsEventSource\"\n        },\n        \"CreatedTime\": {\n          \"type\": \"string\"\n        },\n        \"EventPolicy\": {\n          \"type\": \"string\"\n        }\n      }\n    }\n  }\n}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1509139580, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "events.swagger.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../../../../../sw/api/generated/events/swagger`, &embedded.EmbeddedBox{
		Name: `../../../../../sw/api/generated/events/swagger`,
		Time: time.Unix(1509139580, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"events.swagger.json": file2,
		},
	})
}
