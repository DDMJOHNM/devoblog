---
title: "ADRs"
date: 2022-10-18T16:56:59+13:00
draft: false
---

https://adr.github.io/
https://github.com/adr

adrs 
decision log
A “lightweight” ADR consists of title, status, context, decision, and consequences (according to @mtnygard).

# Decision record template by Michael Nygard

This is the template in [Documenting architecture decisions - Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions).
You can use [adr-tools](https://github.com/npryce/adr-tools) for managing the ADR files.

We think that the considered options with their pros and cons are crucial to understand the reason of a chosen option. MADR — The Markdown Any/Architecture Decision Records (MADR: [ˈmæɾɚ]) in this ADR organization includes such tradeoff analysis information.


In each ADR file, write these sections:

# Title

## Status

What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision


What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?

In short, the Y-statement is as follows:

In the context of <use case/user story u>, facing <concern c> we decided for <option o> to achieve <quality q>, accepting <downside d>.

