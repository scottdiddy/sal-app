# Sal-app
Shop Anything Lagos (SAL) is a SAAS company that lists products from different merchants. Sal-app implements SAL using REST APIs and basic CRUD operations and an inbuilt data structure. 

## Description:
* **Architecture:** Due to the nature of the application (monolithic), **Layered Architecture** was used.
* **Why Layered Architecture?** Layered Architecture was chosen for its clear seperation of concerns which enhances maintainability, scalability, and reusability of components within the system.
* **Benefits:**
  - Modularity: Allows for separation of concerns into distinct layers, such as presentation, business logic, and data access, facilitating easier understanding and maintenance of the system as changes in one layer are less likely to affect another layer.
  - Ease of Use: The basic concept of layered architecture is relatively straightforward.
  - Scalability: Scalability is improved as individual layers can be scaled independently, allowing for better performance optimization.
  - Reusability: Resuable components can be organized into specific layers and will be declared open that is accessible to all other layers.
  - Testability: Each layer can be tested independently, making it easier to write unit tests and ensure the reliability of the system.
<img width="1809" alt="Layered Architecture Sal-App" src="https://github.com/scottdiddy/sal-app/assets/141838693/8ce8a4a4-200c-433e-85d8-f88f25e5046f">

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
