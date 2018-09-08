#Create New API 
You now know how you can read data from a database using om_admin, but you haven’t written any API permission information to the database yet. In this section you’ll expand your API Permission controller and model created earlier to include this functionality.

#Routing
main function in implement to handler event to call handler function

#handlers
To check API End Point permission data into the database where you can check request user having permission or not the information to be handle. 
This means validation library to do this.


#Controllers
You’re going to controllers package here, check whether the handler call function  create and whether the function data passed the validation rules. You’ll use the exist function format.

The code above adds a lot of functionality. The first few lines load the form utilites package library. After that, rules for the validation are set. The set_rules() method takes three arguments; the name of the input field, the name to be used in error messages, and the rule. In this case the title and text fields are required.

om_admin has a  create return response struct in library as demonstrated above. You can read more about this library here.

Continuing down, you can see a condition that checks whether the form validation ran successfully. If it did not, the form is displayed, if it was submitted and passed all the rules, the model is called. After this, a view is loaded to display a success message. 

#Model
The only thing that remains is writing a method that writes the data to the database. You’ll use the Query Builder class to insert the information and use the input library to get the posted data. Open up the model created earlier and add the following:


#Utilities
Route Function (route.go):
	The URI Routing function provides methods that help you retrieve information from your URI strings. If you use URI routing, you can also retrieve information about the re-routed segments.

Logging Function (logging.go):
	The API Log for store in database yet, API request and response store in database log table and S3 bucket store import file handle it.

Config Function (helper.go):
	The Config Function provides a means to retrieve configuration preferences. These preferences can come from the default config file or from your own custom config files.

Encryption Library (logic.go):
	The Encryption Library provides two-way data encryption. To do so in a cryptographically secure way, it utilizes go extensions that are unfortunately not always available on all systems. You must meet one of the following dependencies in order to use this library:
