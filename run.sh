#!/bin/sh

docker build . --tag the_brink

docker run -p 6060:6060 -ti the_brink