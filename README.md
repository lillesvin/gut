# Gut

Gut is a minimal reimplementation of GNU `cut` that supports defining
delimiters as regular expressions.

## Command line options

`-d/-delimiter REGEX` Regular expression to use for splitting input into
fields. Default: `\s+`.

`-f/-fields FIELDS` Collection of fields to output. Can be specified as a
comma separated list of fields and ranges. (Note: Like `cut`, first field is 1,
not 0.) Ranges can be defined as: `2-4`, `2:4`, `2..4` or `2...4`. Default:
Output all fields.

## Examples

```
$ echo "1,2, 3. 4,5" | gut -d "[,\.]\s?" -f 1,3-5
1 3 4 5

$ echo "1,2,3" | gut -d , -d 2,3,1
2 3 1
```

## Install

Currently this is the only option:

```
go get github.com/lillesvin/gut
```

## Todo

 * Access fields by numbers relative to the last field.
 * Crop field output to X bytes

