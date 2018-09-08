# om_admin
om_admin is a lightweight MVC  pattern with control roles management API base framework in Go (Golang) for building fast, scalable and robust database-driven web applications.

# Features:
 Postgres, MySQL, SQLite and Foundation database support
 Modular (you can choose which components to use)
 Middleware support. All compatible Middleware works out of the box
 Gopher spirit (write golang, use all the golang libraries you like)
 Lightweight. Only MVC
 Multiple configuration files support (currently json, yaml, toml and hcl)

# Overview
om is a lightweight MVC framework. It is based on the principles of simplicity, relevance and elegance.

Simplicity. The design is simple, easy to understand, and doesn't introduce many layers between you and the standard library. It is a goal of the project that users should be able to understand the whole framework in a single day.

Relevance. om doesn't assume anything. We focus on things that matter, this way we are able to ensure easy maintenance and keep the system well-organized, well-planned and sweet.

Elegance. om uses golang best practises. We are not afraid of heights, it's just that we need a parachute in our backpack. The source code is heavily documented, any functionality should be well explained and well tested.

Motivation
After two years of playing with golang, I have looked on some of my projects and asked myself: "How golang is that?"

So, om is my reimagining of lightweight MVC, that maintains the golang spirit, and works seamlessly with the current libraries.

# Installation
om works with Go 1.4+

# Tutorials
Contributing
Start with clicking the star button to make the author and his neighbors happy. Then fork the repository and submit a pull request for whatever change you want to be added to this project.

# Application Package Flow
	-> 	main 
	-> 	handlers 
	-> 	utilities
	-> 	controllers 
	-> 	models

1. The main.go serves as the main function, initializing main method is the entry point the base resources needed to run build.
2. The handlers examines the HTTP request to determine what should be done with it. Before the application controller is loaded, the HTTP request and any user submitted data is filtered for security.
3. The Controller loads the model, core libraries, utilities, and any other package needed to process the specific request.

Design and Architectural Goals
maximum performance, capability, and flexibility in the smallest, lightest possible package.


If you have any questions, just open an issue.

# Author
Manish Paneri

# Blog : 
https://manishpaneri.blogspot.com
