# Sal-app
Shop Anything Lagos (SAL) is a SAAS company that lists products from different merchants. Sal-app implements SAL using REST APIs, basic CRUD operations and an inbuilt data structure. 

## API Documentation:
* Click [here](https://documenter.getpostman.com/view/29097143/2sA2rCVNHm) to view the documentation for this API on Postman

## How to run:
* Open git bash and navigate to your preferred directory
* Run `git clone git@github.com:scottdiddy/sal-app.git`
* Open the cloned repo in vscode
* Run `go mod tidy` in your terminal
* Now you can run the code using `go run src/main.go`

## Description:
* **Architecture:** Due to the nature of the application (monolithic), **Layered Architecture** was used.
* **Why Layered Architecture?** Layered Architecture was chosen for its clear seperation of concerns which enhances maintainability, scalability, and reusability of components within the system.
* **Benefits:**
  - Modularity: Allows for separation of concerns into distinct layers, such as presentation, business logic, and data access, facilitating easier understanding and maintenance of the system as changes in one layer are less likely to affect another layer.
  - Ease of Use: The basic concept of layered architecture is relatively straightforward.
  - Scalability: Scalability is improved as individual layers can be scaled independently, allowing for better performance optimization.
  - Reusability: Resuable components can be organized into specific layers and will be declared open that is accessible to all other layers.
  - Testability: Each layer can be tested independently, making it easier to write unit tests and ensure the reliability of the system.
<img width="1809" alt="Layered Architecture Sal-App" src="https://github.com/scottdiddy/sal-app/assets/141838693/163009fc-4a3c-448e-b397-729195c7476e">

## Database Implementation:
* **ERD Design:**
  - Relationship: As shown in the ERD, the database would be implemented as a one to many relationship (mandatory one to mandatory many) between the two entities - Merchant and Products. One to many because a merchant can have different products but a product can have only one merchant. And also it is mandatory because a product must have at least one merchant and a merchant must have at least one product.
* **Justification for ERD Design:**
  - This ERD utilizes two tables, Merchant and Product, with a one-to-many relationship between them. This is suitable for efficient data storage and retrieval, especially when dealing with a large number of merchants and products.
  - The merchant_id in the Product table acts as a foreign key referencing the primary key of the Merchant table, ensuring data integrity and enabling efficient retrieval of products associated with a specific merchant.

* **Considerations for Performance for Large Data:**
  - Indexing: Depending on the queries performed frequently, additional indexes may be added on those columns to optimize query performance.
  - Partitioning: Also partitioning the Product table based on the merchant_id or other relevant criteria can optimize performance by focusing queries on specific subsets of data.
  - Caching: Implementing caching mechanisms can reduce database load by serving frequently accessed data from memory.
 * **Choice of Database:**
   - Type: An SQL database like PostgreSQL is recommended.
   - **Why PostgreSQL?**
     1) Data Involved: PostgreSQL is recommended considering that the data involved is structured and can be arranged efficiently into tables with rows and columns. 
     2) Scalability: Since the platform anticipates significant growth in the number of merchants and products, a highly scalable database like PostgreSQL is best.
     3) Performance: To enhance performance (speed and dependability), choosing a database with efficient query processing capabilities and robust indexing features is crucial.
     4) Cost: Depending on budget constraints, open-source options like PostgreSQL might be preferred over proprietary solutions.

![ERD-Sal-App](https://github.com/scottdiddy/sal-app/assets/141838693/6b5812fa-0f03-47dd-980b-88d26c1be4b2)
