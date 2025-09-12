Créez une structure Character avec les attributs suivants :
- Un nom
- Une classe
- Un niveau
- Des points de vie maximum
- Des points de vie actuels
- Un inventaire

Créez une fonction initCharacter permettant d’initialiser un personnage en
utilisant la structure Character.

Dans votre fonction Main, initialisez un personnage characters.c1 à l’aide de la
fonction initCharacter avec les valeurs suivantes :
- Nom : votre nom
- Classe: Elfe
- Niveau : 1
- Points de vie maximum : 100
- Points de vie actuels : 40
- Inventaire : 3 potions

Créez la fonction displayInfo qui permet d’afficher les
informations du personnage.

Créez la fonction accessInventory qui permet d’afficher tous les
items présents dans l’inventaire du personnage qui seront
utilisables par la suite.

Créez une fonction takePot qui permet d’utiliser une potion dans l’inventaire.
Vous pouvez l’utiliser dans le menu de « Accéder à l’inventaire ». Lorsque vous
utilisez une potion, celle-ci se consomme (supprimée de l’inventaire) et vous
regagnez 50 points de vie actuel. Puis affichez les points de vie actuel sur les
points de vie max du personnage.
Les points de vie actuels ne peuvent pas excéder
les points de vie maximum

Réalisez un menu composé des choix suivants :
- Afficher les informations du personnage
- Accéder au contenu de l’inventaire
- Quitter
Utilisez un switch case
Vous serez amenés à récupérer l’entrée utilisateur

Liez les choix du menu avec les fonctions créées précédemment
Le choix quitter permet d’arrêter le programme
N’oubliez pas d’ajouter le choix de
“Retour” pour naviguer entre les
différentes fonctionnalités

Créez la fonction accessInventory qui permet d’afficher tous les
items présents dans l’inventaire du personnage qui seront
utilisables par la suite. Ajoutez un choix « Marchand » au menu.
Créer l’interface du marchand qui vend :
1 Potion de vie (gratuitement)
TACHE 7 : Marchand
Lorsqu’un item du marchand est choisi, il est ajouté à
l’inventaire et vous devrez afficher le nom de l’item après l’achat.
Astuce : Pour génériser votre code, vous pouvez créer des
fonctions telles que addInventory et removeInventory pour
gérer l’ajout et le retrait d’item de l’inventaire

Créez une fonction isDead qui vérifie si le joueur est à 0 point de
vie, si c’est le cas, il meurt puis il est ressuscité avec 50% de ses
points de vie maximum.

Créez une fonction poisonPot qui inflige 10 points de vie de dégâts
par seconde pendant 3s. A chaque seconde de dégâts, vous afficherez
les points de vie actuels sur les points de vie maximum.
Vous utiliserez la bibliothèque time pour delay les secondes de poison

import "time"

time.Sleep(1 * time.Second)

Ajoutez la Potion de poison dans le menu de vente du marchand.
Même principe que la potion de vie, lorsqu’elle est choisie dans le
marchand, elle est ajoutée à l’inventaire.

Ajoutez l’attribut skill à votre personnage qui prend une liste de sorts.
Modifier votre fonction initCharacter, puis vous ajouterez à votre personnage le
sort de base « Coup de poing ».
Créez maintenant une fonction spellBook qui permet d’ajouter le sort « Boule
de feu » à votre liste de skill.
Attention, un même sort ne peut être appris qu’une seule fois.
Ajouter l’item « Livre de Sort : Boule de Feu » dans le menu de vente du
marchand. A l’achat, il sera ajouté dans votre inventaire et vous pourrez l’utiliser
via le menu « Accéder à l’inventaire ». Appelez la fonction spellBook lors de
l’utilisation de « Livre de Sort : Boule de Feu ».

Créer une fonction characterCreation qui permet à l’utilisateur de créer par luimême un personnage à l’aide de la fonction Init :
Permettre à l’utilisateur de choisir son nom (uniquement des lettres)
Qu’importe la façon d’écrire le nom est modifié pour :
- Commencer par une majuscule
- Le reste en minuscule

