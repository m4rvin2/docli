package docstring

type Grammar struct {
	Description string     `alligopher:"(Let|Num|Pun|Sym|Sep|Oth)*"`
	Arguments   []Argument `alligopher:"@*"`
}

type Argument struct {
	NewLines    string       `alligopher:"'\n'+"`
	Indentation string       `alligopher:"(' '|'\t')+"`
	Identifiers []Identifier `alligopher:"@*"`
}

type Identifier struct {
	Name      string `alligopher:"('-'|Let|Num)+"`
	Separator string `alligopher:"(','' ')"`
}
