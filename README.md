# ATS
This hosts code for Automated Token System

Folder structure  
/app containes server-side code  
/assets holds all raw digital assets created for this project  
/cmd is the entry point to the server  
/config allows different configs to be loaded during initialization, the most important one being db urls  
/frontend contains raw code for the client webapp  
/frontend/public contains compiled code for the client webapp  
  

Infrastructure  
Backend: GoLang (Hosted on Heroku)   
FrontEnd: svelte/javascript, SPA  
API: GraphQL  
WebHooks: (GraphQL, JSON, REST) either  
Database: MongoDB (Hosted on Atlas)  
Design Concept: (Domain Driven Design + Test Driven Design), Agile  

Golang DDD structure:
The entry point  
Main Types:  
meter  
token  
user  

Every type is isolated in it's own package  
Every package associated with these elemets has 2 go files: interactor and repository  
Respository contains a private pointer to the database implementation, as well as methods of all the possible operations on this type  
This allows for the package to implement the database locally, and the only externally visible methods are the operations.  
Interactor contains all the higher-level abstractions of the possible operations that might include interacting with other packages.  
This allows for inter-package communication as well as grouping several operations into one function. The same method might be visible from different packages  
There is one shared database implementation that is shared across all the packages.