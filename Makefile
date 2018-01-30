init:
	go-bindata tmpl/model.html tmpl/x_helpers.html tmpl/x_helpers_test.html && go install

build:
	# Named so it doesn't conflict with the installed modelgen.
	go build -o modelgen_local