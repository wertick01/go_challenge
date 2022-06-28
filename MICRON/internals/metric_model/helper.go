package models

type Mhelper interface {
	beginhelper() bool
	endhelper() bool
	durationhelper() bool
	namehelper() bool
}

type MHelper struct {
	MH Mhelper
}

func (mtr *All) beginhelper() bool {
	if !mtr.Fl.Begin.IsZero() {
		return true
	}
	return false
}

func (mtr *All) endhelper() bool {
	if !mtr.Fl.End.IsZero() {
		return true
	}
	return false
}

func (mtr *All) durationhelper() bool {
	if mtr.Fl.Duration != 0 {
		return true
	}
	return false
}

func (mtr *All) namehelper() bool {
	if mtr.Fl.Name != "" {
		return true
	}
	return false
}