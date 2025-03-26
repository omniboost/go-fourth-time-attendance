module github.com/omniboost/go-fourth-time-attendance

go 1.23.6

require (
	github.com/gorilla/schema v1.4.1
	github.com/hashicorp/go-multierror v1.1.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require github.com/hashicorp/errwrap v1.0.0 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
