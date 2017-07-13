## apiServer
apiServer based-on a micro-service platform, it provide flexibility and hot-plugin ability

### Dependencies
The project `apiServer` depends on the `fabric-cli` tools and `beego` project. 
1. The `fabric-cli` is based on `fabric-sdk-go` and provides the iteraction with `fabric-ca` and `fabric`. Also it parses the protocal message such as `transaction`, `block`, `installed chaincode`, which is very helpful. 
2. The `beego` project is a RESTFUL api platform.

### ADD New Function
1. add models
2. add controllers
3. add router

### Swagger
1. `bee generate docs`
2. `bee run watchall`
Open `IE explorer` and visit `localhost:8080/swagger` to test the restful api.  
Note: @router /func [get/post], the 'func' must be different the func name, it's case-insensitive.

#### Swagger comment
the following line reside on the head of this file, and use // command on each line
//@APIVersion 1.0.0
//@Title beego  API
//Description beego has a very cool tools to autogenerate documents for your API
//@Contact warm3snow@linux.com
//@TermsOfServiceUrl http://beego.me/
//@License Apache 2.0
//@LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
*/
