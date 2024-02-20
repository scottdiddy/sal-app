# sal-app
Shop Anything Lagos (SAL) is a SAAS company that lists products from different merchants. Sal-app implements SAL using REST APIs and basic CRUD operations and an inbuilt data structure. 

## Todo:
* Create Product and Merchant models. Merchant models will have the Merchant id and all the associated SKU id. That is a merchant and all his/her products.
* Initialize middlewares:
  - Body Validator - checks the incoming request body and ensuresn certain fields are present.
* Define controllers for each routes:
  - Create a product.
  - Display all products listed by a merchant.
  - Edit a product.
  - Delete an existing product.
* Similarly define all services for each routes that performs the actual work.

## Usage:
* Open git bash and navigate to your preferred directory
* Run git clone git@github.com:scottdiddy/sal-app.git
* Open the clone repo in vscode and open your terminal
* Run go mod tidy
* cd into src
* Now you can run the code using go run main.go
