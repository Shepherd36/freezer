// Code generated by swaggo/swag. DO NOT EDIT.

package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "ice.io",
            "url": "https://ice.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tokenomics-statistics/adoption": {
            "get": {
                "description": "Returns the current adoption information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokenomics.AdoptionSummary"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics-statistics/top-miners": {
            "get": {
                "description": "Returns the paginated leaderboard with top miners.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "a keyword to look for in the user's username or firstname/lastname",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "max number of elements to return. Default is ` + "`" + `10` + "`" + `.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "number of elements to skip before starting to fetch data",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/tokenomics.Miner"
                            }
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics/{userId}/balance-history": {
            "get": {
                "description": "Returns the balance history for the provided params.\nIf ` + "`" + `startDate` + "`" + ` is after ` + "`" + `endDate` + "`" + `, we go backwards in time: I.E. today, yesterday, etc.\nIf ` + "`" + `startDate` + "`" + ` is before ` + "`" + `endDate` + "`" + `, we go forwards in time: I.E. today, tomorrow, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokenomics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The start date in RFC3339 or ISO8601 formats. Default is ` + "`" + `now` + "`" + ` in UTC.",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The start date in RFC3339 or ISO8601 formats. Default is ` + "`" + `end of day, relative to startDate` + "`" + `.",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The user's timezone. I.E. ` + "`" + `+03:00` + "`" + `, ` + "`" + `-1:30` + "`" + `. Default is UTC.",
                        "name": "tz",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "max number of elements to return. Default is ` + "`" + `24` + "`" + `.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "number of elements to skip before starting to fetch data",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/tokenomics.BalanceHistoryEntry"
                            }
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics/{userId}/balance-summary": {
            "get": {
                "description": "Returns the balance related information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokenomics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokenomics.BalanceSummary"
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics/{userId}/mining-summary": {
            "get": {
                "description": "Returns the mining related information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokenomics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokenomics.MiningSummary"
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "if not found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics/{userId}/pre-staking-summary": {
            "get": {
                "description": "Returns the pre-staking related information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokenomics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokenomics.PreStakingSummary"
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if not allowed",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "if not found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tokenomics/{userId}/ranking-summary": {
            "get": {
                "description": "Returns the ranking related information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokenomics"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokenomics.RankingSummary"
                        }
                    },
                    "400": {
                        "description": "if validations fail",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "if not authorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "if hidden by the user",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "if syntax fails",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    },
                    "504": {
                        "description": "if request times out",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "SOMETHING_NOT_FOUND"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "error": {
                    "type": "string",
                    "example": "something is missing"
                }
            }
        },
        "tokenomics.AdoptionSummary": {
            "type": "object",
            "properties": {
                "milestones": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "achievedAt": {
                                "type": "string",
                                "example": "2022-01-03T16:20:52.156534Z"
                            },
                            "baseMiningRate": {
                                "type": "string",
                                "example": "1,243.02"
                            },
                            "milestone": {
                                "type": "integer",
                                "example": 1
                            },
                            "totalActiveUsers": {
                                "type": "integer",
                                "example": 1
                            }
                        }
                    }
                },
                "totalActiveUsers": {
                    "type": "integer",
                    "example": 11
                }
            }
        },
        "tokenomics.BalanceHistoryBalanceDiff": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "bonus": {
                    "type": "integer",
                    "example": 120
                },
                "negative": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "tokenomics.BalanceHistoryEntry": {
            "type": "object",
            "properties": {
                "balance": {
                    "$ref": "#/definitions/tokenomics.BalanceHistoryBalanceDiff"
                },
                "time": {
                    "type": "string",
                    "example": "2022-01-03T16:20:52.156534Z"
                },
                "timeSeries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tokenomics.BalanceHistoryEntry"
                    }
                }
            }
        },
        "tokenomics.BalanceSummary": {
            "type": "object",
            "properties": {
                "preStaking": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "standard": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "t1": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "t2": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "total": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "totalNoPreStakingBonus": {
                    "type": "string",
                    "example": "1,243.02"
                },
                "totalReferrals": {
                    "type": "string",
                    "example": "1,243.02"
                }
            }
        },
        "tokenomics.Miner": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string",
                    "example": "12345.6334"
                },
                "profilePictureUrl": {
                    "type": "string",
                    "example": "https://somecdn.com/p1.jpg"
                },
                "userId": {
                    "type": "string",
                    "example": "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2"
                },
                "username": {
                    "type": "string",
                    "example": "jdoe"
                }
            }
        },
        "tokenomics.MiningRateBonuses": {
            "type": "object",
            "properties": {
                "extra": {
                    "type": "integer",
                    "example": 300
                },
                "preStaking": {
                    "type": "integer",
                    "example": 300
                },
                "t1": {
                    "type": "integer",
                    "example": 100
                },
                "t2": {
                    "type": "integer",
                    "example": 200
                },
                "total": {
                    "type": "integer",
                    "example": 300
                }
            }
        },
        "tokenomics.MiningRateSummary-string": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string",
                    "example": "1,234,232.001"
                },
                "bonuses": {
                    "$ref": "#/definitions/tokenomics.MiningRateBonuses"
                }
            }
        },
        "tokenomics.MiningRateType": {
            "type": "string",
            "enum": [
                "positive",
                "negative",
                "none"
            ],
            "x-enum-varnames": [
                "PositiveMiningRateType",
                "NegativeMiningRateType",
                "NoneMiningRateType"
            ]
        },
        "tokenomics.MiningRates-tokenomics_MiningRateSummary-string": {
            "type": "object",
            "properties": {
                "base": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "positiveTotalNoPreStakingBonus": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "preStaking": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "standard": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "total": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "totalNoPreStakingBonus": {
                    "$ref": "#/definitions/tokenomics.MiningRateSummary-string"
                },
                "type": {
                    "$ref": "#/definitions/tokenomics.MiningRateType"
                }
            }
        },
        "tokenomics.MiningSession": {
            "type": "object",
            "properties": {
                "endedAt": {
                    "type": "string",
                    "example": "2022-01-03T16:20:52.156534Z"
                },
                "free": {
                    "type": "boolean",
                    "example": true
                },
                "resettableStartingAt": {
                    "type": "string",
                    "example": "2022-01-03T16:20:52.156534Z"
                },
                "startedAt": {
                    "type": "string",
                    "example": "2022-01-03T16:20:52.156534Z"
                },
                "warnAboutExpirationStartingAt": {
                    "type": "string",
                    "example": "2022-01-03T16:20:52.156534Z"
                }
            }
        },
        "tokenomics.MiningSummary": {
            "type": "object",
            "properties": {
                "availableExtraBonus": {
                    "type": "integer",
                    "example": 2
                },
                "miningRates": {
                    "$ref": "#/definitions/tokenomics.MiningRates-tokenomics_MiningRateSummary-string"
                },
                "miningSession": {
                    "$ref": "#/definitions/tokenomics.MiningSession"
                },
                "miningStreak": {
                    "type": "integer",
                    "example": 2
                },
                "remainingFreeMiningSessions": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "tokenomics.PreStakingSummary": {
            "type": "object",
            "properties": {
                "allocation": {
                    "type": "integer",
                    "example": 100
                },
                "bonus": {
                    "type": "integer",
                    "example": 100
                },
                "years": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "tokenomics.RankingSummary": {
            "type": "object",
            "properties": {
                "globalRank": {
                    "type": "integer",
                    "example": 12333
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "latest",
	Host:             "",
	BasePath:         "/v1r",
	Schemes:          []string{"https"},
	Title:            "Tokenomics API",
	Description:      "API that handles everything related to read-only operations for user's tokenomics and statistics about it.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
