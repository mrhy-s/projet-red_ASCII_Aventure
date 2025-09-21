package menus

import (
	// Gardé au cas où il serait utilisé ailleurs
	"ASCII_Aventure/items"
	"ASCII_Aventure/startscreen"
)

// Capitale initialise les états des cases de la carte de la capitale
// et remplit les descriptions des zones spécifiques dans ZonesMap.
func Capitale() {
	startscreen.ClearScreen()
	// Réinitialise toutes les cases à "Inexplorée" par défaut
	// pour s'assurer d'une base propre avant de définir les zones.

	// Ligne 1
	Etat_1_1 = Inexplorée
	Etat_1_2 = Inexplorée
	Etat_1_3 = Inexplorée
	Etat_1_4 = Inexplorée
	Etat_1_5 = Inexplorée
	Etat_1_6 = Inexplorée
	Etat_1_7 = Inexplorée
	Etat_1_8 = Inexplorée
	Etat_1_9 = Inexplorée
	Etat_1_10 = Inexplorée
	Etat_1_11 = Inexplorée

	// Ligne 2
	Etat_2_1 = Inexplorée
	Etat_2_2 = Explorée  // ◦
	Etat_2_3 = Bâtiment  // ■ (Marchand)
	Etat_2_4 = Chemin    // ·
	Etat_2_5 = Chemin    // ·
	Etat_2_6 = Chemin    // ·
	Etat_2_7 = Bâtiment  // ■ (Forgeron)
	Etat_2_8 = Explorée  // ◦
	Etat_2_9 = Explorée  // ◦
	Etat_2_10 = Explorée // ◦
	Etat_2_11 = Inexplorée

	// Ligne 3
	Etat_3_1 = Inexplorée
	Etat_3_2 = Explorée  // ◦
	Etat_3_3 = Chemin    // ·
	Etat_3_4 = Explorée  // ◦
	Etat_3_5 = Chemin    // ·
	Etat_3_6 = Explorée  // ◦
	Etat_3_7 = Chemin    // ·
	Etat_3_8 = Explorée  // ◦
	Etat_3_9 = Explorée  // ◦
	Etat_3_10 = Explorée // ◦
	Etat_3_11 = Inexplorée

	// Ligne 4 (Grande Avenue Centrale)
	Etat_4_1 = Inexplorée
	Etat_4_2 = Chemin  // ·
	Etat_4_3 = Chemin  // ·
	Etat_4_4 = Chemin  // ·
	Etat_4_5 = Chemin  // ·
	Etat_4_6 = Chemin  // ·
	Etat_4_7 = Chemin  // ·
	Etat_4_8 = Chemin  // ·
	Etat_4_9 = Chemin  // ·
	Etat_4_10 = Chemin // ·
	Etat_4_11 = Inexplorée

	// Ligne 5
	Etat_5_1 = Inexplorée
	Etat_5_2 = Explorée  // ◦
	Etat_5_3 = Chemin    // ·
	Etat_5_4 = Bâtiment  // ■ (Bâtiment Quêtes PNJ)
	Etat_5_5 = Explorée  // ◦
	Etat_5_6 = Chemin    // ·
	Etat_5_7 = Explorée  // ◦
	Etat_5_8 = Bâtiment  // ■ (Caserne des Gardes)
	Etat_5_9 = Chemin    // ·
	Etat_5_10 = Explorée // ◦
	Etat_5_11 = Inexplorée

	// Ligne 6
	Etat_6_1 = Inexplorée
	Etat_6_2 = Explorée  // ◦
	Etat_6_3 = Chemin    // ·
	Etat_6_4 = Chemin    // ·
	Etat_6_5 = Chemin    // ·
	Etat_6_6 = Joueur    // ⋄ (Position de départ du joueur au centre)
	Etat_6_7 = Chemin    // ·
	Etat_6_8 = Chemin    // ·
	Etat_6_9 = Chemin    // ·
	Etat_6_10 = Explorée // ◦
	Etat_6_11 = Inexplorée

	// Ligne 7
	Etat_7_1 = Inexplorée
	Etat_7_2 = Explorée  // ◦
	Etat_7_3 = Chemin    // ·
	Etat_7_4 = Explorée  // ◦
	Etat_7_5 = Chemin    // ·
	Etat_7_6 = Explorée  // ◦
	Etat_7_7 = Chemin    // ·
	Etat_7_8 = Explorée  // ◦
	Etat_7_9 = Chemin    // ·
	Etat_7_10 = Explorée // ◦
	Etat_7_11 = Inexplorée

	// Ligne 8
	Etat_8_1 = Inexplorée
	Etat_8_2 = Explorée    // ◦
	Etat_8_3 = Chemin      // ·
	Etat_8_4 = Bâtiment    // ■ (Temple/Bibliothèque)
	Etat_8_5 = Chemin      // ·
	Etat_8_6 = Chemin      // ·
	Etat_8_7 = Chemin      // ·
	Etat_8_8 = Bâtiment    // ■ (Auberge)
	Etat_8_9 = Chemin      // ·
	Etat_8_10 = ZoneCombat // ⚔ (Arène)
	Etat_8_11 = Inexplorée

	// Ligne 9
	Etat_9_1 = Inexplorée
	Etat_9_2 = Explorée // ◦
	Etat_9_3 = Chemin   // ·
	Etat_9_4 = Chemin   // ·
	Etat_9_5 = Chemin   // ·
	Etat_9_6 = Chemin   // ·
	Etat_9_7 = Chemin   // ·
	Etat_9_8 = Chemin   // ·
	Etat_9_9 = Chemin   // ·
	Etat_9_10 = Chemin  // ·
	Etat_9_11 = Inexplorée

	// Ligne 10
	Etat_10_1 = Inexplorée
	Etat_10_2 = Explorée  // ◦
	Etat_10_3 = Explorée  // ◦
	Etat_10_4 = Explorée  // ◦
	Etat_10_5 = Bâtiment  // ■ (Porte Sud / Quartier des Artisans)
	Etat_10_6 = Chemin    // ·
	Etat_10_7 = Explorée  // ◦
	Etat_10_8 = Explorée  // ◦
	Etat_10_9 = Explorée  // ◦
	Etat_10_10 = Explorée // ◦
	Etat_10_11 = Inexplorée

	// Ligne 11
	Etat_11_1 = Inexplorée
	Etat_11_2 = Inexplorée
	Etat_11_3 = Inexplorée
	Etat_11_4 = Inexplorée
	Etat_11_5 = Inexplorée
	Etat_11_6 = Inexplorée
	Etat_11_7 = Inexplorée
	Etat_11_8 = Inexplorée
	Etat_11_9 = Inexplorée
	Etat_11_10 = Inexplorée
	Etat_11_11 = Inexplorée

	// Initialise la map des zones si elle ne l'est pas déjà (LoadMap peut faire cela)
	if ZonesMap == nil {
		ZonesMap = make(map[string]Zone)
	}

	// Met à jour les zones spécifiques de la capitale
	ZonesMap["3-2"] = Zone{
		Nom:         "Boutique du Marchand",
		Description: "Une boutique où l'on peut acheter et vendre toutes sortes d'objets. L'air y est rempli d'odeurs exotiques.",
		Ressources:  []items.Item{}, // Remplissez avec des objets spécifiques si nécessaire
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Marchand Darian", Dialogue: "Bienvenue, explorateur. J'ai de tout, si tu as l'or!", Quete: "Livraison spéciale", Recompense: "Pièces d'or, Potion de soin"}},
		Visitee:     false,
	}

	ZonesMap["7-2"] = Zone{
		Nom:         "Forge du Maître Artisan",
		Description: "L'odeur du métal chaud et le bruit rythmique des marteaux résonnent dans cet atelier ancestral.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Forgeron Grognar", Dialogue: "Besoin de réparations ou d'une nouvelle lame? Montre-moi ce que tu as!", Quete: "Minerai rare", Recompense: "Armure améliorée"}},
		Visitee:     false,
	}

	ZonesMap["4-5"] = Zone{
		Nom:         "Maison des Braves",
		Description: "Un lieu de rassemblement pour les aventuriers, où les histoires de bravoure et les quêtes sont échangées.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Capitaine Elara", Dialogue: "Les bandits frappent encore sur la route. Peux-tu nous aider?", Quete: "Chasse aux bandits", Recompense: "Expérience, Réputation"}},
		Visitee:     false,
	}

	ZonesMap["8-5"] = Zone{
		Nom:         "Caserne des Gardes Royales",
		Description: "Le quartier général de la garde royale. Les soldats s'entraînent ici sans relâche pour la sécurité du royaume.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Sergent Kael", Dialogue: "Reste vigilant. La ville est sûre sous notre surveillance... pour l'instant.", Quete: "", Recompense: ""}},
		Visitee:     false,
	}

	ZonesMap["4-8"] = Zone{
		Nom:         "Ancienne Bibliothèque du Savoir",
		Description: "Des milliers de parchemins et de livres anciens remplissent cette salle silencieuse, gardienne des connaissances oubliées.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Archiviste Lyra", Dialogue: "La connaissance est un pouvoir. Que cherches-tu dans ces pages?", Quete: "Livre perdu", Recompense: "Sortilège, Connaissance"}},
		Visitee:     false,
	}

	ZonesMap["8-8"] = Zone{
		Nom:         "Auberge du Dragon Endormi",
		Description: "Un endroit chaleureux pour se reposer, manger un bon repas et entendre les dernières rumeurs locales.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Aubergiste Boris", Dialogue: "Une chope de bière locale, mon ami? C'est la meilleure de la ville!", Quete: "", Recompense: ""}},
		Visitee:     false,
	}

	ZonesMap["10-8"] = Zone{
		Nom:         "Arène Royale",
		Description: "Le sable imbibé de sang des combats passés attend de nouveaux champions. La foule y est souvent en délire.",
		Ressources:  []items.Item{},
		Monstres:    []string{"Gladiateur Novice", "Bête de l'Arène"}, // Exemples de monstres/ennemis
		PNJs:        []PNJ{{Nom: "Maître de l'Arène", Dialogue: "Prouve ta valeur, ou finis au tapis! Les défis t'attendent.", Quete: "Défi de l'Arène", Recompense: "Gloire, Récompense Unique"}},
		Visitee:     false,
	}

	ZonesMap["5-10"] = Zone{
		Nom:         "Guilde des Artisans / Porte Sud",
		Description: "Un quartier animé où les artisans vendent leurs créations uniques. C'est aussi l'une des portes principales menant hors de la ville, vers le sud.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{{Nom: "Garde du Sud", Dialogue: "Bienvenue dans la capitale. Montrez patte blanche avant d'entrer.", Quete: "", Recompense: ""}},
		Visitee:     false,
	}

	// Zone de départ du joueur - Marque-la comme visitée
	ZonesMap["6-6"] = Zone{
		Nom:         "Place du Marché Central",
		Description: "Le cœur vibrant de la capitale, toujours bondé de monde, de cris de marchands et d'activités incessantes.",
		Ressources:  []items.Item{},
		Monstres:    []string{},
		PNJs:        []PNJ{}, // Ajoutez des PNJs génériques si vous le souhaitez
		Visitee:     true,    // Le joueur commence ici, donc elle est visitée
	}

	// Vous pouvez ajouter des descriptions génériques pour d'autres zones "Explorée" ou "Chemin" si vous voulez
	// qu'elles aient une description spécifique quand le joueur entre dedans.
	// Exemple:
	// ZonesMap["6-5"] = Zone{Nom: "Ruelle Commerçante Ouest", Description: "Une petite ruelle animée avec quelques échoppes.", Visitee: false}
}
