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

- [RFC 8615: Well-Known Uniform Resource Identifiers (URIs)](https://www.rfc-editor.org/rfc/rfc8615)

### Overvierw

TBD

## Overall Description

### Product Perspective

The go-nostr software is designed as a part of the Nostr network, which is a decentralized and extensible social protocol. The software acts as a bridge between Nostr-compatible applications, providing both relay and client functionalities to enable seamless communication across the network.

### Product Functions

go-nostr serves as a component of the Nostr network by facilitating event creation, management, and communication between clients and relays. It provides core functionalities and supports selected NIPs to streamline the development of Nostr-compatible applications and contribute to the network's growth.

### User Characteristics

The primary users of go-nostr include contributors, developers, and end-users. Contributors are those who actively participate in the development process, while developers leverage the software to create Nostr-compatible applications. End-users, on the other hand, interact with the applications built upon go-nostr to access the Nostr network.

### Constraints

go-nostr must adhere to the specifications outlined in the Nostr protocol and maintain compatibility with other Nostr implementations. Furthermore, it must be efficient, secure, and robust while providing a scalable and user-friendly solution.

| Status | NIP                                                                        |
|--------|----------------------------------------------------------------------------|
|        | NIP-01: Basic protocol flow description                                    |
|        | NIP-02: Contact List and Petnames                                          |
|        | NIP-03: OpenTimestamps Attestations for Events                             |
|        | NIP-04: Encrypted Direct Message                                           |
|        | NIP-05: Mapping Nostr keys to DNS-based internet identifiers               |
|        | NIP-06: Basic key derivation from mnemonic seed phrase                     |
|        | NIP-07: window.nostr capability for web browsers                           |
|        | NIP-08: Handling Mentions --- unrecommended: deprecated in favor of NIP-27 |
|        | NIP-09: Event Deletion                                                     |
|        | NIP-10: Conventions for clients' use of e and p tags in text events        |
|        | NIP-11: Relay Information Document                                         |
|        | NIP-12: Generic Tag Queries                                                |
|        | NIP-13: Proof of Work                                                      |
|        | NIP-14: Subject tag in text events.                                        |
|        | NIP-16: Event Treatment                                                    |
|        | NIP-18: Reposts                                                            |
|        | NIP-19: bech32-encoded entities                                            |
|        | NIP-20: Command Results                                                    |
|        | NIP-21: nostr: URL scheme                                                  |
|        | NIP-22: Event created_at Limits                                            |
|        | NIP-23: Long-form Content                                                  |
|        | NIP-25: Reactions                                                          |
|        | NIP-26: Delegated Event Signing                                            |
|        | NIP-27: Text Note References                                               |
|        | NIP-28: Public Chat                                                        |
|        | NIP-33: Parameterized Replaceable Events                                   |
|        | NIP-36: Sensitive Content                                                  |
|        | NIP-39: External Identities in Profiles                                    |
|        | NIP-40: Expiration Timestamp                                               |
|        | NIP-42: Authentication of clients to relays                                |
|        | NIP-45: Counting results                                                   |
|        | NIP-46: Nostr Connect                                                      |
|        | NIP-50: Keywords filter                                                    |
|        | NIP-51: Lists                                                              |
|        | NIP-56: Reporting                                                          |
|        | NIP-57: Lightning Zaps                                                     |
|        | NIP-58: Badges                                                             |
|        | NIP-65: Relay List Metadata                                                |
|        | NIP-78: Application-specific data                                          |

### Assumptions and Dependencies

The development of go-nostr assumes that the Nostr protocol will continue to evolve and incorporate new features and improvements. As such, the software's design must be flexible enough to accommodate these changes. Additionally, go-nostr depends on the stability and reliability of the Nostr network and its underlying infrastructure.

### Apportioning of Requirements

This document focuses on the essential requirements for go-nostr. However, as the Nostr protocol evolves, new requirements may emerge. To ensure the software remains up-to-date and relevant, these requirements will be allocated to subsequent releases or iterations of the software.

## Specific Requirements

### External Interfaces

#### Software Interfaces

##### API

TBD

###### GET /.well-known/nostr.json?name=local-part

Returns the internet identifier for the given name query parameter.

###### GET /health

Returns the current health status of the service.

#### User Interfaces

TBD

##### Client

TBD

###### AccountSettingsPageComponent

TBD

###### AppearanceSettingsPageComponent

TBD

###### BackupSettingsPageComponent

TBD

###### CreateShortTextNoteComponent

TBD

###### HomePageComponent

TBD

###### ListNotificationsPageComponent

TBD

###### NetworkSettingsPageComponent

TBD

###### ProfilePageComponent

TBD

###### SignInPageComponent

TBD

###### SignUpPageComponent

TBD

##### Docs

TBD

### Functions

#### Get Health

Returns the current health [status] of the service implementing the healthcheck endpoint monitoring pattern.

#### Get Internet Identifier

Returns the internet identifier for given name and optionally relays, if enabled. Supports [NIP-05](https://github.com/nostr-protocol/nips/blob/master/05.md).

#### Publish Metadata Event

(Kind 0, NIP 1)

TBD

#### Publish Short Text Note Event

(Kind 1, NIP 1)

TBD

#### Publish Recommend Relay Event

(Kind 2, NIP 1)

TBD

#### Publish Contacts Event

(Kind 3, NIP 2)

TBD

#### Publish Encrypted Direct Messages Event

(Kind 4, NIP 4)

TBD

#### Publish Event Deletion Event

(Kind 5, NIP 9)

TBD

#### Publish Reposts Event

(Kind 6, NIP 18)

TBD

#### Publish Reaction Event

(Kind 7, NIP 25)

TBD

#### Publish Badge Award Event

(Kind 8, NIP 58)

TBD

#### Publish Channel Creation Event

(Kind 40, NIP 28)

TBD

#### Publish Channel Metadata Event

(Kind 41, NIP 28)

TBD

#### Publish Channel Message Event

(Kind 42, NIP 28)

TBD

#### Publish Channel Hide Message Event

(Kind 43, NIP 28)

TBD

#### Publish Channel Mute User Event

(Kind 44, NIP 28)

TBD

#### Publish Reporting Event

(Kind 1984, NIP 56)

TBD

#### Publish Zap Request Event

(Kind 9734, NIP 57)

TBD

#### Publish Zap Event

(Kind 9735, NIP 57)

TBD

#### Publish Mute List Event

(Kind 10000, NIP 51)

TBD

#### Publish Pin List Event

(Kind 10001, NIP 51)

TBD

#### Publish Relay List Metadata Event

(Kind 10002, NIP 65)

TBD

#### Publish Client Authentication Event

(Kind 22242, NIP 42)

TBD

#### Publish Nostr Connect Event

(Kind 24133, NIP 46)

TBD

#### Publish Categorized People List Event

(Kind 30000, NIP 51)

TBD

#### Publish Categorized Bookmark List Event

(Kind 30001, NIP 51)

TBD

#### Publish Profile Badges Event

(Kind 30008, NIP 58)

TBD

#### Publish Badge Definition Event

(Kind 30009, NIP 58)

TBD

#### Publish Create or update a stall Event

(Kind 30017, NIP 15)

TBD

#### Publish Create or update a product Event

(Kind 30018, NIP 15)

TBD

#### Publish Long-form Content Event

(Kind 30023, NIP 23)

TBD

#### Publish Application-specific Data Event

(Kind 30078, NIP 78)

TBD

### Performance Requirements

TBD

### Design Constraints

TBD

### Software System Attributes

TBD

### Other Requirements

TBD
