# go-json-config
Go jsonconfig is a very simple and easy to use configuration library, allowing JSON based config files for your Go application.
Configuration provider read configuration data from config.json file. You can get the string value of a configuration by its section, or
bind an interface to a valid JSON section.

Consider the following config.json file:

```json
{
  "ConnectionStrings": {
    "DbConnection": "Server=.;User Id=app;Password=123;Database=Db",
    "LogDbConnection": "Server=.;User Id=app;Password=123;Database=Log"
  },
  "Caching": {
    "ApplicationKey": "key",
    "Host": "127.0.01"
  },
  "Website": {
    "ActivityLogEnable": "true",
    "ErrorMessages": {
      "InvalidTelephoneNumber": "Invalid Telephone Number",
      "RequestNotFound": "Request Not Found",
      "InvalidConfirmationCode": "Invalid Confirmation Code"
    }
  },
  "Services": {
    "List": [
      {
        "Id": 1,
        "Name": "Service1"
      },
      {
        "Id": 2,
        "Name": "Service2"
      },
      {
        "Id": 3,
        "Name": "Service3"
      }
    ]
  }
}
```
The following code displays how to access several of the preceding configurations settings.
You can get config value via GetSection function with specifying Json sections as string parameter split by ":"

```Go
c, err :=  jsonconfig.GetSection("ConnectionStrings:DbConnection")
```
Any valid JSON is a valid configuration type, you also can bind a struct via jsonconfig. For example, Caching configuration can be
 bind to valid struct:
 
```Go
type Caching struct {
	ApplicationKey string
	Host           string
}

var c Caching
err = jsonconfig.Bind(&c, "Caching")
```
 In the following example, we can bind the services config in the configuration file to a slice
 
 ```Go
type Service struct {
	Id   int
	Name string
}

var s []Service
err = jsonconfig.Bind(&s, "Services:List")
```
