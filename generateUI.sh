docker run -d -p 8080:8080 -e SWAGGER_JSON=/swagger.json -v ${PWD}/swagger.json:/swagger.json swaggerapi/swagger-ui
