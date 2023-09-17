package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/server"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testServer *httptest.Server
var testDb *gorm.DB

func TestMain(m *testing.M) {
	godotenv.Load()
	gormDialector := sqlite.Open(":memory:")
	store.InitSqlDb(gormDialector)
	testDb = store.GetDB()
	m.Run()
}

func TestClients(t *testing.T) {
	t.Run("Test list", func(t *testing.T) {
		setUp()
		clients := []models.Client{}
		testDb.Find(&clients)
		expectedBody := schemas.NewClients(clients)

		response, err := makeRequest("GET", "/clients", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody[0].Name != expectedBody[0].Name {
			t.Errorf("Expected name %s, got %s", expectedBody[0].Name, responseBody[0].Name)
		}

		tearDown()
	})

	t.Run("Test get", func(t *testing.T) {
		setUp()
		client := models.Client{}
		testDb.First(&client)
		expectedBody := schemas.NewClient(client)

		response, err := makeRequest("GET", "/clients/"+fmt.Sprint(client.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		client := schemas.CreateClient{}
		faker.FakeData(&client)
		expectedBody := schemas.Client{}
		expectedBody.Name = client.Name
		expectedBody.Email = client.Email

		response, err := makeRequest("POST", "/clients", client)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		client := models.Client{}
		testDb.First(&client)
		update := schemas.UpdateClient{}
		faker.FakeData(&update)
		expectedBody := schemas.Client{}
		expectedBody.Name = update.Name
		expectedBody.Email = update.Email

		response, err := makeRequest("PUT", "/clients/"+fmt.Sprint(client.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Client{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		client := models.Client{}
		testDb.First(&client)

		response, err := makeRequest("DELETE", "/clients/"+fmt.Sprint(client.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}

		tearDown()
	})

}

func TestProducts(t *testing.T) {

	t.Run("Test list", func(t *testing.T) {
		setUp()
		products := []models.Product{}
		testDb.Find(&products)
		expectedBody := schemas.NewProducts(products)

		response, err := makeRequest("GET", "/products", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody[0].Name != expectedBody[0].Name {
			t.Errorf("Expected name %s, got %s", expectedBody[0].Name, responseBody[0].Name)
		}

		tearDown()
	})

	t.Run("Test get", func(t *testing.T) {
		setUp()
		product := models.Product{}
		testDb.First(&product)
		expectedBody := schemas.NewProduct(product)

		response, err := makeRequest("GET", "/products/"+fmt.Sprint(product.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		product := schemas.CreateProduct{}
		faker.FakeData(&product)
		expectedBody := schemas.Product{}
		expectedBody.Name = product.Name
		expectedBody.Price = product.Price

		response, err := makeRequest("POST", "/products", product)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		product := models.Product{}
		testDb.First(&product)
		update := schemas.UpdateProduct{}
		faker.FakeData(&update)
		expectedBody := schemas.Product{}
		expectedBody.Name = update.Name
		expectedBody.Price = update.Price

		response, err := makeRequest("PUT", "/products/"+fmt.Sprint(product.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Product{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Name != expectedBody.Name {
			t.Errorf("Expected name %s, got %s", expectedBody.Name, responseBody.Name)
		}

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		product := models.Product{}
		testDb.First(&product)

		response, err := makeRequest("DELETE", "/products/"+fmt.Sprint(product.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}

		tearDown()
	})
}

func TestOrders(t *testing.T) {

	t.Run("Test list", func(t *testing.T) {
		setUp()
		orders := []models.Order{}
		testDb.Find(&orders)
		expectedBody := schemas.NewOrders(orders)

		response, err := makeRequest("GET", "/orders", nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := []schemas.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody[0].Quantity != expectedBody[0].Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody[0].Quantity, responseBody[0].Quantity)
		}

		tearDown()
	})

	t.Run("Test get", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)
		expectedBody := schemas.NewOrder(order)

		response, err := makeRequest("GET", "/orders/"+fmt.Sprint(order.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()
		responseBody := schemas.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)

		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}

		tearDown()
	})

	t.Run("Test create", func(t *testing.T) {
		setUp()
		product := models.Product{}
		testDb.First(&product)
		client := models.Client{}
		testDb.First(&client)
		order := schemas.CreateOrder{
			ProductID: product.ID,
			ClientID:  client.ID,
		}
		faker.FakeData(&order)
		expectedBody := schemas.Order{}
		expectedBody.Quantity = order.Quantity

		response, err := makeRequest("POST", "/orders", order)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}

		tearDown()
	})

	t.Run("Test update", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)
		update := schemas.UpdateOrder{}
		faker.FakeData(&update)
		expectedBody := schemas.Order{}
		expectedBody.Quantity = update.Quantity

		response, err := makeRequest("PUT", "/orders/"+fmt.Sprint(order.ID), update)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		responseBody := models.Order{}
		json.NewDecoder(response.Body).Decode(&responseBody)
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
		}

		if responseBody.Quantity != expectedBody.Quantity {
			t.Errorf("Expected quantity %d, got %d", expectedBody.Quantity, responseBody.Quantity)
		}

		tearDown()
	})

	t.Run("Test delete", func(t *testing.T) {
		setUp()
		order := models.Order{}
		testDb.First(&order)

		response, err := makeRequest("DELETE", "/orders/"+fmt.Sprint(order.ID), nil)
		if err != nil {
			t.Error(err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, response.StatusCode)
		}

		tearDown()
	})

}

func setUp() {
	clearDatabase()
	testApp := server.CreateApp()
	testServer = httptest.NewServer(testApp)
	populateTables()
}

func clearDatabase() {
	testDb.Exec("DROP SCHEMA public CASCADE")
	testDb.Exec("CREATE SCHEMA public")
	testDb.Exec("GRANT ALL ON SCHEMA public TO postgres")
	testDb.Exec("GRANT ALL ON SCHEMA public TO public")
}
func tearDown() {
	testServer.Close()
}

func makeRequest(method string, endpoint string, body interface{}) (*http.Response, error) {
	hc := &http.Client{}
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(method, testServer.URL+endpoint, bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		return hc.Do(req)
	}
	req, err := http.NewRequest(method, testServer.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	return hc.Do(req)
}

func populateTables() {
	clients := []models.Client{}
	products := []models.Product{}
	orders := []models.Order{}

	for i := 0; i < 20; i++ {
		var c models.Client
		faker.FakeData(&c)
		clients = append(clients, c)
	}
	for i := 0; i < 20; i++ {
		var p models.Product
		faker.FakeData(&p)
		products = append(products, p)
	}
	for i := 0; i < 20; i++ {
		var o models.Order
		faker.FakeData(&o)
		o.ClientID = int64(i + 1)
		o.ProductID = int64(i + 1)
		orders = append(orders, o)
	}

	testDb.Create(&clients)
	testDb.Create(&products)
	testDb.Create(&orders)
}
