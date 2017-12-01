# ModelGen

ModelGen generates working database interaction code from reading your MySQL / MariaDB database.



## Usage:

```
Usage:
   [command]

Available Commands:
  generate
  help        Help about any command
  migrate

Flags:
  -c, --connection string   user:pass@host:port
  -d, --database string     name of database
  -h, --help                help for this command
  -o, --output string       path to package (default "generated_models")
  -p, --package string      name of package (default "generated_models")
    	
Example:
modelgen -c root:awesome-password@0.0.0.0:3306 -dmy-db -o models generate
```

