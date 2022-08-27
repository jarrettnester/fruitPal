WELCOME TO FRUITPAL!!

I hope you enjoy this application.

Install instructions are as follows:

1. Install the latest Golang version
2. Install Postgres database
3. Add the following database to Postgres:

HOST = "localhost"
PORT = "5432"
USER = "jarrettnester"
PASSWORD = "janejohn123"
DBNAME = "susannester"

4. Add the following tables and populate data:

TABLE commodities
COLUMN INT id
COLUMN VARCHAR name

TABLE countries
COLUMN INT id
COLUMN VARCHAR name
COLUMN VARCHAR country_code

TABLE data_sources
COLUMN INT id
COLUMN VARCHAR name

TABLE data_snapshots
COLUMN INT id
COLUMN TIMESTAMP date_added
COLUMN INT data_source_id

TABLE data_snapshot_values
COLUMN INT id
COLUMN INT commodity_id
COLUMN INT country_id
COLUMN INT data_snapshot_id
COLUMN DECIMAL variable_cost
COLUMN DECIMAL fixed_overhead

5. Navigate in the command line to the apiServerLayer directory and execute:

    A. go build
    B. ./apiServerLayer

6. Navigate in the command line to the client directory and execute:

    npm run start
   
7. Open your browser to localhost:4200 and you should see the following:

    A. a select list of countries in your database
    B. a list of total cost based on hard coded values for a price sent to the 
   API of 2.45 and tonnage of 56 and the commodity id of 1. (NOTE:  this is currently
   hard coded in the service class of the client)
   
FUTURE UPCOMING RELEASES OF FRUITPAL INCLUDE:

1. Inputted values that look nice creating a dynamic user experience for 
fruit total cost searching.
2. Unit tests for functions within the front and back end of the application.
3. Much much more!

Thanks again for trying out Fruitpal!

Just reach out to the CEO/CFO/Founder/PM/Quality Control Specialist/UX extraordinaire
Jarrett Nester with any further inquires.
   
