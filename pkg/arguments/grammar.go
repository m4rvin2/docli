package arguments

type Grammar struct {
	Arguments []Argument `alligopher:"(@@|"`
	_         string     `alligopher:"Pun|Sym|Oth|'\u0009')*"`
}

type Argument struct {
	Identifier string `alligopher:"@('-'|Let|Num)+"`
	Value      Value  `alligopher:"@@?"`
}

type Value struct {
	Assignment string `alligopher:"@'='"`
	String     string `alligopher:"@(Let|Num|Pun|Sym|Sep)+"`
}
