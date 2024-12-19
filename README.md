![1730741125721](https://github.com/user-attachments/assets/be5e55ca-c7f4-41f7-9e13-717e094add34)
# The idea
The idea of the server is to be easy to use and understand.

Unlike servers of other games like Minecraft where the server controls all game logic, Odncore serves as a bridge between clients. It simply:
- passes packets from client to client
- allows clients to process these packets

This makes the server not the most important part of the system, but rather a facilitator of communication between clients.

## Introduction
In the `docs/` folder you can find information such as:
- how the server's API works
- how does the server architecture work and why is this the case?
- what incoming and outgoing network packets exist and their description
- how to set up the config and what each setting is responsible for
