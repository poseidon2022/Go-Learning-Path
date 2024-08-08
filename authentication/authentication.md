## Implementation of Authentication and Authorization functions

* Using JWT and middleware functions


## dependencies

$ go get github.com/dgrijalva/jwt-go
$ go get golang.org/x/crypto/bcrypt //hashing passwords and all

//  user_name and id for token generation


in the authorization part the key returneed by the call back function is used by the parse function to verify the token and all


this is how it is supposed to look like

token, err := jwt.Parse(passedToken, func(token *jwt.Token) (interface{}, error) {
    if ok := token.Method.*JWTSHAmnamn(); !ok {
        error
        return nil, error
    } 
    
    reuturn jwtSecret, nil
})
