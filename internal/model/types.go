package model

type Filter struct {
	Q           string
	Page        int
	Count       int
	WithDeleted bool
}

func (f *Filter) SetQ(q string) *Filter{
	f.Q = q
	return f
}
func (f *Filter) SetPage(i int) *Filter{
	f.Page = i
	return f
}
func (f *Filter) SetCount(i int) *Filter{
	f.Count = i
	return f
}

