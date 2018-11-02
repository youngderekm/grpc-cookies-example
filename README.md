
This code demonstrates how to make gRPC calls create and/or delete session cookies in the gRPC json gateway.  This makes it possible, for example, to setup a session on successful login from a signin call and then delete that session on a signout call.  The gRPC server doesn't have access to the HTTP request or response, but can set metadata flags that custom HTTP and gPRC middleware in the gateway can react to.

### Building
To compile/build:
```
go build ./cmd/grpc
go build ./cmd/gateway
```

To regenerate code from the gRPC .proto file (only necessary if servicedef.proto has been changed):
```
cd servicedef
go generate
```
The protoc compiler is run under docker to regenerate these `servicedef.pb.go` and `servicedef.pb.gw.go`.

### Running
To run:
```
# in one tab:
./grpc
# in another tab
./gateway
```

### Testing
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

## License
[MIT](https://github.com/youngderekm/grpc-cookies-example/blob/master/LICENSE)
