# GORM Gen Utils
`gormgenutils` is a utility package for generating GORM gen API in Golang. With this package, you can generate type-safe DAO APIs for your application models with minimal effort.

## Installation
To install the `gormgenutils` package, run the following command in your terminal:

```bash
go get github.com/problem-company-toolkit/gormgenutils
```

## Usage
Here are the detailed steps on how to use the gormgenutils package in your Golang project:

1. Create a new directory in your project root called generate. This directory will contain the code to generate the GORM gen API for your models.

2. Inside the generate directory, create a main.go file. In this file, you'll configure the gormgenutils.Generate function and run it to generate the API code.

### Configuring the Generation
In the `generate/main.go` file, import the `gormgenutils` package and call the `gormgenutils.Generate` function:

```go
package main

import (
	"github.com/problem-company-toolkit/gormgenutils"
)

func main() {
	opts := gormgenutils.GenerateOpts{
		PackageName: "your-package-name",
		Models: []gormgenutils.Model{
			// Add the models you want to use gen with
		},
		GenericQueriers: func(g *gormgen.GenericQuerier) {
			// Add your custom generic queriers
		},
	}

	gormgenutils.Generate(opts)
}
```
Replace `your-package-name` with the desired output package name for the generated code. If you leave it empty, the default package name `api` will be used. Add your models to the `Models` slice and configure any custom generic queriers in the `GenericQueriers`.

Generating the API
In your terminal, navigate to the `generate` directory and run the `main.go` file:

```bash
cd generate && go run main.go
```
This command will generate the GORM gen API code in the specified output package or in the default package named `api`.

### Using the Generated Code
After generating the API, you can import the package and use the generated code in your Golang project as follows:

```go
package main

import (
	"fmt"

	"your-project-root/your-package-name"
)

func main() {
	api := your_package_name.NewAPI()

	// Use the API
	record, err := api.GetByID("some-record-id")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Record:", record)
}
```
**Note**: Replace `your-project-root` and `your-package-name` with your actual project root and generated package name, respectively. If you used the default package name, it will be `api`.

## GenericQuerier
The `gormgenutils.GenericQuerier` interface provides an API template for generated queriers. Use this interface to implement custom generic queriers that interact with your models. For example:

```go
type MyModelQuerier interface {
	FindByEmail(email string) ([]*MyModel, error)
}
```
In the `gormgenutils.Generate` options, you can pass in functions that configure new queriers for your models:

```go
opts := gormgenutils.GenerateOpts{
	Model: gormgenutils.Model{
		Model: MyModel{},
		GenericQueriers: func(g *gormgen.GenericQuerier) {
			g.ApplyInterface(MyModelQuerier{})
		},
	},
}
```
By following these steps, you can generate and use GORM gen API with the `gormgenutils` package in your Golang projects.