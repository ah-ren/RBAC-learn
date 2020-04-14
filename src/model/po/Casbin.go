package po

type Casbin struct {
	ID     uint32
	PType  string // p
	Role   string // sub
	Path   string // obj
	Method string // act
}
