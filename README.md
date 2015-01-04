npm install grunt-cli -g

npm install

grunt deps

grunt scss

server/conf/mongoConfig.go:

    package conf

    // MongoDB server host
    const MongoHost = "..."

    // MongoDB server port
    const MongoPort = "..."

    // MongoDB database name
    const MongoDatabase = "..."

    // MongoDB username for authorization
    const MongoUser = "..."

    // MongoDB password for authorization
    const MongoPassword = "..."

    // MongoDB full URL
    const MongoURL = "mongodb://" + MongoUser + ":" + MongoPassword + "@" + MongoHost + ":" + MongoPort + "/" + MongoDatabase

server/conf/secret.go:

    package conf

    // Randomly generated strings for using with md5 hash function
    const (
    	Salt1 = "..."
    	Salt2 = "..."
    )

    // CookieStoreKey is needed for storing session values in cookies
    const CookieStoreKey = "..."

cd server
go run app.go
