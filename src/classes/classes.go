package classes

type Classe struct {
	Nom         string
	Description string
}

var Humain *Classe
var Elfe *Classe
var Nain *Classe

func InitClasse(nom string, description string) *Classe {
	return &Classe{
		Nom:         nom,
		Description: description,
	}
}
