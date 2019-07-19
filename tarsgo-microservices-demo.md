# microservices-demo三个Go改写
### Payment
**原接口：**
- GET /health
	- 传入值: 无 
	- 传出值: 
	````
	{
	    "health": [
	        {
	            "service": "payment",
	            "status": "OK",
	            "time": "2019-07-17 09:35:16.7395662 +0800 CST m=+2.569047501"
	        }
	    ]
	}
	````
- POST/payment
	- 传入值: 
	 ````
	 
	{
		"Amount":40
	}
	````
	- 传出值: 
	````
	{
	    "authorised": true,
	    "message": "Payment authorised"
	}
	````

**TarsGo版本：**
- Health()
	- 传入值: 无
	- 传出值: 
	````
	{
	    "health": [
	        {
	            "service": "payment",
	            "status": "OK",
	            "time": "2019-07-17 09:35:16.7395662 +0800 CST m=+2.569047501"
	        }
	    ]
	}, nil
	````
	- 示例调用: 
	````
	app := new(MicroserviceDemo.DoPayment)  
	comm.StringToProxy(obj, app)  
	  
	health, healthErr := app.Health()  
	if healthErr != nil {  
	   fmt.Println(healthErr)  
	   return  
	}  
	
	fmt.Println("resp: ", health)
	````
- Authorise(amount float32)
	- 传入值: amount float32
	- 传出值: 
	````
	{
	    "authorised": true,
	    "message": "Payment authorised"
	}
	````
	- 说明:
	 ````
	var ErrInvalidPaymentAmount = errors.New("Invalid payment amount")
	if amount == 0 {  
	   return "", ErrInvalidPaymentAmount  
	}  
	if amount < 0 {  
	   return "", ErrInvalidPaymentAmount  
	}  
	authorised := false  
	message := "Payment declined"  
	service := Service{99.99}  
	fmt.Println(amount)  
	if amount <= service.declineOverAmount {  
		authorised = true  
		message = "Payment authorised"  
	} else {  
	   message = fmt.Sprintf("Payment declined: amount exceeds %.2f", service.declineOverAmount)  
	}
	````
	- 示例调用: 
	````
	app := new(MicroserviceDemo.DoPayment)  
	comm.StringToProxy(obj, app)  
	  
	authorisation, authorisationErr := app.Authorise(80)   
	if authorisationErr != nil {  
	   fmt.Println(authorisationErr)  
	   return  
	}  
	fmt.Println("resp: ", authorisation)
	````

### Catalogue

**原接口：**
- GET /catalogue
 	- 传入值: 无 
	- 传出值: 
	````
			[
			  {
			    "id": "string",
			    "name": "string",
			    "description": "string",
			    "imageUrl": [
			      "string"
			    ],
			    "price": 0,
			    "count": 0,
			    "tag": [
			      "string"
			    ]
			  }
			]
	````
	- GET /catalogue/{id}
 	- 传入值: id
	- 传出值: 
	````	
			  {
			    "id": "string",
			    "name": "string",
			    "description": "string",
			    "imageUrl": [
			      "string"
			    ],
			    "price": 0,
			    "count": 0,
			    "tag": [
			      "string"
			    ]
			  }	
	````
- GET /catalogue/tags
 	- 传入值: 
	- 传出值: 
	````
		{
		  "size": 0
		}
	````
- GET /tags
 	- 传入值: 
	- 传出值: 
			{
			  "tags": [
			    "string"
			  ]
			}
			
**TarsGo版本：**
- Catalogue()
	- 传入值: 
	- 传出值: 
	````
	[
			  {
			    "id": "string",
			    "name": "string",
			    "description": "string",
			    "imageUrl": [
			      "string"
			    ],
			    "price": 0,
			    "count": 0,
			    "tag": [
			      "string"
			    ]
			  }
			]
	````
	- 示例调用: 
	````
	catalogue, catalogueErr := app.Catalogue()  
	if catalogueErr != nil {  
	   fmt.Println(catalogueErr)  
	   return  
	}  
	fmt.Println("resp: ", catalogue)
	````
- CatalogueId(string id)
	- 传入值: string id
	- 传出值: 
	````
			  {
			    "id": "string",
			    "name": "string",
			    "description": "string",
			    "imageUrl": [
			      "string"
			    ],
			    "price": 0,
			    "count": 0,
			    "tag": [
			      "string"
			    ]
			  }
	````
	- 示例调用:  同上
