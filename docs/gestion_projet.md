# Gestion de Projet - ASCII Aventure

## Vue d'ensemble du projet

ASCII Aventure est un jeu de rôle textuel développé en Go, proposant une expérience RPG complète avec gestion de personnages, système de combat, commerce et artisanat.

## Objectifs du projet

- Développer un RPG textuel fonctionnel en Go
- Implémenter des systèmes de jeu complexes (combat, économie, artisanat)
- Créer une architecture modulaire et maintenable
- Proposer une expérience utilisateur engageante

## Avancement Global

![Progression](https://img.shields.io/badge/Progression-50%25-orange)

**Fonctionnalités complétées :** 20/40  
**À faire :** 20 

## Milestones

### Milestone 1 : Fondations (Complété)
**Objectif :** Mettre en place les bases du jeu  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Créer une structure Character avec les attributs : nom, classe, niveau, points de vie maximum, points de vie actuels, inventaire
- [x] Créer une fonction initCharacter pour initialiser un personnage
- [x] Initialiser un personnage c1 dans la fonction Main avec les valeurs spécifiées
- [x] Créer la fonction displayInfo pour afficher les informations du personnage
- [x] Créer la fonction accessInventory pour afficher l'inventaire
- [x] Créer une fonction takePot pour utiliser une potion (restaure 50 PV, supprime de l'inventaire)

### Milestone 2 : Interface Utilisateur de Base (Complété)
**Objectif :** Créer l'interface principale du jeu  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Réaliser un menu avec switch case : Afficher infos personnage, Accéder inventaire, Quitter
- [x] Lier les choix du menu avec les fonctions créées
- [x] Ajouter option "Retour" pour navigation

### Milestone 3 : Système Commercial de Base (Complété)
**Objectif :** Implémenter le marchand et l'économie de base  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Ajouter choix "Marchand" au menu
- [x] Créer interface marchand vendant 1 Potion de vie (gratuite)
- [x] Créer fonctions addInventory et removeInventory
- [x] Afficher nom de l'item après achat
- [x] Créer fonction isDead (vérifie si joueur à 0 PV, ressuscite à 50% PV max)

### Milestone 4 : Système de Combat (En attente)
**Objectif :** Ajouter poison et mécaniques avancées  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Créer fonction poisonPot (10 dégâts/seconde pendant 3s avec time.Sleep)
- [ ] Ajouter Potion de poison au marchand

### Milestone 5 : Système de Sorts (En attente)
**Objectif :** Implémenter la magie et les compétences  
**Statut :** À faire  
**Date prévue :** À définir

- [x] Ajouter attribut skill à Character (liste de sorts)
- [x] Modifier initCharacter pour ajouter sort "Coup de poing"
- [ ] Créer fonction spellBook pour ajouter "Boule de feu" (pas de doublon)
- [ ] Ajouter "Livre de Sort : Boule de Feu" au marchand

### Milestone 6 : Création de Personnage (En attente)
**Objectif :** Permettre la création personnalisée  
**Statut :** À faire  
**Date prévue :** À définir

- [x] Créer fonction characterCreation pour création personnage par utilisateur
- [x] Permettre choix nom (lettres uniquement, format majuscule + minuscules)
- [x] Permettre choix classe : Humain (100 PV), Elfe (80 PV), Nain (120 PV)
- [x] PV actuels = 50% PV max, niveau 1, sort "Coup de Poing"

### Milestone 7 : Économie Avancée (En attente)
**Objectif :** Système économique complet  
**Statut :** À faire  
**Date prévue :** À définir

- [x] Créer fonction limitant inventaire à 10 items maximum
- [x] Vérifier limite lors ajout d'item
- [x] Ajouter attribut argent à Character (100 pièces d'or au départ)
- [ ] Modifier marchand avec prix : Potion vie (3 or), Potion poison (6 or), Livre sort (25 or)
- [ ] Ajouter items : Fourrure Loup (4 or), Peau Troll (7 or), Cuir Sanglier (3 or), Plume Corbeau (1 or)

### Milestone 8 : Système d'Artisanat (En attente)
**Objectif :** Forgeron et création d'équipements  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Ajouter choix "Forgeron" au menu principal
- [ ] Créer menu forgeron : Chapeau aventurier, Tunique aventurier, Bottes aventurier
- [ ] Implémenter recettes de fabrication avec ressources requises
- [ ] Coût fabrication : 5 pièces d'or + matériaux
- [ ] Gérer messages d'erreur (ressources, argent, inventaire plein)

### Milestone 9 : Système d'Équipement (En attente)
**Objectif :** Équipements et bonus  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Créer structure Equipment (tête, torse, pieds)
- [ ] Ajouter attribut équipement à Character
- [ ] Rendre équipements utilisables (disparaissent de l'inventaire)
- [ ] Modifier PV max selon équipements : Chapeau (+10), Tunique (+25), Bottes (+15)
- [ ] Gérer remplacement d'équipements
- [ ] Créer fonction upgradeInventorySlot (+10 slots, max 3 utilisations)
- [ ] Ajouter "Augmentation d'inventaire" au marchand (30 or)

### Milestone 10 : Système de Combat PvE (En attente)
**Objectif :** Combat contre des monstres  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Créer structure Monster (nom, PV max, PV actuels, points d'attaque)
- [ ] Créer fonction initGoblin (Gobelin d'entrainement : 40 PV, 5 attaque)
- [ ] Créer fonction goblinPattern (100% dégâts, 200% tous les 3 tours)
- [ ] Afficher attaquant, cible, dégâts et PV restants
- [ ] Créer fonction characterTurn avec menu : Attaquer, Inventaire
- [ ] Option Attaquer : Attaque basique 5 dégâts
- [ ] Option Inventaire : utiliser objets de l'inventaire
- [ ] Créer fonction trainingFight (combat tour par tour)
- [ ] Gérer variable tour de combat
- [ ] Alterner characterTurn et goblinPattern
- [ ] Ajouter option "Entrainement" au menu
- [ ] Gérer fin de combat (PV à 0)

### Milestone 11 : Fonctionnalités Bonus (En attente)
**Objectif :** Améliorations et finalisation  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Ajouter attribut initiative (détermine ordre de jeu)
- [ ] Système d'expérience (XP par monstre, montée niveau)
- [ ] Utiliser sorts en combat : Coup de poing (8 dégâts), Boule de feu (18 dégâts)
- [ ] Système de mana (coût sorts, potion mana)
- [ ] Enrichir contenu du jeu
- [ ] Trouver les deux artistes cachés dans parties 2 et 3
- [ ] Ajouter option menu "Qui sont-ils" affichant noms des artistes

## Architecture Technique

### Structure du Projet
```
ASCII_Aventure/
├── src/
│ ├── main.go                  # Point d'entrée
│ ├── characters/              # Gestion des personnages
│ ├── classes/                 # Définition des classes
│ ├── functions_actions/       # Actions du jeu
│ ├── functions_helper/        # Fonctions utilitaires
│ ├── items/                   # Gestion des objets
│ ├── menus/                   # Interface utilisateur
│ ├── skills/                  # Système de compétences
│ └── startscreen/             # système de lancement de l'interface
├── docs/                      # Documentation
└── README.md                  # Instructions projet
```

### Technologies
- **Langage :** Go (Golang) 1.25.0
- **Interface :** Console/Terminal
- **Architecture :** Modulaire avec packages séparés