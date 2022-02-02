# Créer un docker avec Nginx, BDD & Go

Pour lancer le docker :
- Cloner le projet et se rendre à la racine
-  Faire la commande `docker-compose up --build`

Ce repository contient :
- Un dossier go contenant un Dockerfile & un dossier scripts contenant le script go exécuté.
- Un dossier MySQL contenant un Dockerfile & un dump SQL pour le test.
- Un dossier nginx contenant un Dockerfile & une configuration nginx.
- Un docker-compose permettant de relier les différents conteneurs/Dockerfile ci-dessus. 
