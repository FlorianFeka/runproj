package data

type Set struct {
	Id          int `pg:",pk"`
	Name        string
	ProgramSets []*ProgramSet `pg:"rel:has-many"`
}

type Program struct {
	Id          int `pg:",pk"`
	Name        string
	ProgramPath string
}

type Argument struct {
	Id           int `pg:",pk"`
	Argument     string
	Order        int
	ProgramSetId int
	ProgramSet   *ProgramSet `pg:"rel:has-one"`
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
}
