package shoppingproject.dao;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import static org.junit.Assert.*;
import org.junit.runners.Parameterized;
import org.junit.runner.RunWith;
import shoppingproject.domain.Product;

/**
 *
 * @author <Your name here>
 */
@RunWith(Parameterized.class)
public class DaoTester {

    /**
     * To test our DAO we need to create an instance of the interface class 
     * IDAO and then use a constructor to inject an implementation of the interface.
     */

    /**
     * To ensure that each test case runs with isolated test data we should also create
     * 2-4 instances of products that we can add during the setUp and tearDown methods for
     * each of the test cases.
     */


    @Parameterized.Parameters
    public static Collection<?> daoObjectsToTest() {
        return Arrays.asList(new Object[][]{
            {new ListDao()},
        });
    }

    @BeforeClass
    public static void setUpClass() {
    }

    @AfterClass
    public static void tearDownClass() {
    }

    @Before
    public void setUp() {
        //Here we need to create instances of products with data that we can easily test againts
        //Call the IDAO method to add each of the product instances.
     
        dbTest.addProduct(product1);
        dbTest.addProduct(product2);
    }

    @After
    public void tearDown() {
        //Call the IDAO method to remove each of the products added in setUp
    }

    @Test
    public void testDaoGetProductList() {
        //fetch all products in the database

        //test that the two products added in setUp exist after fetching them

        //ensure that we have only added the two test products


        //find products in our returned collection and ensure that all details
        //of the fetched products match the products details used in setUp method
    }

    @Test
    public void testDaoFindById() {
        //Fetch product from db using the local Product models get ID method
 
        //Test local product1 Model against product1 stored in database
        
        /**
         * Note that assertEquals uses equals interface and utilises products equals implemented version
         * and this implementation only compares ID, so we should test against all of the product attributes
         */
 
        //Testing against all product attributes

    }
    
    @Test
    public void testGetCategoryList() {
        //Fetch a list of categories

        //test that only two categories are returned or the amount of categories you added during setUp method

        //test that 1st product category matches fetched category from database

        //test that 2nd product category matches fetched category from database
    }

    @Test
    public void testAddAndDeleteProduct() {
        //create an instance of Product to test add and delete methods

        //add product to database

        //retrieve product after adding using productID and test it is return by comparing to the local instance

        //delete product using instance of saveProduct and test it is not returned
    }

     @Test
    public void testUpdateProduct() {
        //Create two local variable strings for a product name and description that will be used to test updating.

        //fetch product using the product ID of one of the products used in setUp method

        //set the new product name

        //set the new product description

        //with new name and description update the product with updateProduct

        //fetch the product after an update and check the changes have persisted

        //check name has change

        //check description has changed
    }

     @Test
    public void testGetByCategory() {
        //Create local variables which are copies of our test product categories
        
        //Use the getByCategory dao method to fetch a collection of each category
        
        /*
        there should be atleast two different categories with 1 product in each category that
        we can use to test that only one product is returned for each category.

        Adjust this for how ever many products/categories added.
        */

        //Test that the fetched products category matches our local product category 
        
        //we should also test that the returned products dont overlap into wrong categorys
    }
}
