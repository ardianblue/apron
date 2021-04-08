# Exercises

## Exercise 1 - The Apron Web App

We have a simple web app written in Go, based on the user agent you will get a different http code/response. The app supports HTTPS certificate, in this example we have self signed certificates that are saved in the repository (not a good practice to have them on github, just used for the example). The web app redirects http traffic to https. We can put the app on a container by using the Dockerfile, please follow the guide below on how to run the app in Docker.


1. Make sure you have Docker installed and that your user has access to use Docker, you can run `docker ps` and if no errors come back you can continue with the next steps.
2. While you are in the directory where the Dockerfile is located run the following command to build a docker image.

### Build the Docker image
`docker build . -t apron-web-app` - this will take a couple seconds to build the image with the name **apron-web-app**

### Create and run the docker container
`docker run -p 80:80 -p 443:443 apron-web-app` - this will run the the docker container and our web app will be running as well, make sure to have your local ports not in use or otherwise map them to a different port.

Trust the self signed certificate and you should see the web app in the browser with HTTPS.

<img width="1392" alt="image" src="https://user-images.githubusercontent.com/82177263/114083450-fa70df80-987c-11eb-9777-3eb04ebda6c8.png">



## Exercise 2 - The Log Parser
The log parser is a simple script using unix tools to parse the access.log file for unique IP addresses.

### Executing the script
`bash access_log_parser.sh` - the script will produce a `unique_ip_addresses.log` with all unique IP addresses, if you want to count them you can run `cat unique_ip_addresses.log | nl`


## Exercise 3 - CI/CD our Apron WebApp
CI/CD is a very important part of the software developer life-cycle, the main idea behind it, it's to have a workflow that automates most our work that we would otherwise have to do manually. There are many different CI/CD platforms that we can use in this example we are going to use CircleCI.

### Configuration
Every CI/CD platform has some sort of configuration that needs to be taken care, usually it's a `yaml` configuration that you explain how your code needs to be build/tested/delivered, if you have any testing framework you can connect it with your cicd, if you have like a third-party tool to run security checks on your code etc. This configuration file needs to be on your git repository where you code is, this way CirlceCI will know what you are talking about.

While you are in your configuration file there are different configs that you can choose to add or remove, let's take an example for a dev environment. Generally you want your dev environment to run a build after every commit you push to your repository, this can be done easily by just filtering for your dev branch, this won't execute any other configured environments.  

#### CircleCI Example
```yaml
# THIS IS NOT A WORKING CONFIG, JUST AN EXAMPLE OF HOW A CIRCLECI CONFIG FILE LOOKS LIKE.
version: 2.1
jobs:
  build:
    working_directory: ~/repo
    docker:
      - image: circleci/golang:1.15.8
    steps:
      - checkout
      - restore_cache:
          keys:
            - go-mod-v4-{{ checksum "go.sum" }}
      - run:
          name: Install Dependencies
          command: go mod download
      - save_cache:
          key: go-mod-v4-{{ checksum "go.sum" }}
          paths:
            - "/go/pkg/mod"
      - run:
          name: Run tests
          command: |
            mkdir -p /tmp/test-reports
            gotestsum --junitfile /tmp/test-reports/unit-tests.xml
      - store_test_results:
          path: /tmp/test-reports
```

### Step 1 - Configure your ci/cd platform to watch for branch changes, this is generally done by web hooks
Configure CircleCI to look for changes in your dev branch, if a commit has been pushed execute your **BUILD** workflow

### Step 2 - Store your artifect in a safe registry
After your build has been successful, and you ran your tests, security check etc, it's time to build our docker image and push it to a docker registry, you have a few options here, you can use dockerhub, or any of the other cloud providers registry or you could even have your own private registry.

### Step 3 - Deploy the web app
Now that we have a freshly baked docker image, it's time to run/deploy it on your platform, this can be different depending on the way you deploy apps, you can have a kubernetes cluster where you deploye your docker image, or you can have a simple Linux server where you connect securly to it and run a script to pull your image from the registry and run it.

