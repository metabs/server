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
        "id": "298cdfdc-27a5-4c4b-b990-07f460a58dac",
        "name": "Personal WS",
        "customer_id": "1ae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "GET" request "http://localhost/workspaces/298cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 200
    When given the response body
    Then the API must reply with a body containing:
"""
{"id":"298cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Personal WS","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","collections":[],"created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"}
"""

  Scenario: Server returns a 404 when searching for a non existing workspace
    When an HTTP "GET" request "http://localhost/workspaces/298cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 404


  Scenario: Server returns a 200 when deleting an existing workspace
    Given an existing workspace:
    """
    {
        "id": "398cdfdc-27a5-4c4b-b990-07f460a58dad",
        "name": "Work WS",
        "customer_id": "bae3a55d-2c69-4679-808e-1c7772405281",
        "collections": [],
        "created": "2019-02-25T16:04:13.349522Z",
        "updated": "0001-01-01T00:00:00Z"
    }
    """
    When an HTTP "DELETE" request "http://localhost/workspaces/398cdfdc-27a5-4c4b-b990-07f460a58dad":
    """
    """
    Then the API must reply with a status code 200
    When an HTTP "GET" request "http://localhost/workspaces/298cdfdc-27a5-4c4b-b990-07f460a58dac":
    """
    """
    Then the API must reply with a status code 404

  Scenario: Server returns a 404 when deleting a non existing workspace
    When an HTTP "DELETE" request "http://localhost/workspaces/298cdfdc-27a5-4c4b-b990-07f460a58dac":
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
[{"id":"000cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Persona 000","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","collections":[],"created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"},{"id":"222cdfdc-27a5-4c4b-b990-07f460a58dac","name":"Persona 222","customer_id":"1ae3a55d-2c69-4679-808e-1c7772405281","collections":[],"created":"2019-02-25T16:04:13.349522Z","updated":"0001-01-01T00:00:00Z"}]
"""
