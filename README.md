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

![Alt text](/assets/screenshots/intro.png?raw=true "foo")

Level up your stats along the way:

![Alt text](/assets/screenshots/stats1.png?raw=true)

![Alt text](/assets/screenshots/stats2.png?raw=true)

Collect loot (eventually):

![Alt text](/assets/screenshots/inventory.png?raw=true)

Let your character fight for themselves in automatic battles:

![Alt text](/assets/screenshots/battle.png?raw=true)

![Alt text](/assets/screenshots/battle2.png?raw=true)
