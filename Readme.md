This API exposes 5 different end points


POST /contacts - Adds a contact
GET /contacts - Gets all contacts
GET /contacts/{contactID} - Gets a specific contact
PUT /contacts/{contactID} - Updates a speific contact
DELETE /contacts/{contactID} - Deletes a speific contact


The database for this operation has been mimicked to a file called Repo.go


There has been a unity which was built the utility converts the current contacts in the repo to a CSV called ./contact.csv


In order to run the program

go to the root folder and run

go build

and then run

go run *


This will spin up your server on localhost:8080

you can access them by the following endpoints.

localhost:8080/contacts

Make sure you have the following parameters in the header

charset = UTF-8
Content-Type= application/json



