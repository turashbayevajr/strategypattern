package main

import "fmt"

const (
	water = "Water"
	earth = "Earth"
	fire  = "Fire"
	air   = "Air"
)

type Element string
type Elements interface {
	useElement() Element
}
type Aang struct {
	element Element
}
type Katara struct {
	element Element
}
type Zuko struct {
	element Element
}
type Toph struct {
	element Element
}

func (a *Aang) useElement() Element {
	return a.element
}
func (k *Katara) useElement() Element {
	return k.element
}
func (z *Zuko) useElement() Element {
	return z.element
}
func (t *Toph) useElement() Element {
	return t.element
}

type Fight struct {
}

func (f *Fight) fight(w Elements) {
	fmt.Println(fmt.Sprintf("%T Fighting by using element %s", w, w.useElement()))
}

func main() {
	f := &Fight{}
	a := &Aang{air}
	k := &Katara{water}
	t := &Toph{earth}
	z := &Zuko{fire}
	f.fight(a)
	f.fight(k)
	f.fight(z)
	f.fight(t)
	a = &Aang{water}
	f.fight(a)
	a = &Aang{earth}
	f.fight(a)
	a = &Aang{fire}
	f.fight(a)

}
