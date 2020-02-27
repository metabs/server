Feature: Workspace APIs are up and running
  In order to interact with the workspace
  As a client
  I need to have APIs up and running

  Scenario: Server returns a successful response when creating a new workspace
    Given an HTTP "POST" request "http://localhost/workspaces/":
    """
    {
    	"name":"Personal"
    }
    """
    Then the API must reply with a status code 200
    When given the response body
    Then the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal"
    And the API must reply with a body containing an empty list of collections
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing nil update date

  Scenario: Server returns an existing workspace in the response body when it exists
    Given an existing workspace:
    """
    {
        "id": "518cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal WS",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "GET" request "http://localhost/workspaces/518cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 200
    When given the response body
    Then the API must reply with a body containing:
"""
{"id":"518cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Personal WS","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"}
"""

  Scenario: Server returns an non authorized when another customer access workspace owned by another customer
    Given an existing workspace:
    """
    {
        "id": "518cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal WS",
        "customer_id": "0ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "GET" request "http://localhost/workspaces/518cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 401

  Scenario: Server returns a 404 when searching for a non existing workspace
    When an HTTP "GET" request "http://localhost/workspaces/00000000-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 404


  Scenario: Server returns a 200 when deleting an existing workspace
    Given an existing workspace:
    """
    {
        "id": "398cdfdc-27a5-4c4b-b990-07f460a58dad",
        "name": "Work WS",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "DELETE" request "http://localhost/workspaces/398cdfdc-27a5-4c4b-b990-07f460a58dad":
    """
    """
    Then the API must reply with a status code 200
    When an HTTP "GET" request "http://localhost/workspaces/398cdfdc-27a5-4c4b-b990-07f460a58dad":
    """
    """
    Then the API must reply with a status code 404

  Scenario: Server returns a 404 when deleting a non existing workspace
    When an HTTP "DELETE" request "http://localhost/workspaces/298cdfdc-27a5-4c4b-b990-0ae210a58dac":
    """
    """
    Then the API must reply with a status code 404


  Scenario: Server returns an updated workspace in the response body when a patch request is sent
    Given an existing workspace:
    """
    {
        "id": "667cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Persona_X",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "PATCH" request "http://localhost/workspaces/667cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    {
    "name": "Personal :)"
    }
    """
    Then the API must reply with a status code 200
    When given the response body
    Then the API must reply with a body containing an id as "667cdfdc-27a5-4c4b-b990-07f460a58dac"
    And the API must reply with a body containing a name as "Personal :)"
    And the API must reply with a body containing an empty list of collections
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date

  Scenario: Server returns a list of workspace owned by a customer
    Given an existing workspace:
    """
    {
        "id": "000cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Persona 000",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    And an existing workspace:
    """
    {
        "id": "111cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Persona 111",
        "customer_id": "1233a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    And an existing workspace:
    """
    {
        "id": "222cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Persona 222",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    And an existing workspace:
    """
    {
        "id": "333cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Persona 333",
        "customer_id": "1113a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "GET" request "http://localhost/workspaces/":
    """
    """
    Then the API must reply with a status code 200
    When given the response body as list
    Then the API must reply with a body containing:
"""
[{"id":"000cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Persona 000","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"},{"id":"222cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Persona 222","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"}]
"""

  Scenario: Server returns a workspace with a collection once added
    Given an existing workspace:
    """
    {
        "id": "a3bcdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal Workspace",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "POST" request "http://localhost/workspaces/a3bcdfdc-27a5-4c4b-b990-07f460a58dac/collections":
    """
    {
        "name":"Collection 1"
    }
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal Workspace"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Collection 1"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing nil update date

  Scenario: Server returns a workspace with a 2 collection once added
    Given an existing workspace:
    """
    {
        "id": "a3bcdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal Workspace",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "POST" request "http://localhost/workspaces/a3bcdfdc-27a5-4c4b-b990-07f460a58dac/collections":
    """
    {
        "name":"Collection 1"
    }
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal Workspace"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Collection 1"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing nil update date


  Scenario: Server returns a workspace with a collection once added
    Given an existing workspace:
    """
    {
        "id": "a82edfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal - 0",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "POST" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a58dac/collections":
    """
    {
        "name":"Collection 1"
    }
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal - 0"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Collection 1"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing nil update date
    Then an HTTP "POST" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a58dac/collections":
    """
    {
        "name":"Collection 2"
    }
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal - 0"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Collection 1"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing nil update date
    And the API must reply with a body containing a collections at index 1 containing an id
    And the API must reply with a body containing a collections at index 1 containing a name as "Collection 2"
    And the API must reply with a body containing a collections at index 1 containing an creation date
    And the API must reply with a body containing a collections at index 1 containing nil update date


  Scenario: Server returns a workspace with no collections when those are deleted
    Given an existing workspace:
    """
    {
        "id": "a82edfdc-27a5-4c4b-b990-07f460a58ccd",
        "name": "Personal one",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [
        {
            "id": "569f3341-86ee-431e-8223-951ab8875c86",
            "name": "Collection one",
            "tabs":[],
            "created": "2020-02-26T00:29:42.8565151Z",
            "updated": "0001-01-01T00:00:00Z"
        }
        ],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "DELETE" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a58ccd/collections/569f3341-86ee-431e-8223-951ab8875c86":
    """
    """
    Then the API must reply with a status code 200
    Then an HTTP "GET" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a58ccd":
    """
    """
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal one"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing an empty list of collections

  Scenario: Server returns a workspace with patched collection when updated
    Given an existing workspace:
    """
    {
        "id": "a82edfdc-27a5-4c4b-b990-07f460a01ddc",
        "name": "Personal one",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [
        {
            "id": "569f3341-86ee-431e-8223-951ab8875c86",
            "name": "Collection one",
            "tabs":[],
            "created": "2020-02-26T00:29:42.8565151Z",
            "updated": "0001-01-01T00:00:00Z"
        }
        ],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "PATCH" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a01ddc/collections/569f3341-86ee-431e-8223-951ab8875c86":
    """
    {
        "name":"Collection two"
    }
    """
    Then the API must reply with a status code 200
    Then an HTTP "GET" request "http://localhost/workspaces/a82edfdc-27a5-4c4b-b990-07f460a01ddc":
    """
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal one"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id as "569f3341-86ee-431e-8223-951ab8875c86"
    And the API must reply with a body containing a collections at index 0 containing a name as "Collection two"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing an update after create date

  Scenario: Server returns a workspace with a tab once added
    Given an existing workspace:
    """
    {
        "id": "a3bcdfdc-0123-4c4b-b990-07f460a58dac",
        "name": "Personal",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [
        {
           "id": "569f3341-86ee-431e-8223-951ab8875c86",
           "name": "Google",
           "created": "2020-02-26T00:29:42.8565151Z",
           "updated": "0001-01-01T00:00:00Z"
        }
        ],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "2020-02-25T18:04:13.349522Z"
    }
    """
    When an HTTP "POST" request "http://localhost/workspaces/a3bcdfdc-0123-4c4b-b990-07f460a58dac/collections/569f3341-86ee-431e-8223-951ab8875c86/tabs":
    """
    {
        "title":"Spanner",
        "description":"A database used for the service XYZ",
        "icon":"http://console.google.com/spanner.png",
        "link":"http://console.google.com/spanner"
    }
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Google"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a title as "Spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a description as "A database used for the service XYZ"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a icon as "http://console.google.com/spanner.png"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a link as "http://console.google.com/spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a creation date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing nil update date

  Scenario: Server returns a workspace with patched tab when updated
    Given an existing workspace:
    """
    {
    	"id": "a3bcdfdc-0123-4c4b-b990-07f460a58dac",
    	"name": "Personal",
    	"customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
    	"collections": [{
    		"id": "569f3341-86ee-431e-8223-951ab8875c86",
    		"name": "Google",
    		"tabs": [{
    				"id": "0a2be141-86ee-431e-8223-951ab8875c86",
    				"title": "Spanner",
    				"description": "A database used for the service XYZ",
    				"icon": "http://console.google.com/spanner.png",
    				"link": "http://console.google.com/spanner",
    				"created": "2020-02-26T00:29:42.8565151Z",
    				"updated": "0001-01-01T00:00:00Z"
    			},
    			{
    				"id": "0a2be141-86ee-431e-8223-951ab8875c87",
    				"title": "StackDriver Logger",
    				"description": "A logger UI used for the service XYZ",
    				"icon": "http://console.google.com/logger.png",
    				"link": "http://console.google.com/logger",
    				"created": "2020-02-27T00:29:42.8565151Z",
    				"updated": "0001-01-01T00:00:00Z"
    			}
    		],
    		"created": "2020-02-26T00:29:42.8565151Z",
    		"updated": "0001-01-01T00:00:00Z"
    	}],
    	"created": "2019-02-25T16:04:13.349522Z",
    	"updated": "2020-02-25T18:04:13.349522Z"
    }
    """
    When an HTTP "PUT" request "http://localhost/workspaces/a3bcdfdc-0123-4c4b-b990-07f460a58dac/collections/569f3341-86ee-431e-8223-951ab8875c86/tabs/0a2be141-86ee-431e-8223-951ab8875c87":
    """
    {
        "title":"Spanner",
        "description":"A database used for the service XYZ",
        "icon":"http://console.google.com/spanner.png",
        "link":"http://console.google.com/spanner"
    }
    """
    Then the API must reply with a status code 200
    When an HTTP "GET" request "http://localhost/workspaces/a3bcdfdc-0123-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Google"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing an id as "0a2be141-86ee-431e-8223-951ab8875c87"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing a title as "Spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing a description as "A database used for the service XYZ"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing a icon as "http://console.google.com/spanner.png"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing a link as "http://console.google.com/spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing a creation date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 1 containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing an id as "0a2be141-86ee-431e-8223-951ab8875c86"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a title as "Spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a description as "A database used for the service XYZ"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a icon as "http://console.google.com/spanner.png"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a link as "http://console.google.com/spanner"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a creation date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing nil update date

  Scenario: Server returns a 200 when deleting an existing tab
    Given an existing workspace:
    """
    {
    	"id": "a3bcdfdc-0123-4c4b-b990-07f460a58dac",
    	"name": "Personal",
    	"customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
    	"collections": [{
    		"id": "569f3341-86ee-431e-8223-951ab8875c86",
    		"name": "Google",
    		"tabs": [{
    				"id": "0a2be141-86ee-431e-8223-951ab8875c86",
    				"title": "Spanner",
    				"description": "A database used for the service XYZ",
    				"icon": "http://console.google.com/spanner.png",
    				"link": "http://console.google.com/spanner",
    				"created": "2020-02-26T00:29:42.8565151Z",
    				"updated": "0001-01-01T00:00:00Z"
    			},
    			{
    				"id": "0a2be141-86ee-431e-8223-951ab8875c87",
    				"title": "StackDriver Logger",
    				"description": "A logger UI used for the service XYZ",
    				"icon": "http://console.google.com/logger.png",
    				"link": "http://console.google.com/logger",
    				"created": "2020-02-27T00:29:42.8565151Z",
    				"updated": "0001-01-01T00:00:00Z"
    			}
    		],
    		"created": "2020-02-26T00:29:42.8565151Z",
    		"updated": "0001-01-01T00:00:00Z"
    	}],
    	"created": "2019-02-25T16:04:13.349522Z",
    	"updated": "2020-02-25T18:04:13.349522Z"
    }
    """
    When an HTTP "DELETE" request "http://localhost/workspaces/a3bcdfdc-0123-4c4b-b990-07f460a58dac/collections/569f3341-86ee-431e-8223-951ab8875c86/tabs/0a2be141-86ee-431e-8223-951ab8875c86":
    """
    """
    Then the API must reply with a status code 200
    When an HTTP "GET" request "http://localhost/workspaces/a3bcdfdc-0123-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 200
    And given the response body
    And the API must reply with a body containing an id
    And the API must reply with a body containing a name as "Personal"
    And the API must reply with a body containing an creation date
    And the API must reply with a body containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing an id
    And the API must reply with a body containing a collections at index 0 containing a name as "Google"
    And the API must reply with a body containing a collections at index 0 containing an creation date
    And the API must reply with a body containing a collections at index 0 containing an update after create date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing an id as "0a2be141-86ee-431e-8223-951ab8875c87"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a title as "StackDriver Logger"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a description as "A logger UI used for the service XYZ"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a icon as "http://console.google.com/logger.png"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a link as "http://console.google.com/logger"
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing a creation date
    And the API must reply with a body containing a collections at index 0 containing a tab at index 0 containing nil update date
