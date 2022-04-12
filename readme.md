clone sample.env to runtime.env

fill-in google client id, client secret and generate secret string for APP_SECRET

the default callback is http://localhost:3000/auth/{provider}/callback
{provider} : OAuth provider e.g google, twitter, facebook...etc.

start the app
go run *.go