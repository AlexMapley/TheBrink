# TheBrink

## Running
To start up the go binary locally, you can run:

```shell
bash run.sh
```

You can also start up a docker container running the program with:

build image:
```shell
docker-compose build  the_brink
```
run image:
```shell
docker-compose run --rm  the_brink
```

## Dependencies
- Docker
- `docker-compose`

## Mutation
This game is fairly freeform, and I've chosen to reject most principles of functional design - forgoing any seperation of distinct `inputs` and `outputs`.
Most struct methods will instead act as `mutations`, updating the struct and returning a success/failure indicator.

Mutation is meant to be a core theme of the game from the players viewpoint, so we'll embrace that idea in our code design as well.

## Screenshots
Bootup from a simple docker container:

![Intro](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/intro.png)

Level up your stats along the way:

![Stats1](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/stats1.png)

![Stats2](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/stats2.png)

Collect loot (eventually):

![Inventory](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/inventory.png)

Let your character fight for themselves in automatic battles:

![Battle](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/battle.png)

![Battle2](https://github.com/AlexMapley/the_brink/blob/master/assets/screenshots/battle2.png)