Permettre à l’utilisateur de choisir sa classe parmi : Humain, Elfe et Nain
En fonction de sa classe, il aura :
- 100 points de vie maximum pour les Humains
- 80 points de vie maximum pour les Elfes
- 120 points de vie maximum pour les Nains
Points de vie actuels de départ = 50% des points de vie maximum
Niveau de départ est de 1
Sort : Coup de Poing
A noter cette tâche remplace l’initialisation du personnage dans la fonction
main (tâche 3)

Créez une fonction qui empêche le joueur d’avoir plus de 10 items
dans son inventaire. Faites un check avec cette fonction lors de
l’ajout d’un item

MISE EN PLACE DE L’ECONOMIE
ET DU SYSTEME DE FABRICATION
DE L’EQUIPEMENT

Ajoutez un attribut à votre structure Character qui représentera
l’argent du joueur. Donner au joueur 100 pièces d’or au départ.

Modifier votre marchand lorsque le joueur choisit les items suivants :
« Potion de vie » : le joueur perd 3 pièces d’or
« Potion de poison » : le joueur perd 6 pièces d’or
« Livre de Sort : Boule de feu » : le joueur perd 25 pièces d’or
Ajouter au menu de vente du marchand les items suivants :
« Fourrure de Loup » : le joueur perd 4 pièces d’or
« Peau de Troll » : le joueur perd 7 pièces d’or
« Cuir de Sanglier » : le joueur perd 3 pièces d’or
« Plume de Corbeau » : le joueur perd 1 pièce d’or
Si le joueur choisit un item, il est ajouté à l’inventaire et le coût en pièce d’or est déduit de sa bourse
d’argent.

Ajouter au menu principal le choix « Forgeron ». Lorsque le joueur choisit « Forgeron », il doit arriver
dans un autre menu à choix qui va lui proposer la liste d’équipements à fabriquer suivante :
Chapeau de l’aventurier
Tunique de l’aventurier
Bottes de l’aventurier
Si le joueur choisit un objet à fabriquer (et qu’il peut le fabriquer), il perd 5 pièces d’or puis
l’équipement est ajouté à son inventaire.

Attention, la fabrication des équipements nécessite des ressources, il faudra pour :
Chapeau de l’aventurier
1 Plume de Corbeau
1 Cuir de Sanglier
Tunique de l’aventurier
2 Fourrure de loup
1 Peau de Troll
Bottes de l’aventurier
1 Fourrure de loup
1 Cuir de Sanglier
A la fabrication, les matériaux de fabrication sont supprimés de l’inventaire.
N’oubliez pas de personnaliser les messages d’erreur, par exemple, si le joueur n’a pas les ressources
nécessaires ou si le joueur n’a pas l’argent nécessaire ou s’il n’a pas de place dans son inventaire etc.

Créer une structure Equipment qui contient les attributs suivants :
Un équipement de tête
Un équipement pour le torse
Un équipement pour les pieds
Ajouter à votre structure Character, un attribut équipement qui est basé sur la structure Equipment

Rendez utilisable les équipements fabriqués et ajoutez-les au bon emplacement dans votre Structure
Equipement lorsqu'ils sont équipés. Les équipements équipés disparaissent de l’inventaire.
Modifier l’attribut points de vie maximum du personnage en fonction des équipements équipés :
Si Chapeau de l’aventurier équipé : +10 points de vie maximum
Si Tunique de l’aventurier équipé : +25 points de vie maximum
Si Bottes de l’aventurier équipé : +15 points de vie maximum
Lorsque le joueur équipe un équipement :
Si le joueur possède déjà un équipement équipé dans la même section.
L’équipement est remplacé et il récupère dans son inventaire le premier
équipement.

Créez une fonction upgradeInventorySlot qui lorsqu’elle est
utilisée augmente la capacité maximale de l’inventaire du joueur
de +10. Attention, le joueur peut utiliser une augmentation de
l’inventaire seulement 3 fois !
Ajoutez dans le marchand un objet « Augmentation d’inventaire »
pour 30 pièces d’or.


COMBAT
TOUR PAR TOUR

Vous allez créer un combat d’entrainement tour par tour entre le
joueur et un monstre
Créer une nouvelle structure Monster avec comme attribut :
Un nom
Des points de vie maximum
Des points de vie actuels
Des points d’attaque

