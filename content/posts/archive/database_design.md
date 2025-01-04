---
title: "Database Design"
date: 2022-11-13T16:56:59+13:00
draft: true
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

(Above from Coursera course)

https://www.mongodb.com/compare/relational-vs-non-relational-databases
## Relational versus NoSql

# Rdms 
-Relationships between table (Relational Database)

## Advantages 
### Acid Compliant
-Atomicity (All or nothing Update) , Consistency and Durability - standard that guarantees the reliability of Database transactions.

### Data Accurracy
primary and foreign keys ensure no duplication

### Nomalization
Ensuring data is organized in such a way that data anomalies are reduced or eliminated also reduces storage cost

### Simplicity
English like syntax of sql 

## Disadvantages

### Scalability
- Vertical Scaling improve hardware in the machine.
### Flexibility
- Schema is rigid for relational DB

### Performance 
- tightly linked to the complexity of tables and number
as well amount of data in each table.

## Non Relational Database or noSQL 
- Any db that doesnt use tables great for the cloud horizontal scaling involved (Distributed across servers)

### Document Databases
Documents individual units 
Json like structures that support datatypes.
db.product.find({"_id": 23}, {productName: 1, price: 1})
filter collections of data 
eg mongo DB

### Key Value Database 
Data accessed from keys simplicity

### Graph Databases 
Most specialised for of non relational db 
they are flexible because new nodes and edges can easily be added.
not good at querying the whole as relation arent well defined.     

## Wide Columns Databases 
Considered two dimensional key value stores. 

# When to use a relational Database

- Project where data is predictable, structure, size and frequency.
- Nomalization can help reduce the size of data on the disk and limiting duplicate data.
- Are the best choice if relationships between enitites are important
- large dataset with a complex structure.

# When to use a non-relational database
- it your data needs to be flexible given size and shape and be open in the future
- Cloud and scaling 

``` 
Non-relational databases are suitable for both operational and transactional data.
They are more suitable for unstructured big data.
Non-relational databases offer higher performance and availability.
Flexible schema help non-relational databases store more data of varied types that can be changed without major schema changes.
```

## Relational Models
# Normalization 
## NORMAL FORMS
organise data efficiently in db
guidelines 

- avoid data redundancy
- data logically stored

### 1nf
table db or schema
1 - only atomic values
2 - No repeating groups
3 - each table must have a unique primary key

### 2nf
1 - in the first normal form
2 - no partial dependancies 
composite keys


### 3nf
1 - 2nf
2 - all non pk fields  are dependant on the pk

## Denormalization
used on previously normalised databases to increase perfomance 
adding redundant copies of data or by grouping data 
often motivated by scalability or performance in relational databses 
needinf to carry out very large numbers of read operations.
only decognised on a data model otherwised normalised.

eg join operatons on a normalised design may be prohibitively slow
- sql indexed views
- optimise query response 

denormalise logical data design 
so it doesnt become inconsistant
using constarints 

 
