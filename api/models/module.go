package models

import (
	"github.com/google/uuid"
)

type Module struct {
	UUID       uuid.UUID //PRIM
	Structure  []ModuleStructure
	Template   bool       // Is the module a template or custom module
	TemplateID *uuid.UUID // The ID of the template used to create this FKEY
}

type ModuleStructure struct {
	Module     Module
	ModuleUUID uuid.UUID
	Lesson     Lesson
	LessonUUID *uuid.UUID
	Test       Test
	TestUUID   *uuid.UUID
	Rank       string
}

type Lesson struct {
	UUID uuid.UUID
}

type Test struct {
	UUID uuid.UUID
}
