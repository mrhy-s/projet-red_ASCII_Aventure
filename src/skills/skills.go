package skills

type Skill struct {
	Nom         string
	ClasseToUse string
	NiveauToUse int
	Description string
}

var CoupDePoing *Skill
var TirÀLarc *Skill
var CoupDeHache *Skill
var C_temp_skill []string

func InitSkill(nom string, classeToUse string, niveauToUse int, description string) *Skill {
	return &Skill{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
	}
}
