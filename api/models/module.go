package models

import (
	"github.com/google/uuid"
)

type Module struct {
	UUID       uuid.UUID
	Structure  []ModuleStructure
	Template   bool       // Is the module a template or custom module
	TemplateID *uuid.UUID // The ID of the template used to create this FKEY
}

type ModuleStructure struct {
	Module   Module
	ModuleID uuid.UUID
	Lesson   Lesson
	LessonID *uuid.UUID
	Test     Test
	TestID   *uuid.UUID
	Rank     string
}

type Lesson struct {
	UUID uuid.UUID
}

type Test struct {
	UUID uuid.UUID
}
