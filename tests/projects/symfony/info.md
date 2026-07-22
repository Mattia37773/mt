# Code snippets

## Projekt Beschreibung

**Snippet Library** ist eine web Applikation für das ablagern von Codesnippets

### Kernfunktionen:

- **Upload:** Benutzer können eigene Code-Snippets erstellen und hochladen.
- **Suche:** Eine textsuche ermöglicht das schnelle Auffinden von Snippets anhand des Titels.
- **Filterung:** Snippets können nach Programmiersprache und der verwendeten Technologie gefiltert werden.

---

## Techstack & Architektur

Das Projekt nutzt folgende Technologien:

- **Backend:** Symfony mit Doctrine ORM
- **Frontend:** Twig, Tailwind CSS und DaisyUI
- **Datenbank:** MySQL

### Docker-Architektur

Die Anwendung ist vollständig containerisiert und in folgende Services unterteilt:

- `fpm` (codesnippets-fpm): Verarbeitet das PHP-Backend und führt die Applikation aus.
- `nginx`: Der Webserver, der die HTTP-Anfragen entgegennimmt und an den PHP-FPM-Container weiterleitet.
- `db`: Die MySQL-Datenbank zur Speicherung der Benutzer und Snippets.
- `phpmyadmin`: Ein webbasiertes Tool zur einfachen Verwaltung und Überprüfung der Datenbank.
- `mailpit`: Ein lokaler Mail-Server, der von der Applikation gesendete E-Mails abfängt und im Webinterface anzeigt.

---

## Installation

Benötigte Schritte, um das Projekt lokal aufzusetzen und zu starten.

1. Docker Container aufsetzen

    ```bash
    docker compose -f docker/docker-compose.yaml up -d
    ```

2. Backend-Dependencies installieren

    ```bash
    docker exec -it codesnippets-fpm composer i
    ```

3. Frontend-Dependencies installieren

    ```bash
    docker exec -it codesnippets-fpm npm i
    ```

4. Datenbank migrationen

    ```bash
    docker exec -it codesnippets-fpm php bin/console doctrine:migrations:migrate -n
    ```

5. Demodaten laden

    ```bash
    docker exec -it codesnippets-fpm php bin/console doctrine:fixtures:load -n
    ```

    > Alle automatisch generierten Benutzer haben das Passwort "passwort".

6. Frontend Dev-Server starten
    ```bash
    docker exec -it codesnippets-fpm npm run dev
    ```

---

## Dev container

Benötigte Schritte, um das Projekt lokal aufzusetzen und zu starten.

1. Dev container in Vscode starten

    ```bash
    docker compose -f docker/docker-compose.yaml up -d
    ```

2. Nach erfolgreichem start Migrattionen durchführen

    ```bash
    php bin/console doctrine:migrations:migrate -n
    ```

3. Es gibt für die locale entwicklung Demodaten

    ```bash
    php bin/console doctrine:fixtures:load -n
    ```

    > Alle automatisch generierten Benutzer haben das Passwort "passwort".

4. Frontend Dev-Server starten für hmr

    ```bash
    docker exec -it codesnippets-fpm npm run dev
    ```

5. Mit F5 kann in PHP Files der Debugger verwendet werden.
6. Es gibt eine Extansion um mit der DB zu komunizieren in der Seitenleiste. Es muss nur das db password eigegeben werden.

## Local Multistage Build

1.  In den Docker Ordner wechseln

    ```bash
    cd dcoker
    ```

2.  Den Docker Stack bauen

    ```bash
    docker compose -f docker-compose-build.yml up --build
    ```

3. Migrationen laufen lassen

    ```bash
    php bin/console doctrine:migrations:migrate -n
    ```

## Docker Hub Build

1.  In den Docker Ordner wechseln

    ```bash
    cd dcoker
    ```

2.  Den Docker Stack von Dockerhub pullen

    ```bash
    docker compose -f docker-compose-hub.yml up
    ```

3. Migrationen laufen lassen

    ```bash
    php bin/console doctrine:migrations:migrate -n
    ```

---

## Erreichbarkeit der Services

Sobald die Container laufen, sind die einzelnen Dienste über folgende URLs auf Ihrem localhost erreichbar:

- Webseite: http://localhost:8080
- phpMyAdmin: http://localhost:8081
- Mailpit: http://localhost:8025

```

```

```
