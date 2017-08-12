# GenerateSASTokenGo
Generate a SaS token for Azure Service Hub with Go

See the following page for context [Generate SaS Token](https://docs.microsoft.com/en-us/rest/api/eventhub/generate-sas-token)

You can invoke the function like this:
```
SaS := createSharedAccessToken("http://yournamespace.servicebus.windows.net/orders", "yourpolicyname", "yoursecretkey")
```