- Count()
 	- 传入值: 
	- 传出值: 
	````
		{"size":10}
	````
	- 示例调用:  同上
- Tags()
 	- 传入值: 
	- 传出值: 
	````
		{"Tags":["brown","geek","formal","blue","skin","red","action","sport","black","magic","green"]}
	````
	- 示例调用:  同上
	
### Users
**原接口：(数据库中数据未发现email字段)**
- GET /login (测试无效永远返回空，应该为登录返回登录人信息)
- POST /register
 	- 传入值: 
 	````
 	{
	  "username": "string",
	  "password": "string",
	  "email": "string"
	}
 	````
	- 传出值: 
````{"id": "string"}````

- GET /customers
	- 传入值: 
	- 传出值: 
````	
 {"_embedded":{"customer":[{"firstName":"Eve","lastName":"Berger","username":"Eve_Berger","id":"57a98d98e4b00679b4a830af",  
  "_links":{"addresses":{"href":"http://user/customers/57a98d98e4b00679b4a830af/addresses"}, "cards":{"href":"http://user/customers/57a98d98e4b00679b4a830af/cards"},"customer":{"href":"http://user/customers/57a98d98e4b00679b4a830af"},  
  "self":{"href":"http://user/customers/57a98d98e4b00679b4a830af"}}}]}}
  ````

- DELETE /customers/{id}
	- 传入值:  id
	- 传出值: 
````
{
  "status": true
}
````
  - GET/customers/{id}
	- 传入值:  id
	- 传出值: 
````
	{
	  "firstName": "string",
	  "lastName": "string",
	  "username": "string",
	  "_links": {
	    "self": {
	      "href": "string"
	    },
	    "customer": {
	      "href": "string"
	    },
	    "addresses": {
	      "href": "string"
	    },
	    "cards": {
	      "href": "string"
	    }
	  }
	}
````
  - GET /customers/{id}/cards
	- 传入值:  id
	- 传出值: 
````
{
  "_embedded": {
    "card": [
      {
        "longNum": "string",
        "expires": "string",
        "ccv": "string",
        "_links": {
          "self": {
            "href": "string"
          },
          "card": {
            "href": "string"
          }
        }
      }
    ]
  },
  "_links": {},
  "page": {}
}
````
  - GET /customers/{id}/address
	- 传入值:  id
	- 传出值: 
````
{
  "_embedded": {
    "address": [
      {
        "number": "string",
        "street": "string",
        "city": "string",
        "postcode": "string",
        "country": "string",
        "_links": {
          "self": {
            "href": "string"
          },
          "address": {
            "href": "string"
          }
        }
      }
    ]
  },
  "_links": {},
  "page": {}
}
````
 - GET /cards
	- 传入值:  
	- 传出值: 
````
	{
	  "_embedded": {
	    "card": [
	      {
	        "longNum": "string",
	        "expires": "string",
	        "ccv": "string",
	        "_links": {
	          "self": {
	            "href": "string"
	          },
	          "card": {
	            "href": "string"
	          }
	        }
	      }
	    ]
	  },
	  "_links": {},
	  "page": {}
	}
````
 - POST /cards
	- 传入值:  
````
{
  "longNum": "string",
  "expires": "string",
  "ccv": "string",
  "userID": "string"
}
````
	- 传出值: 
````
	{
	  "id": "string"
	}
````
 - DELETE /cards/{id}
	- 传入值:  id
	- 传出值: 
````
	{
	  "status": true
	}
````
 - GET /cards/{id}
	- 传入值:  id
	- 传出值: 
````
{
  "longNum": "string",
  "expires": "string",
  "ccv": "string",
  "_links": {
    "self": {
      "href": "string"
    },
    "card": {
      "href": "string"
    }
  }
}
````


 - GET /addresses
	- 传入值:  
	- 传出值: 
````
{
  "_embedded": {
    "address": [
      {
        "number": "string",
        "street": "string",
        "city": "string",
        "postcode": "string",
        "country": "string",
        "_links": {
          "self": {
            "href": "string"
          },
          "address": {
            "href": "string"
          }
        }
      }
    ]
  },
  "_links": {},
  "page": {}
}
````
 - POST /addresses
	- 传入值:  
