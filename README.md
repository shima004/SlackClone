# Clean Architecture - ChatServer

## Description

This is a chat server that uses Clean Architecture.

## Directory Structure

### Enterprise Business Rules

- entities

### Application Business Rules

- usecases
  - inputport
  - interactor
  - repository

### Interface Adapters

- controllers

  - web
    - handler

- gateways
  - datasource
  - infra
  - repository

### Frameworks & Drivers

- frameworks
  - web
    - handler
    - router

### Others

- config
- mock
- utility
