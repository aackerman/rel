package arel

type Predicator struct{}

// TODO: fix interface type
func (p *Predicator) NotEqual(other ...interface{}) *NotEqualNode {
	return NewNotEqualNode()
}

func (p *Predicator) NotEqualAny() {

}

func (p *Predicator) NotEqualAll() {

}

func (p *Predicator) Equal() {

}

func (p *Predicator) EqualAny() {

}

func (p *Predicator) EqualAll() {

}

func (p *Predicator) In() {

}

func (p *Predicator) InAny() {

}

func (p *Predicator) InAll() {

}

func (p *Predicator) NotIn() {

}

func (p *Predicator) NotInAny() {

}

func (p *Predicator) NotInAll() {

}

func (p *Predicator) Matches() {

}

func (p *Predicator) MatchesAny() {

}

func (p *Predicator) MatchesAll() {

}

func (p *Predicator) DoesNotMatch() {

}

func (p *Predicator) DoesNotMatchAny() {

}

func (p *Predicator) DoesNotMatchAll() {

}

func (p *Predicator) GreaterThanOrEqual() {

}

func (p *Predicator) GreaterThanOrEqualAny() {

}

func (p *Predicator) GreaterThanOrEqualAll() {

}

func (p *Predicator) GreaterThan() {

}

func (p *Predicator) GreaterThanAny() {

}

func (p *Predicator) GreaterThanAll() {

}

func (p *Predicator) LessThan() {

}

func (p *Predicator) LessThanAny() {

}

func (p *Predicator) LessThanAll() {

}

func (p *Predicator) LessThanOrEqual() {

}

func (p *Predicator) LessThanOrEqualAny() {

}

func (p *Predicator) LessThanOrEqualAll() {

}
