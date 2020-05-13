module github.com/ivan-avalos/linkbucket-go/server

go 1.13

require (
	github.com/badoux/goscraper v0.0.0-20190827161153-36995ce6b19f
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/ivan-avalos/gorm-paginator/pagination v0.0.0-20200420193221-15ef16f02ab0
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.16
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/metal3d/go-slugify v0.0.0-20160607203414-7ac2014b2f23
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/ivan-avalos/linkbucket-go => ./
