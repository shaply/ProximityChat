1. Need to fix the JWT token transmission so that it doesn't transmit with a GET request and in the url path of the GET request
2. Need to create the receiver method that listens indefinitely
3. The upgrade to socket doesn't work? Probably because when using Postman, am just sending an HTML request and might need to switch the the websocket thing