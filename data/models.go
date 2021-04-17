package data

type Set struct {
	Id          int `pg:",pk"`
	Name        string
	ProgramSets []*ProgramSet `pg:"rel:has-many" json:"-"`
	IsActive    bool          `json:"-"`
}

func NewSet(name string) Set {
	return Set{
		Name:     name,
		IsActive: true,
	}
}

type Program struct {
	Id          int `pg:",pk"`
	Name        string
	ProgramPath string
	IsActive    bool `json:"-"`
}

func NewProgram(name, programPath string) Program {
	return Program{
		Name:        name,
		ProgramPath: programPath,
		IsActive:    true,
	}
}

type Argument struct {
	Id           int `pg:",pk"`
	Argument     string
	Order        int
	ProgramSetId int
	ProgramSet   *ProgramSet `pg:"rel:has-one"`
	IsActive     bool        `json:"-"`
}

func NewArgument(
	argument string,
	order, programSetId int) Argument {
	return Argument{
		Argument:     argument,
		Order:        order,
		ProgramSetId: programSetId,
		IsActive:     true,
	}
}

type ProgramSet struct {
	Id              int `pg:",pk"`
	SetId           int
	Set             *Set `pg:"rel:has-one"`
	ProgramId       int
	Program         *Program `pg:"rel:has-one"`
	Monitor         int
	SnappedPosition string
	Arguments       []*Argument `pg:"rel:has-many"`
	IsActive        bool        `json:"-"`
}

func NewProgramSet(
	setId, programId, monitor int,
	snappedPosition string) ProgramSet {
	return ProgramSet{
		SetId:           setId,
		ProgramId:       programId,
		Monitor:         monitor,
		SnappedPosition: snappedPosition,
		IsActive:        true,
	}
}
