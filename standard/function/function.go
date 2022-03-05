package function

type baseStruct struct {
	val string
}

func NewBaseStruct(val string) *baseStruct{
	return &baseStruct{
		val: val,
	}
}

type ISuper interface {
	getValue() int
	selfMultiplying()
	getStruct() *baseStruct
	setStructVal(string)
}

type Quote struct {
	val int
	stc *baseStruct
}

func NewQuote(val int, stc *baseStruct) *Quote{
	return &Quote{
		val: val,
		stc: stc,
	}
}

func (q *Quote) getValue() int{
	return q.val
}

func (q *Quote) selfMultiplying() {
	q.val *= q.val
}

func (q *Quote) getStruct() *baseStruct{
	return q.stc
}

func (q *Quote) setStructVal(val string) {
	q.stc.val = val
}

//
type Direct struct {
	val int
	stc *baseStruct
}

func NewDirect(val int, stc *baseStruct) *Direct{
	return &Direct{
		val: val,
		stc: stc,
	}
}

func (d Direct) getValue() int{
	return d.val
}

func (d Direct) selfMultiplying(){
	d.val *= d.val
}

func (d Direct) getStruct() *baseStruct{
	return d.stc
}

func (d Direct) setStructVal(val string) {
	d.stc.val = val
}
