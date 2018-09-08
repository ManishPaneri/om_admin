# Installation

`om_admin` works with Go 1.4+

# Overview
`om_admin` is a lightweight MVC framework. It is based on the principles of simplicity, relevance and elegance.

* Simplicity. The design is simple, easy to understand, and doesn't introduce many layers between you and the standard library. It is a goal of the project that users should be able to understand the whole framework in a single day.

* Relevance. `om_admin` doesn't assume anything. We focus on things that matter, this way we are able to ensure easy maintenance and keep the system well-organized, well-planned and sweet.

* Elegance. `om_admin` uses golang best practises. We are not afraid of heights, it's just that we need a parachute in our backpack. The source code is heavily documented, any functionality should be well explained and well tested.


#Features
Features in and of themselves are a very poor way to judge an application since they tell you nothing about the user experience, or how intuitively or intelligently it is designed. Features don’t reveal anything about the quality of the code, or the performance, or the attention to detail, or security practices. The only way to really judge an app is to try it and get to know the code.
	- 	Model-View-Controller Based System
	- 	Extremely Light Weight
	- 	Full Featured database classes with support for several platforms.
	- 	Query Builder Database Support
	- 	Session Management
	- 	Email Sending Function.
	- 	File Uploading Class
	- 	Localization
	- 	Pagination
	- 	Data Encryption
	- 	API Logging
	- 	Application Profiling
	- 	Flexible URI Routing
	- 	Support for Package Extensions
	- 	Library of "utilities” Package


#Application Package Flow
	-> 	main 
	-> 	handlers 
	-> 	controllers 
	-> 	models
	-> 	redshifts
	-> 	utilities

1. The main.go serves as the main function, initializing main method is the entry point the base resources needed to run build.
2. The handlers examines the HTTP request to determine what should be done with it. Before the application controller is loaded, the HTTP request and any user submitted data is filtered for security.
3. The Controller loads the model, core libraries, utilities, and any other package needed to process the specific request.


#Design and Architectural Goals
maximum performance, capability, and flexibility in the smallest, lightest possible package.

To meet this goal we are committed to benchmarking, re-factoring, and simplifying at every step of the development process, rejecting anything that doesn’t further the stated objective.

From a technical and architectural standpoint, om_admin was created with the following objectives:

Dynamic Instantiation. In om_admin, components are loaded and routines executed only when requested, rather than globally. No assumptions are made by the system regarding what may be needed beyond the minimal core resources, so the system is very light-weight by default. The events, as triggered by the HTTP request, and the controllers and views you design will determine what is invoked.
Loose Coupling. Coupling is the degree to which components of a system rely on each other. The less components depend on each other the more reusable and flexible the system becomes. Our goal was a very loosely coupled system.
Component Singularity. Singularity is the degree to which components have a narrowly focused purpose. In om_admin, each class and its functions are highly autonomous in order to allow maximum usefulness.