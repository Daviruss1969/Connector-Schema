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
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$id": "https://gitlab.hardis-group.com/r-d-technique/reflex-platform/deployment/platform-setup/-/raw/develop/100.upgrade-release/connectors/connector-RFX-latest.json?ref_type=heads",
    "title": "Connector",
    "description": "A product from Reflex Platform solution",
    "type": "object",
    "properties": {
        "Name": {
            "description": "Name of the connector",
            "type": "string"
        },
        "Type": {
            "description": "Type of the connector",
            "type": "string"
        },
        "Description": {
            "description": "Description of the connector",
            "type": "string"
        },
        "Version": {
            "description": "Version of the connector",
            "type": "string"
        },
        "RepoUrl": {
            "description": "Repository url of the connector",
            "type": "string"
        },
        "MainClassName": {
            "description": "Main class name of the connector",
            "type": "string"
        },
        "JarName": {
            "description": "",
            "type": "string"
        },
        "SetupJson": {
            "description": "",
            "type": "string"
        },
        "DistributionName": {
            "description": "Distribution name of the connector",
            "type": "string"
        },
        "MetadataJSONParameters": {
            "description": "Parameters of the connector",
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "name": {
                        "description": "Name of the parameter",
                        "type": "string"
                    },
                    "label": {
                        "description": "Label of the parameter",
                        "type": "string"
                    },
                    "help_text": {
                        "description": "Description of the parameter",
                        "type": "string"
                    },
                    "mandatory": {
                        "description": "Describe if the parameter is mandatory",
                        "type": "boolean"
                    },
                    "group_id": {
                        "description": "",
                        "type": "string"
                    },
                    "order_id": {
                        "description": "",
                        "type": "number",
                        "exclusiveMinimum": 0
                    },
                    "type": {
                        "description": "Type of the parameter",
                        "type": "string",
                        "format": "type-format"
                    },
                    "custom_info": {
                        "description": "More informations about the parameter",
                        "type": "object",
                        "properties": {
                            "$suffix": {
                                "description": "The extention of the file for FILE format",
                                "type": "string"
                            },
                            "$regex": {
                                "description": "A regex to match for string format",
                                "type": "string",
                                "format": "regex"
                            },
                            "$default_value": {
                                "description": "A default value for BOOLEAN or SET formats",
                                "type": "string",
                                "enum": ["true", "false", "all"]
                            },
                            "$settype": {
                                "description": "The type of the set (multiple or single)",
                                "type": "string",
                                "enum": ["multiple", "single"]
                            }
                        }
                    }
                },
                "required": ["name", "label", "help_text", "mandatory", "group_id", "order_id", "type"]
            }
        },
        "Permissions": {
            "description": "Permissions of the connector",
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "Client": {
                        "description": "Target of the permission",
                        "type": "string"
                    },
                    "Scopes": {
                        "description": "Scopes of the permission",
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "Service": {
                                    "description": "Service of the scope",
                                    "type": "string"
                                },
                                "Method": {
                                    "description": "Method of the scope",
                                    "type": "string"
                                },
                                "RscType": {
                                    "description": "Type of the scope (projet, orga, connection...)",
                                    "type": "string"
                                },
                                "PlatformReserved": {
                                    "description": "Describe if the platform is reserved",
                                    "type": "boolean"
                                },
                                "TransientPermissions": {
                                    "description": "",
                                    "type": "array"
                                }     
                            },
                            "required": ["Service", "Method", "RscType", "PlatformReserved", "TransientPermissions"]
                        }
                    }
                },
                "required": ["Client", "Scopes"]
            }
        }
    },
    "required": [ "Name", "Type", "Description", "Version", "RepoUrl", "MainClassName", "JarName", "SetupJson", "DistributionName", "MetadataJSONParameters"]
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