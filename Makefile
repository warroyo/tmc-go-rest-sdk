SWAGGER_URL="https://vdc-download.vmware.com/vmwb-repository/dcr-public/faa99fd1-ee55-452c-9954-0448e2989ec4/fadf9bf3-1bad-461a-a57f-a73039e93ea7/public-swagger.yaml"

generate: download-swagger strip-swagger generate-swagger

download-swagger:
	echo "Downloading TMC swagger spec"
	curl -s -o public-swagger.yaml "${SWAGGER_URL}"

strip-swagger:
	sed -i -e 's/vmware\.tanzu\.manage\.v1alpha1\.//' public-swagger.yaml

generate-swagger:
	swagger generate client -f public-swagger.yaml --skip-validation

clean-swagger:
	rm -Rf client models

