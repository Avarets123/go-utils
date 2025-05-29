package index

import "reflect"

type Index struct {
	Idx map[string]int
}

func New[T any](data []T, field string) *Index {
	mI := map[string]int{}

	for i, v := range data {
		iV := reflect.ValueOf(v).FieldByName(field).Interface()

		lv, _ := iV.(string)

		mI[lv] = i

	}

	return &Index{
		Idx: mI,
	}

}

func (i *Index) Add(value int, index string) {
	i.Idx[index] = value
}

func (i *Index) Delete(index string) {
	delete(i.Idx, index)
}

func (i *Index) Exists(index string) (int, bool) {
	b, has := i.Idx[index]
	return b, has
}
