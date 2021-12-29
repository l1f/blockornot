# Block or Not

## the idea
From time to time I scroll through twitter and block many accounts out of self-protection.
That's 3-4 clicks every time, which I find very annoying. 

The idea is to log in with your own twitter account. 
And then to be able to search for hashtags. (not done)

The tweets that can be found there are then displayed one after the other and then there is the possibility 
to swipe left and right to block the account that wrote the tweet or to move to the next tweet. (not done)

## development
To be able to work on the project you need a Twitter development account.

### backend
#### build
```
go build cmd/api/api.go 
```

Generate a config file
```
./api -c ./config.toml init
or 
go build cmd/api/api.go -c ./config.toml init
```

Next, `ConsumerKey` and `ConsumerSecret` must be entered into the configuration. 

Now the backend can be started in the following way:

```
./api -c ./config.toml run
or 
go build cmd/api/api.go -c ./config.toml run
```

### frontend
The frontend is written with react using typescript. 

**Note:** in frontend/src/setupProxy.js is defined how the frontend can talk 
to the backend during development 

#### start

```
yarn start
```