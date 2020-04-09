docker run --rm -v ${PWD}:/local --network="host" swaggerapi/swagger-codegen-cli generate \
-i http://localhost:8080/swagger.json \
-l js \
-o /local/out/js
