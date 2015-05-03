## Qu'est-ce ?

Ce script fonctionne en duo avec le HTTP Files Client. Il monte un serveur HTTP mettant à la disposition du client les endpoints suivants :

- GET /listFiles : sert un flux json représentant l'ensemble des fichiers disponibles au téléchargement
- GET /downloadFile : sert un fichier
- DELETE /deleteFile : supprime un fichier

## Comment utiliser le script

Il est possible d'utilisaton `go install` pour générer un executable. Si `go` est installé sur la machine, il est également possible de lancer le script avec la commmande `go run main.go`.

Le script attend les arguments suivants (dans l'ordre) :

- répertoire des fichiers à servir
- le domaine + le port sur lesquels servir
- le token de sécurité

## Exemple d'utilisation

`go run main.go ./files 0.0.0.0:8888 azerty`
