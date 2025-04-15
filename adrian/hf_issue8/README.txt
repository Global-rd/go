To run the dockerized application:

1. Start the database with docker composer:
 docker-compose up

2. Build the dockerized application:
 docker build -t webservice-docker-app .

3. Run the dockerized app:
 docker run --rm -p 8080:8080 webservice-docker-app

In case the application is having trouble to resolve the host.docker.internal hostname (e.g. on linux), run it as follows:
 docker run --add-host=host.docker.internal:<The IP of the docker0 interface> --rm -p 8080:8080 webservice-docker-app

