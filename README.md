# ModelGen

ModelGen generates working database interaction code from reading your MySQL / MariaDB database.



## Usage:

```
Usage of modelgen:
  -d string
      name of database
  -dsn string
    	root:@tcp(localhost:3306)/database_name?parseTime=true
  -o string
    	path to package (default "generated_models")
  -p string
    	name for generated package (default "models")
    	
    	
Example:
modelgen -d databaseName -o models -dsn user:pass@tcp(localhost:3306)/databaseName?parseTime=true
```

