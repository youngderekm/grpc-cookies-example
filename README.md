
This code demonstrates how to make gRPC calls create and/or delete session cookies in the gRPC json gateway.  This makes it possible, for example, to setup a session on successful login from a signin call and then delete that session on a signout call.  The gRPC server doesn't have access to the HTTP request or response, but can set metadata flags that custom HTTP and gPRC middleware in the gateway can react to.


To build:
```
go build ./cmd/grpc
go build ./cmd/gateway
```

To run:
```
# in one tab:
./grpc
# in another tab
./gateway
```

To test:
```
curl -v --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"tester","password":"1234"}' \
  http://localhost:8081/v1/authapi/signin
```
You should see a Set-Cookie value in the output, like this:

```
Set-Cookie: grpc-cookies-example-session=MTU0MTE3MTUwOHxEdi1CQkFFQ180SUFBUkFCRUFBQVFfLUNBQUVHYzNSeWFXNW5EQWdBQm5WelpYSkpaQTBxYldGcGJpNVRaWE56YVc5dV80TURBUUVIVTJWemMybHZiZ0hfaEFBQkFRRUdWWE5sY2tsRUFRUUFBQUFKXzRRR0FmMERBNW9BfG1GuOnjcScQCaptcXYeKBjbhYG-cq0ongAMO_vn1dc4; Path=/; Expires=Fri, 02 Nov 2018 19:11:48 GMT; Max-Age=14400; HttpOnly
```
