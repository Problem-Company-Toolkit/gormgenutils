package gormgenutils

import (
	"fmt"
	"path"

	"gorm.io/gen"
)

const (
	GORM_DEFAULT_PACKAGE_NAME = "api"
)

type Model struct {
	// Must be a POINTER to the model. Model must be exported from the package (this is applicable for all GORM things in genreal).
	Model           interface{}
	GenericQueriers func(*gen.Generator)
}
type GenerateOpts struct {
	PackageName     string
	Models          []Model
	GenericQueriers func(*gen.Generator)
}

type GenericQuerier interface {
	// where("id=@id")
	GetByID(id string) (gen.T, error)
}

func Generate(
	opts GenerateOpts,
) {
	packageName := opts.PackageName
	if packageName == "" {
		packageName = GORM_DEFAULT_PACKAGE_NAME
	}

	// Should end up being relative to where this is being called from.
	outputPath := path.Join("..", packageName)
	fmt.Printf("Generating Golang API from models")
	fmt.Printf("Package name: %s", packageName)

	allModels := []interface{}{}
	for i := range opts.Models {
		allModels = append(allModels, opts.Models[i].Model)
	}

	cfg := gen.Config{
		OutPath:       outputPath,
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		ModelPkgPath:  packageName,
		FieldNullable: true,
	}

	g := gen.NewGenerator(cfg)

	// Generate basic type-safe DAO API for structs.
	g.ApplyBasic(allModels...)
	g.ApplyInterface(func(GenericQuerier) {}, allModels...)
	if opts.GenericQueriers != nil {
		opts.GenericQueriers(g)
	}
	for i := range opts.Models {
		if opts.Models[i].GenericQueriers != nil {
			opts.Models[i].GenericQueriers(g)
		}
	}

	// Generate the code
	g.Execute()
}
