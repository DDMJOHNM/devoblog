---
title: "Database Design"
date: 2022-11-13T16:56:59+13:00
draft: false
---

## Entity Relationship Models

DBMS
- data
- information
- knowledge 
- intelligence (value of data) 

From data to actionable insights 

File System Solution
- Separated, isolated data and data duplication
- Dependancy on applications 
- Incompatible files
- Lack of data sharing 

Database Solution
- A structure that contains data
- Organise data through entities, attributes, relationships
- Collection of programs that users interact with a database.

Popular DBMS
- Microsoft Access
- MySql
- Progres
- MongoDb (NoSql) 

Advantages 
- more information from same amount of data 
- share our data 
- balancing conflicting requirements 
- controlling redundacy 
- consistency 
- Referential integrity
- expanding security
- increasing productivity
- providing data independence.


Disaadvantages
- Complex
- Greater Impact of Failure 
- More differcult recovery
- Larger file size (Also structure, user and application)

When is a DBMS not appropriate - when advantages outweigh disadvantages
- Database small with simple structure
- Applications are simple and relatively static
- Applications have real time requirements 
- Concurrent, multi access to data not required  

Contents of a DBMS

User Data
- tables, relationships
- columns or row db structure

Meta Data 
- How to Describe the data 
- System Tables 
- Indexes access records without going through entire table 

Applications Metadata 
- Storage facility for forms 

Design a Relational Database Model 
Entity relationship model 
Describes
- Entities : class group abstract structure and no value (Object of Interest)
- Attributes (Identifier) : Characteristics of an entity of interest 
- Relationships among these entities : Instance - Object with values 
- Identifiers : special attribute used to identify a specific instance of an entity - artificial or natural (Created)

Relationship between two or more entities
eg employees and companies 
Domain knowledge consider answer in context
employee multiple companies 
degree of relations how many enities involved in the relationship
- Unary or Recursive - (1)
- Binary (2) - relational db cannot handle 

Cardinality - number of instances involved in the relationship (max cardinality or multiplicity)
- 1:n One to many
- N:1 Many to one 
- N:M Many to many 

Participation - Manditory or Optional (Optionality or Minimal Cardinatilty)
or require instance of entity here also there 

Entity Relationship Models 

Based on your understanding so far, what are the entities, attributes, identifiers, and relationships in the Entity Relationship Model? You can lay out the initial outline here, and the owner will meet with you to answer your questions.

{{< figure src="/images/dbcourse1.jpg" title="Entity RelationShip Model Pizza Franchise" >}}


Entity Relationship Diagram and Notation 








## Relational Models
## Normalization 
## Case Study 