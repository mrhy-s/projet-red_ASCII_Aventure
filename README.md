# ASCII_Aventure

## Présentation du projet

ASCII_Aventure est un jeu de rôle textuel développé en Go. Le joueur incarne un aventurier qui peut explorer, combattre, gérer son inventaire et interagir avec des marchands dans un monde entièrement basé sur du texte et de l'art ASCII.

### Fonctionnalités principales

- **Système de personnages** : Créez et gérez jusqu'à 2 personnages avec différentes classes (Humain, Elfe, Nain)
- **Système de classes** : Chaque classe possède ses propres caractéristiques et capacités
- **Gestion d'inventaire** : Collectez et utilisez différents objets (potions, armes, équipements)
- **Système de compétences** : Apprenez et utilisez diverses compétences selon votre classe
- **Boutique du marchand** : Achetez des objets avec vos pièces d'or
- **Système de vie/mort** : Gestion des points de vie avec résurrection automatique
- **Interface ASCII** : Menus et affichages esthétiques en art ASCII



## Installation et configuration

### Prérequis

- **Go 1.25.0** (Testé uniquement dans cette version)
- Terminal compatible UTF-8 pour un affichage optimal des caractères spéciaux

### Installation du projet

1. **Clonez le dépôt** :
```bash
git clone https://github.com/mrhy-s/projet-red_ASCII_Aventure
cd projet-red_ASCII_Aventure
```

2. **Naviguez vers le dossier source** :
```bash
cd src
```

3. **Initialisez le module Go** (si ce n'est pas déjà fait) :
```bash
go mod tidy
```

---

## Lancement du projet

### Démarrage rapide

```bash
cd src
go run main.go
```

### Compilation (optionnel)

Si vous souhaitez créer un exécutable :

```bash
# Depuis le dossier src/
go build -o ascii_aventure main.go

# Puis lancez l'exécutable
./ascii_aventure        # Linux/macOS
ascii_aventure.exe      # Windows
```

## Guide d'utilisation

### Première utilisation

1. **Lancement** : Exécutez le jeu avec `go run main.go`
2. **Personnage par défaut** : Le jeu démarre avec Zinrel (Elfe niveau 1)
3. **Menu principal** : Utilisez les options numérotées pour naviguer

## Structure du projet

```
projet-red_N/
├── README.md                    # Ce fichier
├── docs/                        # Documentation projet
│   └── gestion_projet.md        # Document de gestion
└── src/                         # Code source
    ├── main.go                  # Point d'entrée principal
    ├── go.mod                   # Module Go
    ├── characters/              # Gestion des personnages
    │   └── characters.go
    ├── classes/                 # Classes de personnages
    │   └── classes.go
    ├── functions_actions/       # Actions du jeu
    │   └── functions_actions.go
    ├── functions_helper/        # Fonctions utilitaires
    │   └── helper.go
    ├── items/                   # Objets et équipements
    │   └── items.go
    ├── menus/                   # Interface utilisateur
    │   └── menus.go
    └── skills/                  # Système de compétences
        └── skills.go
```

## Classes disponibles

| Classe | Description | Spécialité |
|--------|-------------|-----------|
| **Humain** | Polyvalent et adaptable | Équilibré dans tous les domaines |
| **Elfe** | Agile et magique | Maîtrise du tir à l'arc |
| **Nain** | Robuste et artisan | Expert en combat à la hache |

## Système d'inventaire

- **Capacité** : 10 objets maximum
- **Types d'objets** :
  - Potions de soin (+50 PV)
  - Potions de poison (-20 PV)
  - Armes (dégâts variables)
  - Équipements divers

## Système économique

- **Pièces d'or** : Monnaie du jeu
- **Boutique** : Prix évoluent selon les tours
- **Potion gratuite** : Une potion offerte au début

## Système de combat

- **Points de vie** : Gestion automatique
- **Compétences** : Spécifiques à chaque classe
- **Mort** : Résurrection automatique avec pénalités

---

## Licence

Ce projet est développé dans le cadre d'un exercice pédagogique.

## 👥 Contribution

Louis Guérout
Wilfrid Delamare

**Bon jeu ! ◕‿◕ **
