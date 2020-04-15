# RBAC-learn

+ A restful api backend based on RBAC written by [gin](https://github.com/gin-gonic/gin), [gorm](https://github.com/jinzhu/gorm), [casbin](https://github.com/casbin/casbin) and [xdi](https://github.com/Aoi-hosizora/ahlib)

### Functions

+ Authorization bases on jwt
+ User CRUD in mysql
+ Token saved in Redis and Refresh
+ RBAC access management through url

### Run

```bash
# Change database config in ./src/config/config.yaml

go build -i -o ./build/rbac-learn.out main.go
./rbac-learn.out
```

### References

+ [Tutorial: Integrate Gin with Casbin](https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0)
+ [用Go写后台系统API--记录心得(二)](https://studygolang.com/topics/6999)
+ [Goの認証ライブラリCasbinをGin,GORMと連携する](https://www.zaneli.com/blog/20181203)
