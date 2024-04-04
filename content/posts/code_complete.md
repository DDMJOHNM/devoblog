---
title: "Code Complete - Steve McConnell"
date: 2024-01-01T16:56:59+13:00
draft: false
---

## Chapter One
### Welcome to Software Construction

■ Software construction is the central activity in software development; construc- tion is the only activity that’s guaranteed to happen on every project.
■ The main activities in construction are detailed design, coding, debugging, inte- gration, and developer testing (unit testing and integration testing).
■ Other common terms for construction are “coding” and “programming.”
■ The quality of the construction substantially affects the quality of the software.
■ In the final analysis, your understanding of how to do construction determines how good a programmer you are, and that’s the subject of the rest of the book.


## Chapter Two 
### Metaphors for a Richer Understanding of Software Development


■ Metaphors are heuristics, not algorithms. As such, they tend to be a little sloppy.
■ Metaphors help you understand the software-development process by relating it
to other activities you already know about.
■ Some metaphors are better than others.
■ Treating software construction as similar to building construction suggests that careful preparation is needed and illuminates the difference between large and small projects.
■ Thinking of software-development practices as tools in an intellectual toolbox suggests further that every programmer has many tools and that no single tool is right for every job. Choosing the right tool for each problem is one key to being an effective programmer.
■ Metaphors are not mutually exclusive. Use the combination of metaphors that works best for you.

## Chapter Three 
### Measure Twice Cut Once: Upstream Prerequisites
Paying attention to quality is the best way to improve productivity
Risk Reduction 
Upstream work 
Construction Prerequistes
If Requirements are contaminated they contaminate the architecture and that contaminates construction 
cost to fix a defect rises from the time its introduced to when its detected

3.1 Importance of Prerequisites
3.2 Determine the Kind of Software You’re Working On
3.3 Problem-Definition Prerequisite
3.4 Requirements Prerequisite
25% changes in requirements during development 
3.5 Architecture Prerequisite

Handling Requirements Changes During Construction

design—architecture refers to design constraints that apply systemwide, whereas high- level design refers to design constraints that apply at the subsystem or multiple-class level, but not necessarily systemwide.

Achitectural Components 
Program Organization
Major Classes
Data Design 
Business Rules
User Interface Design
Resource Management
Security
Performance 
Scalability
Interoperability
Internationalization/Localization
Input/Output
Error Processing **
Fault Tolerance 
Architectural Feasibility
Overengineering
Buy vs Build Decsions
Reuse Decisions
Change Strategy
General Architectural Quality
General Software-Development Approaches
Many books are available that map out different approaches to conducting a software project. Some are more sequential, and some are more iterative

■ The overarching goal of preparing for construction is risk reduction. Be sure your preparation activities are reducing risks, not increasing them.
■ If you want to develop high-quality software, attention to quality must be part of the software-development process from the beginning to the end. Attention to quality at the beginning has a greater influence on product quality than atten- tion at the end.
■ Part of a programmer’s job is to educate bosses and coworkers about the soft- ware-development process, including the importance of adequate preparation before programming begins.
■ The kind of project you’re working on significantly affects construction prereq- uisites—many projects should be highly iterative, and some should be more sequential.
■ If a good problem definition hasn’t been specified, you might be solving the wrong problem during construction.
Key Points 59
Checklist: Upstream Prerequisites
❑ Have you identified the kind of software project you’re working on and tai- lored your approach appropriately?
❑ Are the requirements sufficiently well defined and stable enough to begin construction? (See the requirements checklist for details.)
❑ Is the architecture sufficiently well defined to begin construction? (See the architecture checklist for details.)
❑ Have other risks unique to your particular project been addressed, such that construction is not exposed to more risk than necessary?
60 Chapter 3: Measure Twice, Cut Once: Upstream Prerequisites
■ If good requirements work hasn’t been done, you might have missed important details of the problem. Requirements changes cost 20 to 100 times as much in the stages following construction as they do earlier, so be sure the requirements are right before you start programming.
■ If a good architectural design hasn’t been done, you might be solving the right problem the wrong way during construction. The cost of architectural changes increases as more code is written for the wrong architecture, so be sure the archi- tecture is right, too.
■ Understand what approach has been taken to the construction prerequisites on your project, and choose your construction approach accordingly.

## Chapter 4
### Key Construction Decisions
Every programming language has strengths and weaknesses. Be aware of the specific strengths and weaknesses of the language you’re using.
■ Establish programming conventions before you begin programming. It’s nearly impossible to change code to match them later.
■ More construction practices exist than you can use on any single project. Con- sciously choose the practices that are best suited to your project.
■ Ask yourself whether the programming practices you’re using are a response to the programming language you’re using or controlled by it. Remember to pro- gram into the language, rather than programming in it.
■ Your position on the technology wave determines what approaches will be effec- tive—or even possible. Identify where you are on the technology wave, and adjust your plans and expectations accordingly.


## Chapter 5
### Design In Construction 

Here’s a summary of major design heuristics:
■ Find Real-World Objects
■ Form Consistent Abstractions
■ Encapsulate Implementation Details ■ Inherit When Possible
■ Hide Secrets (Information Hiding) ■ Identify Areas Likely to Change
■ Keep Coupling Loose
■ Look for Common Design Patterns

The following heuristics are sometimes useful too:
■ Aim for Strong Cohesion
■ Build Hierarchies
■ Formalize Class Contracts
■ Assign Responsibilities
■ Design for Test
■ Avoid Failure
■ Choose Binding Time Consciously
■ Make Central Points of Control
■ Consider Using Brute Force
■ Draw a Diagram
■ Keep Your Design Modular

5.4 Design Practices

