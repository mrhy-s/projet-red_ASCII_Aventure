package skills

type Skill struct {
	Nom         string
	ClasseToUse string
	NiveauToUse int
	Description string
	Degats      int
}

var CoupDePoing *Skill
var TirÀLarc *Skill
var CoupDeHache *Skill
var BouleDeFeu *Skill
var Griffure *Skill
var CoupDeRage *Skill
var Soin *Skill
var C_temp_skill []string

func InitSkill(nom string, classeToUse string, niveauToUse int, description string, degats int) *Skill {
	return &Skill{
		Nom:         nom,
		ClasseToUse: classeToUse,
		NiveauToUse: niveauToUse,
		Description: description,
		Degats:      degats,
	}
}
