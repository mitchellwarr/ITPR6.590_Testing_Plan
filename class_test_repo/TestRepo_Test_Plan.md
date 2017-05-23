# Test Repo Test Plan


## Test Plan Identifier

Plan number JDDD - 930

Written by Mitchell Warr


## Introduction

This is essentially the executive summary part of the plan.

Identify the Scope of the plan 

As this is the "Executive Summary" keep information brief and to the point.


## Test Items (Functions)

#

### add product to database
- create product
- add product to database
- retrieve product from database
- check that the retrieved has the same values as the product

#

### remove product from database
- create product
- add product to database
- check that the product was actually added
  - retrieve product from database
  - check that the retrieved has the same values as the product
- remove product from database
- retrieve product from database
- check that the retrieved no longer exists

#

### retrieve product from database by ID
- create product
- add product to database
- retrieve product by its ID
- check that retrieved is equal to the product

#

### retrieve all product from database
- create multiple differetn products
- add products to database
- retrieve products from databse
- check that retrieved is equal to the products

#

### retrieve product from database by category
- create multiple products with different categories
- add products to database
- retrieve products from databse by category
- check that retrieved is equal to the products of only that category

#

### retrieve all categories from database
- create multiple products with different categories
- add products to database
- retrieve categories from databse
- check that retrieved is equal to the products categories

#

### upate product in database
- create product
- add product to database
  - retrieve product from database
  - check that the retrieved has the same values as the product
- update product in databse with different etails
- retrieve product from database
- check that the retrieved has the same values as the product

#
