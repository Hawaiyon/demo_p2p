// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a P2P trade system.",
    "title": "Sample P2P Trade",
    "contact": {
      "email": "hawaiyon@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "hawaiyon.ml",
  "basePath": "/v1",
  "paths": {
    "/debt/{baseUserId}/{toUserId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "transaction"
        ],
        "summary": "根据基准用户ID，返回跟另外一个用户之间的债务情况",
        "operationId": "getDebtInfo",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "基准用户 ID",
            "name": "baseUserId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "另外一个用户的ID",
            "name": "toUserId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Debt"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/loan": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "transaction"
        ],
        "summary": "借款",
        "operationId": "addLoan",
        "parameters": [
          {
            "description": "借款信息",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Loan"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Loan"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/loan/{transactionId}": {
      "get": {
        "description": "Find loan by ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transaction"
        ],
        "summary": "Find loan by ID",
        "operationId": "getLoanById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of transaction",
            "name": "transactionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Loan"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/repayment": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "transaction"
        ],
        "summary": "还款",
        "operationId": "addRepay",
        "parameters": [
          {
            "description": "还款信息",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Repayment"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Repayment"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/repayment/{transactionId}": {
      "get": {
        "description": "Find repay by ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "transaction"
        ],
        "summary": "Find repay by ID",
        "operationId": "getRepayById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of transaction",
            "name": "transactionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Loan"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/user": {
      "post": {
        "description": "创建用户",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "创建用户",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "json格式，包含用户名和初始金额",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    },
    "/user/{userID}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "根据用户id获取具体信息，包括对当前余额、借出总额和借入总额",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "需要获取信息的用户id",
            "name": "userID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "user found",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/api_error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Debt": {
      "type": "object",
      "required": [
        "baseUserId",
        "toUserId",
        "borrow",
        "lend"
      ],
      "properties": {
        "baseUserId": {
          "description": "基准用户 ID",
          "type": "integer",
          "format": "int64"
        },
        "borrow": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "lend": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "toUserId": {
          "description": "另外一个用户的ID",
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Loan": {
      "type": "object",
      "required": [
        "BorrowerId",
        "LenderId",
        "amount"
      ],
      "properties": {
        "BorrowerId": {
          "type": "integer",
          "format": "int64"
        },
        "LenderId": {
          "type": "integer",
          "format": "int64"
        },
        "amount": {
          "type": "integer",
          "format": "int64"
        },
        "createdDate": {
          "type": "string",
          "format": "date-time",
          "readOnly": true
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "status": {
          "description": "Transaction Status",
          "type": "string",
          "enum": [
            "failed",
            "success"
          ],
          "readOnly": true
        }
      },
      "xml": {
        "name": "Loan"
      }
    },
    "Repayment": {
      "type": "object",
      "required": [
        "BorrowerId",
        "LenderId",
        "repay_amount"
      ],
      "properties": {
        "BorrowerId": {
          "type": "integer",
          "format": "int64"
        },
        "LenderId": {
          "type": "integer",
          "format": "int64"
        },
        "createdDate": {
          "type": "string",
          "format": "date-time",
          "readOnly": true
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "repay_amount": {
          "type": "integer",
          "format": "int64"
        },
        "status": {
          "description": "Transaction Status",
          "type": "string",
          "enum": [
            "failed",
            "success"
          ],
          "readOnly": true
        }
      },
      "xml": {
        "name": "Repayment"
      }
    },
    "User": {
      "type": "object",
      "required": [
        "username",
        "balance"
      ],
      "properties": {
        "balance": {
          "type": "integer",
          "format": "int64"
        },
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "total_borrow": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "total_lend": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    },
    "api_error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "用户交易相关接口。任何时候，用户的余额都一定是大于等于0，用户不能还款超过其借款。",
      "name": "transaction"
    },
    {
      "description": "用户相关接口",
      "name": "user"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
