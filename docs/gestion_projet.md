# Gestion de Projet - ASCII Aventure

## Vue d'ensemble du projet

ASCII Aventure est un jeu de rôle textuel développé en Go, proposant une expérience RPG complète avec gestion de personnages, système de combat, commerce et artisanat.

## Objectifs du projet

- Développer un RPG textuel fonctionnel en Go
- Implémenter des systèmes de jeu complexes (combat, économie, artisanat)
- Créer une architecture modulaire et maintenable
- Proposer une expérience utilisateur engageante

## Avancement Global

![Progression](https://img.shields.io/badge/Progression-87.5%25-green)

**Fonctionnalités complétées :** 35/40  
**À faire :** 5

## Milestones

### Milestone 1 : Fondations (Complété)
**Objectif :** Mettre en place les bases du jeu  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Créer une structure Character avec les attributs : nom, classe, niveau, points de vie maximum, points de vie actuels, inventaire (Louis)
- [x] Créer une fonction initCharacter pour initialiser un personnage (Louis)
- [x] Initialiser un personnage c1 dans la fonction Main avec les valeurs spécifiées (Louis)
- [x] Créer la fonction displayInfo pour afficher les informations du personnage (Louis)
- [x] Créer la fonction accessInventory pour afficher l'inventaire (Louis)
- [x] Créer une fonction takePot pour utiliser une potion (restaure 50 PV, supprime de l'inventaire) (Louis)

### Milestone 2 : Interface Utilisateur de Base (Complété)
**Objectif :** Créer l'interface principale du jeu  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Réaliser un menu avec switch case : Afficher infos personnage, Accéder inventaire, Quitter (Louis)
- [x] Lier les choix du menu avec les fonctions créées (Louis)
- [x] Ajouter option "Retour" pour navigation (Louis)

### Milestone 3 : Système Commercial de Base (Complété)
**Objectif :** Implémenter le marchand et l'économie de base  
**Statut :** Terminé  
**Date de fin :** Complété

- [x] Ajouter choix "Marchand" au menu (Louis)
- [x] Créer interface marchand vendant 1 Potion de vie (gratuite) (Louis)
- [x] Créer fonctions addInventory et removeInventory (Louis)
- [x] Afficher nom de l'item après achat (Louis)
- [x] Créer fonction isDead (vérifie si joueur à 0 PV, ressuscite à 50% PV max) (Wilfrid + Louis)

### Milestone 4 : Système de Combat (En attente)
**Objectif :** Ajouter poison et mécaniques avancées  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Créer fonction poisonPot (10 dégâts/seconde pendant 3s avec time.Sleep) (Wilfrid)
- [x] Ajouter Potion de poison au marchand (Wilfrid)

### Milestone 5 : Système de Sorts (En attente)
**Objectif :** Implémenter la magie et les compétences  
**Statut :** Terminé  
**Date prévue :** Complété

- [x] Ajouter attribut skill à Character (liste de sorts) (Louis)
- [x] Modifier initCharacter pour ajouter sort "Coup de poing" (Louis)
- [x] Créer fonction spellBook pour ajouter "Boule de feu" (pas de doublon) (Louis)
- [x] Ajouter "Livre de Sort : Boule de Feu" au marchand (Louis)

### Milestone 6 : Création de Personnage (En attente)
**Objectif :** Permettre la création personnalisée  
**Statut :** Terminé 
**Date prévue :** Complété

- [x] Créer fonction characterCreation pour création personnage par utilisateur (Louis)
- [x] Permettre choix nom (lettres uniquement, format majuscule + minuscules) (Louis)
- [x] Permettre choix classe : Humain (100 PV), Elfe (80 PV), Nain (120 PV) (Louis)
- [x] PV actuels = 50% PV max, niveau 1, sort "Coup de Poing" (Louis + Wilfird)

### Milestone 7 : Économie Avancée (En attente)
**Objectif :** Système économique complet  
**Statut :** Terminé
**Date prévue :** Complété

- [x] Créer fonction limitant inventaire à 10 items maximum (Louis)
- [x] Vérifier limite lors ajout d'item (Louis)
- [x] Ajouter attribut argent à Character (100 pièces d'or au départ) (Louis)
- [x] Modifier marchand avec prix : Potion vie (3 or), Potion poison (6 or), Livre sort (25 or) (Wilfried)
- [x] Ajouter items : Fourrure Loup (4 or), Peau Troll (7 or), Cuir Sanglier (3 or), Plume Corbeau (1 or) (Wilfrid)

### Milestone 8 : Système d'Artisanat (En attente)
**Objectif :** Forgeron et création d'équipements  
**Statut :** Terminé  
**Date prévue :** Complété

- [x] Ajouter choix "Forgeron" au menu principal (Louis)
- [x] Créer menu forgeron : Chapeau aventurier, Tunique aventurier, Bottes aventurier (Louis)
- [x] Implémenter recettes de fabrication avec ressources requises (Louis)
- [x] Coût fabrication : 5 pièces d'or + matériaux (Louis)
- [x] Gérer messages d'erreur (ressources, argent, inventaire plein) (Louis)

### Milestone 9 : Système d'Équipement (En attente)
**Objectif :** Équipements et bonus  
**Statut :** À faire  
**Date prévue :** À définir

- [x] Créer structure Equipment (tête, torse, pieds) (Louis)
- [x] Ajouter attribut équipement à Character (Louis)
- [x] Rendre équipements utilisables (disparaissent de l'inventaire) (Louis)
- [x] Modifier PV max selon équipements : Chapeau (+10), Tunique (+25), Bottes (+15) (Louis)
- [x] Gérer remplacement d'équipements (Louis)
- [x] Créer fonction upgradeInventorySlot (+10 slots, max 3 utilisations) (Louis)
- [x] Ajouter "Augmentation d'inventaire" au marchand (30 or) (Louis)

### Milestone 10 : Système de Combat PvE (En attente) (Louis ET Wilfrid)
**Objectif :** Combat contre des monstres  
**Statut :** À faire  
**Date prévue :** À définir

- [x] Créer structure Monster (nom, PV max, PV actuels, points d'attaque) (Wilfrid et Louis)
- [x] Créer fonction initGoblin (Gobelin d'entrainement : 40 PV, 5 attaque) (Louis)
- [x] Créer fonction goblinPattern (100% dégâts, 200% tous les 3 tours) (Louis)
- [x] Afficher attaquant, cible, dégâts et PV restants (Louis)
- [x] Créer fonction characterTurn avec menu : Attaquer, Inventaire (Louis)
- [x] Option Attaquer : Attaque basique 5 dégâts (Louis)
- [x] Option Inventaire : utiliser objets de l'inventaire (Louis)
- [x] Créer fonction trainingFight (combat tour par tour) (Louis)
- [x] Gérer variable tour de combat (Louis)
- [x] Alterner characterTurn et goblinPattern (Louis)
- [x] Ajouter option "Entrainement" au menu (Louis)
- [x] Gérer fin de combat (PV à 0) (Louis)

### Milestone 11 : Fonctionnalités Bonus (En attente)
**Objectif :** Améliorations et finalisation  
**Statut :** À faire  
**Date prévue :** À définir

- [ ] Ajouter attribut initiative (détermine ordre de jeu)
- [ ] Système d'expérience (XP par monstre, montée niveau)
- [ ] Utiliser sorts en combat : Coup de poing (8 dégâts), Boule de feu (18 dégâts)
- [ ] Système de mana (coût sorts, potion mana)
- [ ] Enrichir contenu du jeu
- [x] Trouver les deux artistes cachés dans parties 2 et 3
- [x] Ajouter option menu "Qui sont-ils" affichant noms des artistes
- [x] Ajouter un écran de chargement (Louis)
