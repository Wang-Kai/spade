# spade - CLI tool for [`Go template`]()

Command line tool that render go template with values from YAML files.

## Install

```
go get github.com/Wang-Kai/spade
```

## Usage

Say you have a `template` file like this:

```
Hi, my name {{ .name }}, i am {{ .age }} years ago.
```

add `values.yaml` YAML file like this one:

```yaml
name: steve
age: 100
```

You can compile the template like this:

```
spade -f ./example/template -v ./example/values.yaml
```