# acml
There exist 4 files, which are main.go, DockerFile, docker-compose.yml and homepage.html.

Also, there is a folder that is called Form that includes several templates.

Docker: It is known as containerization and we used it to run the main.go file with its dependencies so that it can run more quickly and reliably; because docker container image is light weight, standalone, executable package of software that includes everything needed to run an application.

Docker-compose: it is used to run more than a docker application, mainly used when having more than 1 service and want to use them in parallel so that they can run in isolated environment to run: docker-compose up.

Config: The twelve-factor app stores config in environment variables (often shortened to env vars or env). Env vars are easy to change between deploys without changing any code; unlike config files, there is little chance of them being checked into the code repo accidentally; and unlike custom config files, or other config mechanisms.
