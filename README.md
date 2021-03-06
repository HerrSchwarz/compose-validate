# compose-validate

[![Build Status](https://travis-ci.org/HerrSchwarz/compose-validate.svg?branch=develop)](https://travis-ci.org/HerrSchwarz/compose-validate) [![Circle CI] (https://circleci.com/gh/HerrSchwarz/compose-validate.png?style=shield&circle-token=a1708741450afe114a9717a6718d05d897707397)](https://circleci.com/gh/HerrSchwarz/compose-validate)

[![codecov](https://codecov.io/gh/HerrSchwarz/compose-validate/branch/develop/graph/badge.svg)](https://codecov.io/gh/HerrSchwarz/compose-validate)

Ever found yourself write a bunch of docker-compose files, ending up with hundreds of lines of compose config for each environment? And you do not know, if the config is still doing, what it is supposed to do? No possibility to unit test the config? compose-validate can help a little. Compose-validate can be used to validate docker compose files against a given rule set. You can check:

- if some services exist
- if these services are connected to some networks
- if the network_mode is set to a certain value
- and if there are certain labels present on these services 

## Complexity

Let's do the math. We assume, a project has:

- 15 Services (four micro-service deployment services, apache, mysql, registrator, swarm-agent, swarm-master, consul, consul-template, heka, elasticsearch, kibana, curator)
- each Service has a compose file with 20 lines
- we have five environments:
  - local (for development)
  - continous integration
  - test (internal tests)
  - stage (external tests
  - live
- there are specific diferences between the environments (e.g. sizing)

There would be 15 * 20 * 5 = 1500 lines of compose config. And projects can easily be bigger.

## validation rules

Validation rules can be written in a .yml file. Here a small example:

```
rules:
  frontend_services:
    services:
      - apache
      - mysql
    networks:
      - frontend
      - vpn
    network_mode: overlay
    labels:
      - SERVICE_NAME
      - health-check-port
  
  backend_services:
    services:
      - admin
      - api
      - elasticsearch
      - kibana
      - swarm-agent
      - swarm-master
      - consul
      - consul-template
      - heka
    network_mode: host
    labels:
      - SERVICE_NAME
      
  deployment-services:
    services:
      - micro-service-1
      - micro-service-2
      - micro-service-3
      - micro-service-4
    network_mode: none
    labels:
      - SERIVCE_IGNORE
```

The idea is to find services, which are similar and to validate the similiar properties. Maybe some services should be in the same network or all services should have some labels present.

This is a spare time project and my first experiment with go. Any Comments, reviews or PRs are welcome.
