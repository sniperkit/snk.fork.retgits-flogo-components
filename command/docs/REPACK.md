# Flogo - Extended

Usage
-----
Install the package:

```
go get -u github.com/palantir/amalgomate
```

Run the command:

```
amalgomate --config flogo-extended.yml --output-dir outpkg --pkg main
```

The above command runs `amalgomate` on the files specified in `repackage.yml` and writes the output source files into a
new directory called `outpkg`. `outpkg` will contain an `amalgomated` directory that contains all of the repacked
projects and a `main.go` file that contains a `main` method for invoking the repacked libraries.

Configuration
-------------
`amalgomate` uses a configuration file to determine the packages that should be used as input and the name of the 
command that should be used for that package. The configuration is a `yml` file that contains an entry for each program
that should be repackaged:

```yml
packages:
  sample:
    main: github.com/nmiyake/go-sample
  inner:
    main: github.com/nmiyake/go-project/main
    distance-to-project-pkg: 1
```

Each package must have a unique name (this will be the value that the generated Go wrapper will use to reference the
program). The package must specify a `main` package. The package will be resolved in the same way it would if it were in
a Go source file contained in the output directory (including vendoring behavior). If the program being wrapped is in a
subdirectory of a main project, then the `distance-to-project-pkg` parameter can be used to specify the distance between
the `main` package and the project root package. When a program is being wrapped, the project package is copied into the
vendor directory of the output directory, so this parameter can be used in cases where the `main` package is in a
subdirectory of a project but more files need to be copied in order for the import to function correctly.


## References
- github.com/palantir/amalgomate