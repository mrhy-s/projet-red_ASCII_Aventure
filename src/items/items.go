package items

type Item struct {
	Nom                  string
	ClasseToUse          string
	NiveauToUse          int
	DurabilitéMaximum    int
	DurabilitéesActuelle int
	Description          string
	Type                 string
	Effet                string
	Degats               int
	Defense              int
}

var Potion_de_soin *Item
var Potion_de_poison *Item
var Epee_en_fer *Item
var Spell_book_bdf *Item
var Plume_de_corbeau *Item
var Fourrure_de_loup *Item
var Peau_de_troll *Item
var Cuir_de_sanglier *Item
var Tunique_de_laventurier *Item
var Chapeau_de_laventurier *Item
var Bottes_de_laventurier *Item

func InitItem(nom string, classeToUse string, niveauToUse int, durabilitéMaximum int, durabilitéesActuelle int, description string) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMaximum,
		DurabilitéesActuelle: durabilitéesActuelle,
		Description:          description,
	}
}

func InitPotion(nom string, classeToUse string, niveauToUse int, description string, effet string) *Item {
	return &Item{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
		Type:        "Consommable",
		Effet:       effet,
	}
}

func InitArme(nom string, classeToUse string, niveauToUse int, durabilitéMax int, description string, degats int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMax,
		DurabilitéesActuelle: durabilitéMax,
		Description:          description,
		Type:                 "Arme",
		Degats:               degats,
	}
}

func InitSpellBook(nom string, classeToUse string, niveauToUse int, description string) *Item {
	return &Item{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
		Type:        "Consommable",
	}
}

func InitRessources(nom string, description string) *Item {
	return &Item{
		Nom:         nom,
		Description: description,
		Type:        "Ressources",
	}
}

func InitArmure(nom string, classeToUse string, niveauToUse int, durabilitéMax int, description string, defense int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMax,
		DurabilitéesActuelle: durabilitéMax,
		Description:          description,
		Type:                 "Armure",
		Defense:              defense,
	}
}
