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
	Gold                 int
	PrixRevente          int
}

var Potion_de_soin *Item
var Potion_de_poison *Item
var Épée_en_fer *Item
var Spell_book_bdf *Item
var Plume_de_corbeau *Item
var Fourrure_de_loup *Item
var Peau_de_troll *Item
var Cuir_de_sanglier *Item
var Tunique_de_laventurier *Item
var Chapeau_de_laventurier *Item
var Bottes_de_laventurier *Item
var Dague_rouillée *Item
var Gourdin_clouté *Item
var Arc_tordu *Item
var Cuir_bouilli_rapiécé *Item
var Casque_bosselé *Item
var Bourse_de_cuir *Item
var Fer_brut *Item

func InitItem(nom string, classeToUse string, niveauToUse int, durabilitéMaximum int, durabilitéesActuelle int, description string, prixRevente int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMaximum,
		DurabilitéesActuelle: durabilitéesActuelle,
		Description:          description,
		PrixRevente:          prixRevente,
	}
}

func InitPotion(nom string, classeToUse string, niveauToUse int, description string, effet string, prixRevente int) *Item {
	return &Item{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
		Type:        "Consommable",
		Effet:       effet,
		PrixRevente: prixRevente,
	}
}

func InitArme(nom string, classeToUse string, niveauToUse int, durabilitéMax int, description string, degats int, prixRevente int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMax,
		DurabilitéesActuelle: durabilitéMax,
		Description:          description,
		Type:                 "Arme",
		Degats:               degats,
		PrixRevente:          prixRevente,
	}
}

func InitArmeMonster(nom string, classeToUse string, niveauToUse int, durabilitéActuelle int, durabilitéMax int, description string, degats int, prixRevente int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMax,
		DurabilitéesActuelle: durabilitéActuelle,
		Description:          description,
		Type:                 "Arme",
		Degats:               degats,
		PrixRevente:          prixRevente,
	}
}

func InitSpellBook(nom string, classeToUse string, niveauToUse int, description string, prixRevente int) *Item {
	return &Item{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
		Type:        "Consommable",
		PrixRevente: prixRevente,
	}
}

func InitRessources(nom string, description string, prixRevente int) *Item {
	return &Item{
		Nom:         nom,
		Description: description,
		Type:        "Ressources",
		PrixRevente: prixRevente,
	}
}

func InitArmure(nom string, classeToUse string, niveauToUse int, durabilitéMax int, durabilitéActuelle int, description string, prixRevente int) *Item {
	return &Item{
		Nom:                  nom,
		ClasseToUse:          classeToUse,
		NiveauToUse:          niveauToUse,
		DurabilitéMaximum:    durabilitéMax,
		DurabilitéesActuelle: durabilitéActuelle,
		Description:          description,
		Type:                 "Armure",
		PrixRevente:          prixRevente,
	}
}

func InitBourse(gold int) *Item {
	prixRevente := 3
	return &Item{
		Nom:         "Bourse de cuir",
		Description: "Quelques pièces d'or",
		Type:        "Objet",
		Gold:        gold,
		PrixRevente: prixRevente,
	}
}
