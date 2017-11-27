init:
	go-bindata tmpl/model.html tmpl/x_helpers.html tmpl/x_helpers_test.html && go install
