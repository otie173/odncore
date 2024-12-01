![1730741125721](https://github.com/user-attachments/assets/be5e55ca-c7f4-41f7-9e13-717e094add34)
# The idea
The goal of the core is to be easy to use and understand, and this applies to:
- the core code
- how it works
- the settings of the core itself

Unlike cores of other games like Minecraft, where the core is the brain of the whole system, Odncore serves as an intermediary between clients. It simply:
- passes packets from client to client
- allows clients to process these packets

This makes the core not the most important part of the system, but rather a facilitator of communication between clients.

## Introduction
In the **docs/** folder you can find information such as:
- how the server's API works
- how does the server architecture work and why is this the case?
- what incoming and outgoing network packets exist and their description
- how to set up the config and what each setting is responsible for

## Contributing
Odncore is open source for a reason. Anyone can contribute. If you have any difficulties, please contact me via Discord. In order for your changes to be accepted, they must:
- the code must be easy to read and understandable
- your changes must fit the basic idea of the architecture: simplicity

### Pull Request Process
- Make sure that the code matches the style guide
- Add tests for the new functionality
- Update the documentation
- Create a PR with a description of the changes

### Code Style
- Use talking names
- Comment on non-obvious solutions
- Follow the KISS principle