````
{
  "street": "string",
  "number": "string",
  "country": "string",
  "city": "string",
  "postcode": "string",
  "userID": "string"
}
````
	- 传出值: 
````
	{
	  "id": "string"
	}
````
 - DELETE /addresses/{id}
	- 传入值:  id
	- 传出值: 
````
	{
	  "status": true
	}
````
 - GET /addresses/{id}
	- 传入值:  id
	- 传出值: 
````
{
  "number": "string",
  "street": "string",
  "city": "string",
  "postcode": "string",
  "country": "string",
  "_links": {
    "self": {
      "href": "string"
    },
    "address": {
      "href": "string"
    }
  }
}
````

**TarsGo版本：**
- login(username string, password string)
	- 传入值: username string, password string
	- 传出值: 
	 ````
	 {"firstName":"User",
	 "lastName":"Name",
	 "username":"user",
	 "_id":"57a98d98e4b00679b4a830b2"
	 }
	````	
	- 示例调用:  同上
- register(string username, string password, string email, string first, string last)
	- 传入值: string username, string password, string email, string first, string last
	- 传出值: 
	 ````
	 {"id" :"5d316cdc51c9832a6c14bb25"}
	````	
	- 示例调用:  同上

- register(string username, string password, string email, string first, string last)
	- 传入值: string username, string password, string email, string first, string last
	- 传出值: 
	 ````
	 {"id" :"5d316cdc51c9832a6c14bb25"}
	````	
	- 示例调用:  同上
- customers()
	- 传入值: 
	- 传出值: 
	 ````
	  [{"firstName":"Eve","lastName":"Berger","username":"Eve_Berger","_id":"57a98d98e4b00679b4a830af"}]
	````	
	- 示例调用:  同上
- deleteCustomerById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	{
	  "status": true
	}
	````	
	- 示例调用:  同上
 - findCustomerById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	{"firstName":"333","lastName":"444","username":"user2","_id":"5d316c9251c98337fc10b24d"}
	````	
	- 示例调用:  同上
 - findCustomerCardById(string id)
	- 传入值:  string id
	- 传出值: 
	 ````
	  [{"longNum":"5953580604169678","expires":"08/19","ccv":"678","_id":"57a98d98e4b00679b4a830ae"}]
	````	
	- 示例调用:  同上
 - findCustomerAddressById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	[{"street":"Whitelees Road","number":"246","country":"United Kingdom","city":"Glasgow","postcode":"G67 3DL","_id":"57a98d98e4b00679b4a830ad"}]
	````	
	- 示例调用:  同上
 - Cards()
	- 传入值: 
	- 传出值: 
	 ````
	 [{"longNum":"5953580604169678","expires":"08/19","ccv":"678","_id":"57a98d98e4b00679b4a830ae"}]
	````	
	- 示例调用:  同上
 - AddCard(string longNum, string expires, string ccv, string userID)
	- 传入值: string longNum, string expires, string ccv, string userID
	- 传出值: 
	 ````
	  {"id" :"5d31b0b551c9832478801665"}
	````	
	- 示例调用:  同上
 - DeleteCardById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	  {"status": true}
	````	
	- 示例调用:  同上
 - FindCardById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	{"longNum":"5953580604169678","expires":"08/19","ccv":"678","_id":"57a98d98e4b00679b4a830ae"}
	````	
	- 示例调用:  同上

 - Addresses()
	- 传入值: 
	- 传出值: 
	 ````
	 [{"street":"Whitelees Road","number":"246","country":"United Kingdom","city":"Glasgow","postcode":"G67 3DL","_id":"57a98d98e4b00679b4a830ad"}]
	````	
	- 示例调用:  同上
 - AddAddress(string street, string number, string country, string city, string postcode, string userID)
	- 传入值: string street, string number, string country, string city, string postcode, string userID
	- 传出值: 
	 ````
	  {"id" :"5d31b0b551c9832478801665"}
	````	
	- 示例调用:  同上
 - DeleteAddressById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	  {"status": true}
	````	
	- 示例调用:  同上
 - FindAddressById(string id)
	- 传入值: string id
	- 传出值: 
	 ````
	{"street":"my road","number":"3","country":"UK","city":"London","postcode":"","_id":"57a98ddce4b00679b4a830d1"}
	````	
	- 示例调用:  同上


