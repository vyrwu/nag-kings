docker run -d -p 8080:8080 -e SWAGGER_JSON=/swaggerui/swagger.json -v ${PWD}/swagger.json:/swagger.json swaggerapi/swagger-ui
