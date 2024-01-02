<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">Connector Schema</h3>

  <p align="center">
    Validation connector structure and syntax
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

The Connector Schema validates a given input file according to the specified format of an N1 connector. If the input file is not ok, it will print a message according to the error in order to understand the mistake.

Below the schema that validates the input :

```json
{
	"type": "object",
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"pattern": "#",
	"properties": {
		"Name": {
			"type": "string",
			"pattern": "[A-Za-z]*"
		},
		"Type": {
			"type": "string",
			"pattern": "[A-Z]*"
		},
		"Description": {
			"type": "string"
		},
		"Version": {
			"type": "string",
			"pattern": "[0-9]+\\.[0-9]*\\.[0-9]*"
		},
		"RepoUrl": {
			"type": "string",
			"pattern": "https?://(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"
		},
		"MainClassName": {
			"type": "string",
			"pattern": "MainClassName"
		},
		"JarName": {
			"type": "string",
			"pattern": "([a-zA-Z\\d]*-)+[0-9]+\\.[0-9]*\\.[0-9]*\\.jar"
		},
		"SetupJson": {
			"type": "string"
		},
		"DistributionName": {
			"type": "string",
			"pattern": "([a-zA-Z\\d]*-)+[0-9]+\\.[0-9]*\\.[0-9]*"
		},
		"MetadataJSONParameters": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"Name": {
						"type": "string"
					},
					"Label": {
						"type": "string"
					},
					"Help_text": {
						"type": "string"
					},
					"Mandatory": {
						"type": "boolean"
					},
					"Group_id": {
						"type": "string"
					},
					"Order_id": {
						"type": "integer",
						"pattern": "[0-9]*"
					},
					"Type": {
						"type": "string",
						"pattern": "FILE|BOOLEAN|STRING|SET|PASSWORD"
					},
					"Custom_info": {
            "type": "object",
						"properties": {
							"$suffix": {
									"type": "string"
							},
							"$regex": {
									"type": "string",
									"format": "regex"
							},
							"$default_value": {
									"type": "string",
									"enum": ["true", "false", "all"]
							},
							"$settype": {
									"type": "string",
									"enum": ["multiple", "single"]
							}
						}
					}
				},
        "required": ["Name", "Label", "Help_text", "Mandatory", "Group_id", "Order_id", "Type"]
			}
		},
		"Permissions": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"Client": {
						"type": "string",
						"pattern": "[A-Za-z]+[A-Za-z0-9\\-]*"
					},
					"Scopes": {
						"type": "array",
						"items": {
							"type": "object",
							"properties": {
								"Service": {
									"type": "string",
									"pattern": "[A-Za-z]+[A-Za-z0-9\\-]*"
								},
								"Method": {
									"type": "string",
									"pattern": "[A-Za-z]+[A-Za-z0-9\\-]*"
								},
								"RscType": {
									"type": "string",
									"pattern": "Project|..."
								},
								"PlatformReserved": {
									"type": "boolean"
								}
							},
              "required": ["Service", "Method", "RscType", "PlatformReserved"]
						}
					}
				},
        "required": ["Client", "Scopes"]
			}
		},
		"Actions": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"ActionId": {
						"type": "string",
						"pattern": "$start|$stop|$delete|$create|[a-z]+[a-z0-9\\-]*"
					},
					"Label": {
						"type": "string",
						"pattern": "[A-Za-z0-9]+[A-Za-z0-9\\-]*"
					}
				},
        "required": ["ActionId"]
			}
		},
		"Resources": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"Type": {
						"type": "string",
						"enum": ["NATS","FLINK","KC"]
					},
					"Object": {
						"type": "string"
					},
					"CustomInfo": {
						"type": "object"
					}
				},
        "required": ["Type", "Object"]
			}
		}
	},
  "required": ["Name", "Type", "Description", "Version", "MainClassName", "JarName", "MetadataJSONParameters", "Permissions", "Actions"]
}
```

### Built With
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


<!-- GETTING STARTED -->
## Getting Started
### Prerequisites

To use the Connector Schema, you need to be able to compile go file. If you can't, check this link to get started : [https://go.dev/doc/install](https://go.dev/doc/install)

Type this command to be sure the installation is correct
  ```sh
  go version
  ```

### Installation


1. Clone the repo
   ```sh
   git clone https://github.com/Daviruss1969/Connector-Schema.git
   ```
2. Go into the directory
    ```sh
    cd Connector-Schema
    ```
3. If you are not using Linux, you can build the binary (to be sure the binary will run according to your OS)
   ```sh
   go build -o csch.exe
   ```
4. Up to you to adding it to your $PATH if you want to call it everywhere
    1. Linux :
         1. Add the following line at the end of your ~/.bashrc
        ```sh
        export PATH="/home/user/yourPATH/ConnectorSchema:$PATH"
        ```
        1. run the following command 
        ```sh
        . ~/.bashrc
        ```
     2. Windows :
        
        Follow this tutorial https://stackoverflow.com/questions/44272416/how-to-add-a-folder-to-path-environment-variable-in-windows-10-with-screensho

<!-- USAGE EXAMPLES -->
## Usage

You can now use the Connector Schema 
```sh
./csch.exe -i "inputFilePath"
```

You need to give one valid file path.

The file contains the information about the connector, the field "MetadataJSONParameters" can be parsed or not, the 2 cases will work.