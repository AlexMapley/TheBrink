docker build . --tag the_brink

image=$(docker images | grep the_brink | awk '{print $1}')

echo $image