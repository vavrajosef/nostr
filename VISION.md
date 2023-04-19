# VISION

## Introduction

### Purpose

The purpose of this requirements document is to outline the necessary features, functionality, and specifications for the Nostr Implementation Possibilities (NIPs) monorepo. This document will serve as a reference for contributors, developers, and users who are interested in understanding the scope of the project, its goals, and the expectations for implementing and interacting with the monorepo.

By providing clear and comprehensive requirements, we aim to facilitate effective collaboration, efficient development, and a consistent understanding of the project among all stakeholders. This document will guide the design, implementation, and maintenance of the monorepo to ensure it meets the needs of the Nostr community and adheres to the principles outlined in the NIPs.

### Scope

This requirements document focuses on the development of the go-nostr software, a Nostr-compatible implementation that includes both relay and client functionalities. The software will provide the necessary features and tools for developers and users to interact with the Nostr protocol, allowing for the creation and management of events, as well as communication between clients and relays. The go-nostr software aims to streamline the development of Nostr-compatible applications, promote the adoption of the Nostr network, and support the growth of the decentralized social platform. While not covering the implementation of every single NIP, it will focus on core functionalities and selected NIPs that are essential for a functional and robust Nostr system. 

### Definitions, Acronyms, and Abbreviations

TBD

### References

TBD

### Overvierw

TBD

## Overall Description

### Product Perspective

The go-nostr software is designed as a part of the Nostr network, which is a decentralized and extensible social protocol. The software acts as a bridge between Nostr-compatible applications, providing both relay and client functionalities to enable seamless communication across the network.

### Product Functions

go-nostr serves as an essential component of the Nostr network by facilitating event creation, management, and communication between clients and relays. It provides core functionalities and supports selected NIPs to streamline the development of Nostr-compatible applications and contribute to the network's growth.

### User Characteristics

The primary users of go-nostr include contributors, developers, and end-users. Contributors are those who actively participate in the development process, while developers leverage the software to create Nostr-compatible applications. End-users, on the other hand, interact with the applications built upon go-nostr to access the Nostr network.

### Constraints

go-nostr must adhere to the specifications outlined in the Nostr protocol and maintain compatibility with other Nostr implementations. Furthermore, it must be efficient, secure, and robust while providing a scalable and user-friendly solution.

### Assumptions and Dependencies

The development of go-nostr assumes that the Nostr protocol will continue to evolve and incorporate new features and improvements. As such, the software's design must be flexible enough to accommodate these changes. Additionally, go-nostr depends on the stability and reliability of the Nostr network and its underlying infrastructure.

### Apportioning of Requirements

This document focuses on the essential requirements for go-nostr. However, as the Nostr protocol evolves, new requirements may emerge. To ensure the software remains up-to-date and relevant, these requirements will be allocated to subsequent releases or iterations of the software.

## Specific Requirements

### External Interfaces

#### Software Interfaces

##### gRPC

TBD

##### HTTP

TBD

##### Websocket

TBD

#### User Interfaces

TBD

### Functions

#### Metadata (Kind 0, NIP 1)

TBD

#### Short Text Note (Kind 1, NIP 1)

TBD

#### Recommend Relay (Kind 2, NIP 1)

TBD

#### Contacts (Kind 3, NIP 2)

TBD

#### Encrypted Direct Messages (Kind 4, NIP 4)

TBD

#### Event Deletion (Kind 5, NIP 9)

TBD

#### Reposts (Kind 6, NIP 18)

TBD

#### Reaction (Kind 7, NIP 25)

TBD

#### Badge Award (Kind 8, NIP 58)

TBD

#### Channel Creation (Kind 40, NIP 28)

TBD

#### Channel Metadata (Kind 41, NIP 28)

TBD

#### Channel Message (Kind 42, NIP 28)

TBD

#### Channel Hide Message (Kind 43, NIP 28)

TBD

#### Channel Mute User (Kind 44, NIP 28)

TBD

#### Reporting (Kind 1984, NIP 56)

TBD

#### Zap Request (Kind 9734, NIP 57)

TBD

#### Zap (Kind 9735, NIP 57)

TBD

#### Mute List (Kind 10000, NIP 51)

TBD

#### Pin List (Kind 10001, NIP 51)

TBD

#### Relay List Metadata (Kind 10002, NIP 65)

TBD

#### Client Authentication (Kind 22242, NIP 42)

TBD

#### Nostr Connect (Kind 24133, NIP 46)

TBD

#### Categorized People List (Kind 30000, NIP 51)

TBD

#### Categorized Bookmark List (Kind 30001, NIP 51)

TBD

#### Profile Badges (Kind 30008, NIP 58)

TBD

#### Badge Definition (Kind 30009, NIP 58)

TBD

#### Create or update a stall (Kind 30017, NIP 15)

TBD

#### Create or update a product (Kind 30018, NIP 15)

TBD

#### Long-form Content (Kind 30023, NIP 23)

TBD

#### Application-specific Data (Kind 30078, NIP 78)

TBD

### Performance Requirements

TBD

### Design Constraints

TBD

### Software System Attributes

TBD

### Other Requirements

TBD