Créer une fonction initGoblin qui avec la structure Monster
initiatialise les paramètres de base d’un gobelin d’entrainement :
Son nom sera : Gobelin d’entrainement
Ses points de vie maximum seront égaux à 40
Ses points de vie actuels seront égaux à ses points de vie
maximum
Ses points d’attaque sont égaux à 5

Créer une fonction goblinPattern qui va utiliser un Gobelin d’entrainement créé
par la fonction InitGoblin. Vous allez créer un pattern (schéma) de combat.
Chaque tour de combat, le Gobelin inflige 100% de son attaque en dégâts au
joueur. Tous les 3 tours (c’est-à-dire en fonction du tour de combat), il inflige 200%
de son attaque en dégâts au lieu de 100%.
Vous afficherez le nom de l’attaquant, le nom de la personne attaquée ainsi que le
montant de dégâts infligés.
Ex : « Gobelin d’entrainement inflige à Personnage 5 de dégâts »
Puis vous afficherez les points de vie actuel sur les points de vie max de la
personne attaquée.

Créer une fonction charaterTurn qui va simuler le tour de jeu du joueur.
Créer l’interface de combat de l’utilisateur :
Menu
Attaquer
Inventaire
Si le joueur choisit « Attaquer »
Il utilise « Attaque basique », infligez aux points de vie de l’adversaire 5 dégâts.
Vous afficherez, le nom de l'attaque utilisée, les dégâts infligés puis les points de
vie restants de l’adversaire. Puis c’est au tour du monstre de jouer.

Si le joueur choisit « Inventaire »
Afficher la liste des objets présents dans l’inventaire, si l’utilisateur choisit l’un
d’eux, il est utilisé et son effet s’applique. Puis c’est au tour du monstre de jouer.
N’oubliez pas d’afficher tout ce qui est réalisé.
Ex : « Personnage inflige 5 dégâts à Gobelin d’entrainement »
Ex : « Gobelin d’entrainement » « PV : 35 / 40 »
Ex : « Vous utilisez Potion de soin »
Etc…

Créer une fonction trainingFight qui va lancer un combat
d’entrainement contre un monstre. Le système de combat sera
tour par tour, c’est-à-dire que le joueur et le monstre vont jouer
l’un après l’autre.
Créer une variable qui permettra de
savoir à quel tour de combat
on se situe

Continuez la fonction trainingFight, appelez l’une après l’autre les
fonctions charTurn et goblinPattern.
Au début de chaque tour, afficher le tour de jeu
Ajouter au menu une option « Entrainement » pour appeler
trainingFight
N’oubliez pas, si les points de vie du monstre ou du joueur
tombent à 0 ou moins, le combat est terminé. Le joueur est alors
renvoyé au menu de départ.


Missions
Les missions vous permettent de
récupérer des points bonus


Ajoutez un attribut initiative au Personnage et au Gobelin, le tour
de jeu commence avec celui qui a le plus d’initiative

Ajouter un système d’expérience qui donne des points d’expérience
propre à chaque monstre au joueur à la fin d’un combat. Vous
pouvez ajouter des attributs Expérience actuelle et Expérience max
qui lorsque l’expérience maximum est atteinte fait augmenter le
niveau du joueur.
Attention à compter l’excès d’expérience pour le prochain niveau.
Vous pouvez également augmenter le nombre d’expérience
nécessaire à chaque niveau passé.
Tout gain de statistiques du personnage avec une montée en
niveau sera compté en bonus.

Permettre à l'utilisateur pour pouvoir utiliser des sorts durant le
combat qui infligent des dégâts à l'adversaire.
Coup de poing : 8 dégâts
Boule de feu : 18 dégâts
Vous pouvez ajouter d’autres sorts comme des sorts de soin, des
bonus ou des malus

Ajoutez un système de mana avec un attribut mana et mana max au
Personnage.
Les sorts « Coup de poing » et « Boule de feu » consomment du mana.
Réduisez le mana du personnage à leur utilisation et rendez impossible
leur utilisation si le mana est insuffisant.
Vous pouvez ajouter une potion de mana à l’image de la potion de vie
au marchand afin de regagner du mana.

Faites parler votre imagination pour enrichir le contenu de votre jeu !

Deux artistes sont cachés dans la partie 2 et 3 qui sont-ils?
Ajouter dans votre menu une option « Qui sont-ils » qui affiche le nom
des artistes cachés.

